package models

import (
	"github.com/google/uuid"
	masterModels "gitlab.com/odma1/odma-be/models/master"
	"gitlab.com/odma1/odma-be/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OperatingActivityProductOrm struct {
	db *gorm.DB
}

type OperatingActivityProduct struct {
	ID         uuid.UUID `gorm:"index:id,unique;type:uuid;default:gen_random_uuid();" json:"id"`
	Quantity   float32   `json:"quantity,omitempty" gorm:"not null"`
	VATRate    int16     `json:"vatRate,omitempty" gorm:"not null"`
	SubTotal   float64   `json:"subTotal,omitempty" gorm:"not null;"`
	GrandTotal float64   `json:"grandTotal,omitempty" gorm:"not null;"`

	OperatingActivityID uuid.UUID          `json:"operatingActivityId" gorm:"type:uuid;not null;default:NULL;"`
	OperatingActivity   *OperatingActivity `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:OperatingActivityID;references:ID" json:"operatingActivity"`
	ProductID           uuid.UUID          `json:"productId" gorm:"type:uuid;not null;default:NULL;"`
	Product             *Product           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ProductID;references:ID" json:"product"`

	OperatingActivityProductCreatedBy uuid.UUID          `json:"createdBy" gorm:"type:uuid;not null;default:NULL;"`
	OperatingActivityProductUpdatedBy uuid.UUID          `json:"updatedBy" gorm:"type:uuid;default:NULL;"`
	CreatedByUser                     *masterModels.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:OperatingActivityProductCreatedBy;references:ID" json:"createdByUser"`
	UpdatedByUser                     *masterModels.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:OperatingActivityProductUpdatedBy;references:ID" json:"updatedByUser"`

	types.DefaultModelProperty
}

type OperatingActivityProductModelAction interface {
	GetAllOperatingActivityProduct(userId uuid.UUID) (m []OperatingActivityProduct, err error)
	GetOneOperatingActivityProductByID(id uuid.UUID) (m OperatingActivityProduct, err error)
	GetOneOperatingActivityProductByOperatingActivityId(operatingActivityId uuid.UUID) (m OperatingActivityProduct, err error)

	InsertOperatingActivityProduct(p OperatingActivityProduct) (err error)
	UpdateOperatingActivityProduct(id uuid.UUID, p OperatingActivityProduct) (err error)
}

func NewOperatingActivityProductAction(db *gorm.DB) OperatingActivityProductModelAction {
	return &OperatingActivityProductOrm{db}
}

func (o *OperatingActivityProductOrm) GetAllOperatingActivityProduct(userId uuid.UUID) (m []OperatingActivityProduct, err error) {
	result := o.db.Model(&m).
		Preload("CreatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.
				Select([]string{"ID", "Name", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt"}).
				First(&masterModels.User{}, userId)
		}).
		Preload("UpdatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedAt", "UpdatedAt"})
		}).
		Preload("OperatingActivity").
		Preload("Product").
		Find(&m)
	return m, result.Error
}

func (o *OperatingActivityProductOrm) GetOneOperatingActivityProductByID(id uuid.UUID) (m OperatingActivityProduct, err error) {
	result := o.db.Model(&m).
		Preload("CreatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt"})
		}).
		Preload("UpdatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt"})
		}).
		Preload("OperatingActivity").
		Preload("Product").
		First(&m, id)
	return m, result.Error
}

func (o *OperatingActivityProductOrm) GetOneOperatingActivityProductByOperatingActivityId(operatingActivityId uuid.UUID) (m OperatingActivityProduct, err error) {
	result := o.db.Model(&m).Where("operating_activity_id = ?", operatingActivityId).First(&m)
	return m, result.Error
}

func (o *OperatingActivityProductOrm) InsertOperatingActivityProduct(p OperatingActivityProduct) (err error) {
	result := o.db.Model(&p).Omit(clause.Associations).Create(&p)
	return result.Error
}

func (o *OperatingActivityProductOrm) UpdateOperatingActivityProduct(id uuid.UUID, p OperatingActivityProduct) (err error) {
	result := o.db.Model(&p).Where("id = ?", id).Updates(&p)
	return result.Error
}
