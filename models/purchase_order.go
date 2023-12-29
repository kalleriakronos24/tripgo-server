package models

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	database "gitlab.com/odma1/odma-be/db"
	masterModels "gitlab.com/odma1/odma-be/models/master"
	"gitlab.com/odma1/odma-be/types"
	"gitlab.com/odma1/odma-be/utils"
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

	DocumentID           uuid.UUID               `json:"documentId,omitempty" gorm:"not null;"`
	Document             *Document               `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:DocumentID;references:ID" json:"document"`
	OperatingActivityID  uuid.UUID               `json:"operatingActivityId" gorm:"type:uuid;not null;default:NULL;"`
	OperatingActivity    *OperatingActivity      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:OperatingActivityID;references:ID" json:"operatingActivity"`
	PurchaseOrderProduct []*PurchaseOrderProduct `json:"purchaseOrderProduct,omitempty"`

	PurchaseOrderCreatedBy uuid.UUID          `json:"createdBy" gorm:"type:uuid;not null;default:NULL;"`
	PurchaseOrderUpdatedBy uuid.UUID          `json:"updatedBy" gorm:"type:uuid;default:NULL;"`
	CreatedByUser          *masterModels.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:PurchaseOrderCreatedBy;references:ID" json:"createdByUser"`
	UpdatedByUser          *masterModels.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:PurchaseOrderUpdatedBy;references:ID" json:"updatedByUser"`

	types.DefaultModelProperty
}

type CustomResponsePurchaseOrder struct {
	ID             uuid.UUID `json:"id"`
	Number         string    `json:"number,omitempty"`
	Type           string    `json:"type,omitempty"`
	Recipient      string    `json:"recipient,omitempty"`
	RecipientEmail string    `json:"recipientEmail,omitempty"`
	Date           time.Time `json:"date,omitempty"`

	OperatingActivityID uuid.UUID          `json:"operatingActivityId"`
	DocumentID          uuid.UUID          `json:"documentId,omitempty"`
	Document            *Document          `json:"document"`
	OperatingActivity   *OperatingActivity `json:"operatingActivity"`
	Product             []any              `json:"product"`

	CreatedByUser *masterModels.User `json:"createdByUser"`
	UpdatedByUser *masterModels.User `json:"updatedByUser"`

	types.DefaultModelProperty
}

type PurchaseOrderModelAction interface {
	GetAllPurchaseOrder(userId uuid.UUID) (m []PurchaseOrder, err error)
	GetAllPurchaseOrderPaginated(c *gin.Context, userId uuid.UUID) (*database.Pagination, error)
	GetOnePurchaseOrderByID(id uuid.UUID) (m CustomResponsePurchaseOrder, err error)
	GetOnePurchaseOrderByOperatingActivityId(operatingActivityId uuid.UUID) (m PurchaseOrder, err error)

	InsertPurchaseOrder(p PurchaseOrder) (err error)
	UpdatePurchaseOrder(id uuid.UUID, p PurchaseOrder) (err error)
	DeletePurchaseOrderByOptActID(id uuid.UUID, tx *gorm.DB) (err error)
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

func (o *PurchaseOrderOrm) GetAllPurchaseOrderPaginated(c *gin.Context, userId uuid.UUID) (*database.Pagination, error) {

	var mArr []*PurchaseOrder
	var pagination database.Pagination

	o.db.
		Scopes(database.Paginator(c, &mArr, []string{"CreatedByUser", "UpdatedByUser", "Document", "OperatingActivity"}, &pagination)).
		Where("purchase_order_created_by", userId).
		Find(&mArr)
	pagination.Data = &mArr
	return &pagination, nil
}

func (o *PurchaseOrderOrm) GetOnePurchaseOrderByID(id uuid.UUID) (m CustomResponsePurchaseOrder, err error) {
	purchaseOrderModel := PurchaseOrder{}
	result := o.db.Model(&purchaseOrderModel).
		Preload("CreatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt"})
		}).
		Preload("UpdatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt"})
		}).
		Preload("OperatingActivity").
		Preload("Document").
		First(&purchaseOrderModel, id)

	if purchaseOrderModel.Type == "in" {

		var operatingActivityProduct []OperatingActivityProduct
		err := o.db.Model(&operatingActivityProduct).Preload("Product").Where("operating_activity_id = ?", purchaseOrderModel.OperatingActivityID).Find(&operatingActivityProduct)

		if err.Error != nil {
			return m, err.Error
		}

		converted := utils.UnpackArray(operatingActivityProduct)

		m := CustomResponsePurchaseOrder{
			ID:                purchaseOrderModel.ID,
			Number:            purchaseOrderModel.Number,
			Type:              purchaseOrderModel.Type,
			Recipient:         purchaseOrderModel.Recipient,
			RecipientEmail:    purchaseOrderModel.RecipientEmail,
			Date:              purchaseOrderModel.Date,
			DocumentID:        purchaseOrderModel.DocumentID,
			Document:          purchaseOrderModel.Document,
			OperatingActivity: purchaseOrderModel.OperatingActivity,
			Product:           converted,
			CreatedByUser:     purchaseOrderModel.CreatedByUser,
			UpdatedByUser:     purchaseOrderModel.UpdatedByUser,
		}
		return m, result.Error
	}

	if purchaseOrderModel.Type == "out" {
		var purchaseOrderProductModel []PurchaseOrderProduct
		err := o.db.Model(&purchaseOrderProductModel).
			Preload("Product").
			Preload("Client").
			Where("purchase_order_id = ?", purchaseOrderModel.ID).
			Find(&purchaseOrderProductModel)

		if err.Error != nil {
			return m, err.Error
		}

		converted := utils.UnpackArray(purchaseOrderProductModel)

		m := CustomResponsePurchaseOrder{
			ID:                purchaseOrderModel.ID,
			Number:            purchaseOrderModel.Number,
			Type:              purchaseOrderModel.Type,
			Recipient:         purchaseOrderModel.Recipient,
			RecipientEmail:    purchaseOrderModel.RecipientEmail,
			Date:              purchaseOrderModel.Date,
			DocumentID:        purchaseOrderModel.DocumentID,
			Document:          purchaseOrderModel.Document,
			OperatingActivity: purchaseOrderModel.OperatingActivity,
			Product:           converted,
			CreatedByUser:     purchaseOrderModel.CreatedByUser,
			UpdatedByUser:     purchaseOrderModel.UpdatedByUser,
		}
		return m, result.Error
	}
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

func (o *PurchaseOrderOrm) DeletePurchaseOrderByOptActID(id uuid.UUID, tx *gorm.DB) (err error) {
	result := tx.Model(&PurchaseOrder{}).Where("operating_activity_id = ?", id).Delete(&PurchaseOrder{})
	return result.Error
}
