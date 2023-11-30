package models

import (
	"github.com/google/uuid"
	masterModels "gitlab.com/odma1/odma-be/models/master"
	"gitlab.com/odma1/odma-be/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type PurchaseOrderOrm struct {
	db *gorm.DB
}

type PurchaseOrder struct {
	ID             uuid.UUID `gorm:"index:id,unique;type:uuid;default:gen_random_uuid();" json:"id"`
	Number         string    `json:"number,omitempty" gorm:"not null;unique"`
	Type           string    `json:"type,omitempty" gorm:"not null"`
	Recipient      string    `json:"recipient,omitempty" gorm:"not null;"`
	RecipientEmail string    `json:"recipientEmail,omitempty" gorm:"not null;"`
	Date           time.Time `json:"date,omitempty" gorm:"not null"`

	DocumentID          uuid.UUID          `json:"documentId,omitempty" gorm:"not null;"`
	Document            *Document          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:DocumentID;references:ID" json:"document"`
	OperatingActivityID uuid.UUID          `json:"operatingActivityId" gorm:"type:uuid;not null;default:NULL;"`
	OperatingActivity   *OperatingActivity `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:OperatingActivityID;references:ID" json:"operatingActivity"`

	PurchaseOrderCreatedBy uuid.UUID          `json:"createdBy" gorm:"type:uuid;not null;default:NULL;"`
	PurchaseOrderUpdatedBy uuid.UUID          `json:"updatedBy" gorm:"type:uuid;default:NULL;"`
	CreatedByUser          *masterModels.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:PurchaseOrderCreatedBy;references:ID" json:"createdByUser"`
	UpdatedByUser          *masterModels.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:PurchaseOrderUpdatedBy;references:ID" json:"updatedByUser"`

	types.DefaultModelProperty
}

type PurchaseOrderModelAction interface {
	GetAllPurchaseOrder(userId uuid.UUID) (m []PurchaseOrder, err error)
	GetOnePurchaseOrderByID(id uuid.UUID) (m PurchaseOrder, err error)
	GetOnePurchaseOrderByOperatingActivityId(operatingActivityId uuid.UUID) (m PurchaseOrder, err error)

	InsertPurchaseOrder(p PurchaseOrder) (err error)
	UpdatePurchaseOrder(id uuid.UUID, p PurchaseOrder) (err error)
}

func NewPurchaseOrderAction(db *gorm.DB) PurchaseOrderModelAction {
	return &PurchaseOrderOrm{db}
}

func (o *PurchaseOrderOrm) GetAllPurchaseOrder(userId uuid.UUID) (m []PurchaseOrder, err error) {
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
		Preload("Document").
		Find(&m)
	return m, result.Error
}

func (o *PurchaseOrderOrm) GetOnePurchaseOrderByID(id uuid.UUID) (m PurchaseOrder, err error) {
	result := o.db.Model(&m).
		Preload("CreatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt"})
		}).
		Preload("UpdatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt"})
		}).
		Preload("OperatingActivity").
		Preload("Document").
		First(&m, id)
	return m, result.Error
}

func (o *PurchaseOrderOrm) GetOnePurchaseOrderByOperatingActivityId(operatingActivityId uuid.UUID) (m PurchaseOrder, err error) {
	result := o.db.Model(&m).Where("operating_activity_id = ?", operatingActivityId).First(&m)
	return m, result.Error
}

func (o *PurchaseOrderOrm) InsertPurchaseOrder(p PurchaseOrder) (err error) {
	result := o.db.Model(&p).Omit(clause.Associations).Create(&p)
	return result.Error
}

func (o *PurchaseOrderOrm) UpdatePurchaseOrder(id uuid.UUID, p PurchaseOrder) (err error) {
	result := o.db.Model(&p).Where("id = ?", id).Updates(&p)
	return result.Error
}
