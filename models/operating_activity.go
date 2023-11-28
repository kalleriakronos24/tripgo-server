package models

import (
	"github.com/google/uuid"
	masterModels "gitlab.com/odma1/odma-be/models/master"
	"gitlab.com/odma1/odma-be/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
)

type operatingActivityOrm struct {
	db *gorm.DB
}

type OperatingActivity struct {
	ID                    uuid.UUID `gorm:"index:id,unique;type:uuid;default:gen_random_uuid();" json:"id"`
	TaxInvoiceNumber      string    `json:"taxInvoiceNumber,omitempty" gorm:"not null"`
	DeliveryReceiptNumber string    `json:"deliveryReceiptNumber,omitempty" gorm:"not null"`
	Status                string    `json:"status,omitempty" gorm:"not null;default:ongoing"`

	ClientID       uuid.UUID            `json:"clientId" gorm:"type:uuid;not null;default:NULL;"`
	Client         *masterModels.Client `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ClientID;references:ID" json:"client"`
	ProductHistory []*ProductHistory    `json:"productHistory,omitempty"`
	Quotation      []*Quotation         `json:"quotation,omitempty"`
	Payment        []*Payment           `json:"payment,omitempty"`

	OperatingActivityCreatedBy uuid.UUID          `json:"createdBy" gorm:"type:uuid;not null;default:NULL;"`
	OperatingActivityUpdatedBy uuid.UUID          `json:"updatedBy" gorm:"type:uuid;default:NULL;"`
	CreatedByUser              *masterModels.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:OperatingActivityCreatedBy;references:ID" json:"createdByUser"`
	UpdatedByUser              *masterModels.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:OperatingActivityUpdatedBy;references:ID" json:"updatedByUser"`

	types.DefaultModelProperty
}

type OperatingActivityModelAction interface {
	GetAllOperatingActivity(userId uuid.UUID) (m []OperatingActivity, err error)
	GetOneOperatingActivityByID(id uuid.UUID) (m OperatingActivity, err error)
	GetOneOperatingActivityByTaxNumber(taxNumber string) (m OperatingActivity, err error)

	InsertOperatingActivity(p OperatingActivity) (err error)
	UpdateOperatingActivity(id uuid.UUID, p OperatingActivity) (err error)
}

func NewOperatingActivityAction(db *gorm.DB) OperatingActivityModelAction {
	return &operatingActivityOrm{db}
}

func (o *operatingActivityOrm) GetAllOperatingActivity(userId uuid.UUID) (m []OperatingActivity, err error) {
	result := o.db.Model(&m).
		Preload("CreatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.
				Select([]string{"ID", "Name", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt"}).
				First(&masterModels.User{}, userId)
		}).
		Preload("UpdatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedAt", "UpdatedAt"})
		}).
		Preload("Client", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedAt", "UpdatedAt"})
		}).
		Preload("Quotation").
		Preload("ProductHistory").
		Find(&m)
	return m, result.Error
}

func (o *operatingActivityOrm) GetOneOperatingActivityByID(id uuid.UUID) (m OperatingActivity, err error) {
	result := o.db.Model(&m).
		Preload("CreatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt"})
		}).
		Preload("UpdatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt"})
		}).
		Preload("Client", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedAt", "UpdatedAt"})
		}).
		First(&m, id)
	return m, result.Error
}

func (o *operatingActivityOrm) GetOneOperatingActivityByTaxNumber(taxNumber string) (m OperatingActivity, err error) {
	result := o.db.Model(&m).Where("lower(tax_invoice_number) = ?", strings.ToLower(taxNumber)).First(&m)
	return m, result.Error
}

func (o *operatingActivityOrm) InsertOperatingActivity(p OperatingActivity) (err error) {
	result := o.db.Model(&p).Omit(clause.Associations).Create(&p)
	return result.Error
}

func (o *operatingActivityOrm) UpdateOperatingActivity(id uuid.UUID, p OperatingActivity) (err error) {
	result := o.db.Model(&p).Where("id = ?", id).Updates(&p)
	return result.Error
}
