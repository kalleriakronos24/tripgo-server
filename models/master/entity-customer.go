package master

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	database "github.com/kalleriakronos24/khaimal-group/db"
	"github.com/kalleriakronos24/khaimal-group/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CustomerOrm struct {
	db *gorm.DB
}

type Customer struct {
	ID           uuid.UUID `gorm:"index:id,unique;type:uuid;default:gen_random_uuid();" json:"id"`
	Name         string    `json:"name" gorm:"not null" binding:"required"`
	Phone        string    `json:"phone,omitempty" gorm:"default:NULL"`
	Status       string    `json:"status,omitempty" binding:"required" gorm:"not null;default:active;"`
	RefferalCode string    `json:"refferalCode,omitempty" gorm:"default:NULL"`

	CredentialsID uuid.UUID    `json:"credentialsId" gorm:"type:uuid;not null"`
	CreatedBy     uuid.UUID    `json:"createdBy,omitempty" gorm:"type:uuid;default:NULL"`
	UpdatedBy     uuid.UUID    `json:"updatedBy,omitempty" gorm:"type:uuid;default:NULL"`
	Credentials   *Credentials `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:CredentialsID;references:ID" json:"credentials,omitempty"`

	types.DefaultModelProperty
}

type CustomerModelAction interface {
	GetOneByID(id uuid.UUID) (m Customer, err error)
	GetOneByCustomerName(Customername string) (m Customer, err error)
	GetAllCustomerPaginated(c *gin.Context) (*database.Pagination, error)
	InsertCustomer(p Customer, tx *gorm.DB) (err error)
	DeleteCustomer(id uuid.UUID, tx *gorm.DB) (err error)
}

func NewCustomerAction(db *gorm.DB) CustomerModelAction {
	return &CustomerOrm{db}
}

func (o *CustomerOrm) GetAllCustomerPaginated(c *gin.Context) (*database.Pagination, error) {
	var mArr []*Customer
	var pagination database.Pagination
	o.db.
		Scopes(database.Paginator(c, &mArr, []string{"Credentials"}, &pagination)).
		Order("created_at DESC").
		Find(&mArr)
	pagination.Data = &mArr
	return &pagination, nil
}

func (o *CustomerOrm) GetOneByID(id uuid.UUID) (Customer Customer, err error) {
	result := o.db.Model(&Customer).
		Where("credentials_id = ?", id).
		Preload(clause.Associations).
		First(&Customer)
	return Customer, result.Error
}

func (o *CustomerOrm) GetOneByCustomerName(Customername string) (m Customer, err error) {
	result := o.db.Model(&m).Where("name = ?", Customername).First(&m)

	if m.Status == "inactive" {
		return m, errors.New("Customer has no company related. please ask your admin for verification")
	}

	return m, result.Error
}

func (o *CustomerOrm) InsertCustomer(p Customer, tx *gorm.DB) (err error) {
	result := tx.Model(&p).Create(&p)
	return result.Error
}

func (o *CustomerOrm) DeleteCustomer(id uuid.UUID, tx *gorm.DB) (err error) {
	result := tx.Model(&Customer{}).Delete(&Customer{}, id)
	return result.Error
}
