package master

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	database "github.com/kalleriakronos24/khaimal-group/db"
	"github.com/kalleriakronos24/khaimal-group/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type InternalOrm struct {
	db *gorm.DB
}

type Internal struct {
	ID       uuid.UUID `gorm:"index:id,unique;type:uuid;default:gen_random_uuid();" json:"id"`
	Name     string    `json:"name" gorm:"not null" binding:"required"`
	Email    string    `gorm:"email:id,unique" json:",omitempty" binding:"required"`
	Password string    `json:"password,omitempty" binding:"required" gorm:"not null"`
	Phone    string    `json:"phone,omitempty" gorm:"default:NULL"`
	Status   string    `json:"status,omitempty" binding:"required" gorm:"not null;default:active;"`

	CredentialsID uuid.UUID `json:"credentialsId" gorm:"type:uuid;not null"`
	CompanyID     uuid.UUID `json:"companyId" gorm:"type:uuid;default:NULL"`
	CreatedBy     uuid.UUID `json:"createdBy,omitempty" gorm:"type:uuid;default:NULL"`
	UpdatedBy     uuid.UUID `json:"updatedBy,omitempty" gorm:"type:uuid;default:NULL"`

	Credentials *Credentials `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:CredentialsID;references:ID" json:"credentials,omitempty"`
	Company     *Company     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:CompanyID;references:ID" json:"company,omitempty"`

	types.DefaultModelProperty
}

type InternalModelAction interface {
	GetOneByID(id uuid.UUID) (m Internal, err error)
	GetOneByInternalName(Internalname string) (m Internal, err error)
	GetOneByEmail(email string) (m Internal, err error)
	GetAllInternalPaginated(c *gin.Context, InternalId uuid.UUID) (*database.Pagination, error)

	InsertInternal(p Internal, tx *gorm.DB) (err error)

	UpdateInternalToInactive(id uuid.UUID, tx *gorm.DB) (err error)
	RemoveInternalFromCompany(id uuid.UUID, tx *gorm.DB) (err error)
	DeleteInternal(id uuid.UUID, tx *gorm.DB) (err error)
}

func NewInternalAction(db *gorm.DB) InternalModelAction {
	return &InternalOrm{db}
}

func (o *InternalOrm) GetAllInternalPaginated(c *gin.Context, InternalId uuid.UUID) (*database.Pagination, error) {
	var mArr []*Internal
	var pagination database.Pagination
	o.db.
		Scopes(database.Paginator(c, &mArr, []string{"CreatedBy", "UpdatedBy", "Company"}, &pagination)).
		Where("created_by", InternalId).
		Find(&mArr)
	pagination.Data = &mArr
	return &pagination, nil
}

func (o *InternalOrm) GetOneByID(id uuid.UUID) (Internal Internal, err error) {
	result := o.db.Model(&Internal).
		Where("id = ?", id).
		Preload(clause.Associations).
		First(&Internal)
	return Internal, result.Error
}

func (o *InternalOrm) GetOneByEmail(email string) (m Internal, err error) {
	result := o.db.Model(&m).Where("email = ?", email).First(&m)
	return m, result.Error
}

func (o *InternalOrm) GetOneByInternalName(Internalname string) (m Internal, err error) {
	result := o.db.Model(&m).Where("Internalname = ?", Internalname).First(&m)

	if m.Status == "inactive" {
		return m, errors.New("Internal has no company related. please ask your admin for verification")
	}

	return m, result.Error
}

func (o *InternalOrm) UpdateInternalToInactive(id uuid.UUID, tx *gorm.DB) (err error) {
	result := tx.Model(&Internal{}).Where("created_by = ? AND role != 'superadmin'", id).Update("status", "inactive")
	return result.Error
}

func (o *InternalOrm) RemoveInternalFromCompany(id uuid.UUID, tx *gorm.DB) (err error) {
	result := tx.Model(&Internal{}).Where("id", id).Update("company_id", nil)
	return result.Error
}

func (o *InternalOrm) InsertInternal(p Internal, tx *gorm.DB) (err error) {
	result := tx.Model(&p).Create(&p)
	return result.Error
}

func (o *InternalOrm) DeleteInternal(id uuid.UUID, tx *gorm.DB) (err error) {
	result := tx.Model(&Internal{}).Delete(&Internal{}, id)
	return result.Error
}

// TENTANTS
func (o *InternalOrm) InsertInternalOwner(p Internal) (err error) {
	fmt.Printf("%v", p)
	result := o.db.Model(&p).Create(&p)
	return result.Error
}

func (o *InternalOrm) InsertInternalAdmin(p Internal) (err error) {
	fmt.Printf("%v", p)
	result := o.db.Model(&p).Create(&p)
	return result.Error
}
