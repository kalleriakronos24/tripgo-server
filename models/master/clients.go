package master

import (
	"github.com/google/uuid"
	"gitlab.com/odma1/odma-be/types"
	"gorm.io/gorm"
	"strings"
)

type clientOrm struct {
	db *gorm.DB
}

type Client struct {
	ID          uuid.UUID `json:"id" gorm:"index:id,unique;type:uuid;default:gen_random_uuid();"`
	Name        string    `json:"name" gorm:"not null;default:NULL" `
	PhoneNumber string    `json:"phoneNumber,omitempty" gorm:"not null;default:NULL"`
	Email       string    `gorm:"index:email,unique;not null;default:NULL" json:"email,omitempty"`
	Address     string    `json:"address,omitempty" gorm:"default:NULL"`

	CompanyID uuid.UUID `json:"companyId" gorm:"type:uuid;default:NULL"`
	Company   *Company  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:CompanyID;references:ID" json:"company"`

	ClientCreatedBy uuid.UUID `json:"createdBy" gorm:"type:uuid;not null;default:NULL"`
	ClientUpdatedBy uuid.UUID `json:"updatedBy" gorm:"type:uuid;default:NULL"`
	CreatedByUser   *User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ClientCreatedBy;references:ID" json:"createdByUser"`
	UpdatedByUser   *User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ClientUpdatedBy;references:ID" json:"updatedByUser"`

	types.DefaultModelProperty
}

type ClientModelAction interface {
	GetAllClient(userId uuid.UUID) (m []Client, err error)
	GetOneClientByID(id uuid.UUID) (m Client, err error)
	GetOneClientByName(name string) (m Client, err error)
	GetOneClientByEmail(email string) (m Client, err error)

	InsertClient(p Client) (err error)

	UpdateClient(id uuid.UUID, p Client) (err error)
}

func NewClientAction(db *gorm.DB) ClientModelAction {
	return &clientOrm{db}
}

func (o *clientOrm) GetAllClient(userId uuid.UUID) (m []Client, err error) {
	result := o.db.Model(&m).
		Preload("CreatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.
				Select([]string{"ID", "Name", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt"}).
				First(&User{}, userId)
		}).
		Preload("UpdatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedBy", "UpdatedAt"})
		}).
		Preload("Company", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedAt", "UpdatedAt"})
		}).
		Find(&m)
	return m, result.Error
}

func (o *clientOrm) GetOneClientByID(id uuid.UUID) (m Client, err error) {
	result := o.db.Model(&m).
		Preload("CreatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt"})
		}).
		Preload("UpdatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedBy", "UpdatedAt"})
		}).
		Preload("Company", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedAt", "UpdatedAt"})
		}).
		Where("id = ?", id).
		First(&m)
	return m, result.Error
}

func (o *clientOrm) GetOneClientByName(name string) (m Client, err error) {
	result := o.db.Model(&m).Where("lower(name) = ?", strings.ToLower(name)).First(&m)
	return m, result.Error
}

func (o *clientOrm) GetOneClientByEmail(email string) (m Client, err error) {
	result := o.db.Model(&m).Where("email = ?", email).First(&m)
	return m, result.Error
}

func (o *clientOrm) InsertClient(p Client) (err error) {
	result := o.db.Model(&p).Create(&p)
	return result.Error
}

func (o *clientOrm) UpdateClient(id uuid.UUID, p Client) (err error) {
	result := o.db.Model(&p).Where("id = ?", id).Updates(&p)
	return result.Error
}
