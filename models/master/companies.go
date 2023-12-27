package master

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	database "gitlab.com/odma1/odma-be/db"
	"gitlab.com/odma1/odma-be/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
)

type companyOrm struct {
	db *gorm.DB
}

type Company struct {
	ID                uuid.UUID `gorm:"index:id,unique;type:uuid;default:gen_random_uuid();" json:"id"`
	Name              string    `json:"name" gorm:"not null"`
	PhoneNumber       string    `json:"phoneNumber,omitempty" gorm:"not null"`
	Email             string    `gorm:"index:email,unique;not null" json:"email,omitempty"`
	Address           string    `json:"address,omitempty" gorm:"default:NULL"`
	PICName           string    `json:"picName,omitempty" gorm:"not null"`
	PICDesignation    string    `json:"picDesignation,omitempty" gorm:"not null"`
	BankAccountName   string    `json:"bankAccountName,omitempty" gorm:"not null"`
	BankAccountNumber int       `json:"bankAccountNumber,omitempty" gorm:"not null;default:0"`

	Client           []*Client `json:"client,omitempty"`
	User             []*User   `json:"user,omitempty"`
	CompanyCreatedBy uuid.UUID `json:"createdBy" gorm:"type:uuid;not null;default:NULL;"`
	CompanyUpdatedBy uuid.UUID `json:"updatedBy" gorm:"type:uuid;default:NULL;"`
	CreatedByUser    *User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:CompanyCreatedBy;references:ID" json:"createdByUser"`
	UpdatedByUser    *User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:CompanyUpdatedBy;references:ID" json:"updatedByUser"`

	types.DefaultModelProperty
}

type CompanyModelAction interface {
	GetAllCompany(userId uuid.UUID) (m []Company, err error)
	GetAllCompanyPaginated(c *gin.Context, userId uuid.UUID) (*database.Pagination, error)
	GetOneCompanyByID(id uuid.UUID) (m Company, err error)
	GetOneCompanyByName(name string) (m Company, err error)
	GetOneCompanyByEmail(email string) (m Company, err error)
	GetOneCompanyByUserID(userId uuid.UUID) (m Company, err error)

	InsertCompany(p Company) (err error)
	UpdateCompany(id uuid.UUID, p Company) (err error)
	DeleteCompany(id uuid.UUID, tx *gorm.DB) (err error)
}

func NewCompanyAction(db *gorm.DB) CompanyModelAction {
	return &companyOrm{db}
}

func (o *companyOrm) GetAllCompanyPaginated(c *gin.Context, userId uuid.UUID) (*database.Pagination, error) {

	var mArr []*Company
	var pagination database.Pagination

	o.db.
		Scopes(database.Paginator(c, &mArr, []string{"CreatedByUser", "UpdatedByUser", "Client"}, &pagination)).
		Where("company_created_by", userId).
		Find(&mArr)
	pagination.Data = &mArr
	return &pagination, nil
}

func (o *companyOrm) GetOneCompanyByUserID(userId uuid.UUID) (m Company, err error) {
	user := User{
		ID: userId,
	}
	userResult := o.db.Model(&user).First(&user)
	if userResult.Error != nil {
		return m, userResult.Error
	}
	result := o.db.Model(&m).
		Preload("CreatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.
				Select([]string{"ID", "Name", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt"})
		}).
		Preload("UpdatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedBy", "UpdatedAt"})
		}).
		Preload("Client").
		Where("company_created_by = ? ", user.ID).First(&m)
	return m, result.Error
}

func (o *companyOrm) GetAllCompany(userId uuid.UUID) (m []Company, err error) {
	result := o.db.Model(&m).
		Preload("CreatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.
				Select([]string{"ID", "Name", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt"}).
				First(&User{}, userId)
		}).
		Preload("UpdatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedBy", "UpdatedAt"})
		}).
		Preload("Client").
		Find(&m)
	return m, result.Error
}

func (o *companyOrm) GetOneCompanyByID(id uuid.UUID) (m Company, err error) {
	result := o.db.Model(&m).
		Preload("CreatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt"})
		}).
		Preload("UpdatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt"})
		}).
		Preload("Client").
		First(&m, id)

	fmt.Printf("%v", &m)

	return m, result.Error
}

func (o *companyOrm) GetOneCompanyByName(name string) (m Company, err error) {
	result := o.db.Model(&m).Where("lower(name) = ?", strings.ToLower(name)).First(&m)
	return m, result.Error
}

func (o *companyOrm) GetOneCompanyByEmail(email string) (m Company, err error) {
	result := o.db.Model(&m).Where("email = ?", email).First(&m)
	return m, result.Error
}

func (o *companyOrm) InsertCompany(p Company) (err error) {
	result := o.db.Model(&p).Omit(clause.Associations).Create(&p)
	return result.Error
}

func (o *companyOrm) UpdateCompany(id uuid.UUID, p Company) (err error) {
	result := o.db.Model(&p).Where("id = ?", id).Updates(&p)
	return result.Error
}

func (o *companyOrm) DeleteCompany(id uuid.UUID, tx *gorm.DB) (err error) {
	result := tx.Model(&Company{}).Delete(&Company{}, id)
	return result.Error
}
