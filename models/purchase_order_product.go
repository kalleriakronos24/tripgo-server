package models

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	database "gitlab.com/odma1/odma-be/db"
	masterModels "gitlab.com/odma1/odma-be/models/master"
	"gitlab.com/odma1/odma-be/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PurchaseOrderProductOrm struct {
	db *gorm.DB
}

type PurchaseOrderProduct struct {
	ID         uuid.UUID `gorm:"index:id,unique;type:uuid;default:gen_random_uuid();" json:"id"`
	Quantity   float32   `json:"quantity,omitempty" gorm:"not null"`
	VATRate    int16     `json:"vatRate,omitempty" gorm:"not null"`
	SubTotal   float64   `json:"subTotal,omitempty" gorm:"not null;"`
	GrandTotal float64   `json:"grandTotal,omitempty" gorm:"not null;"`

	PurchaseOrderID uuid.UUID            `json:"purchaseOrderId" gorm:"type:uuid;not null;default:NULL;"`
	PurchaseOrder   *PurchaseOrder       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:PurchaseOrderID;references:ID" json:"purchaseOrder"`
	ProductID       uuid.UUID            `json:"productId" gorm:"type:uuid;not null;default:NULL;"`
	Product         *Product             `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:ProductID;references:ID" json:"product"`
	ClientID        uuid.UUID            `json:"clientId" gorm:"type:uuid;not null;default:NULL;"`
	Client          *masterModels.Client `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:ClientID;references:ID" json:"client"`

	PurchaseOrderProductCreatedBy uuid.UUID          `json:"createdBy" gorm:"type:uuid;not null;default:NULL;"`
	PurchaseOrderProductUpdatedBy uuid.UUID          `json:"updatedBy" gorm:"type:uuid;default:NULL;"`
	CreatedByUser                 *masterModels.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:PurchaseOrderProductCreatedBy;references:ID" json:"createdByUser"`
	UpdatedByUser                 *masterModels.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:PurchaseOrderProductUpdatedBy;references:ID" json:"updatedByUser"`

	types.DefaultModelProperty
}

type PurchaseOrderProductModelAction interface {
	GetAllPurchaseOrderProduct(userId uuid.UUID) (m []PurchaseOrderProduct, err error)
	GetAllPurchaseOrderProductPaginated(c *gin.Context, userId uuid.UUID) (*database.Pagination, error)
	GetOnePurchaseOrderProductByID(id uuid.UUID) (m PurchaseOrderProduct, err error)
	GetOnePurchaseOrderProductByPurchaseOrderId(purchaseOrderId uuid.UUID) (m PurchaseOrderProduct, err error)
	InsertPurchaseOrderProduct(p PurchaseOrderProduct) (err error)
	UpdatePurchaseOrderProduct(id uuid.UUID, p PurchaseOrderProduct) (err error)
}

func NewPurchaseOrderProductAction(db *gorm.DB) PurchaseOrderProductModelAction {
	return &PurchaseOrderProductOrm{db}
}

func (o *PurchaseOrderProductOrm) GetAllPurchaseOrderProduct(userId uuid.UUID) (m []PurchaseOrderProduct, err error) {
	result := o.db.Model(&m).
		Preload("CreatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.
				Select([]string{"ID", "Name", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt"}).
				First(&masterModels.User{}, userId)
		}).
		Preload("UpdatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedAt", "UpdatedAt"})
		}).
		Preload("PurchaseOrder").
		Preload("Product").
		Preload("Client").
		Find(&m)
	return m, result.Error
}

func (o *PurchaseOrderProductOrm) GetAllPurchaseOrderProductPaginated(c *gin.Context, userId uuid.UUID) (*database.Pagination, error) {

	var mArr []*PurchaseOrderProduct
	var pagination database.Pagination

	o.db.
		Scopes(database.Paginator(c, &mArr, []string{"CreatedByUser", "UpdatedByUser", "PurchaseOrder", "Product", "Client"}, &pagination)).
		Where("purchase_order_product_created_by", userId).
		Find(&mArr)
	pagination.Data = &mArr
	return &pagination, nil
}

func (o *PurchaseOrderProductOrm) GetOnePurchaseOrderProductByID(id uuid.UUID) (m PurchaseOrderProduct, err error) {
	result := o.db.Model(&m).
		Preload("CreatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt"})
		}).
		Preload("UpdatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt"})
		}).
		Preload("PurchaseOrder").
		Preload("Product").
		Preload("Client").
		First(&m, id)
	return m, result.Error
}

func (o *PurchaseOrderProductOrm) GetOnePurchaseOrderProductByPurchaseOrderId(purchaseOrderId uuid.UUID) (m PurchaseOrderProduct, err error) {
	result := o.db.Model(&m).Where("purchase_order_id = ?", purchaseOrderId).First(&m)
	return m, result.Error
}

func (o *PurchaseOrderProductOrm) InsertPurchaseOrderProduct(p PurchaseOrderProduct) (err error) {
	result := o.db.Model(&p).Omit(clause.Associations).Create(&p)
	return result.Error
}

func (o *PurchaseOrderProductOrm) UpdatePurchaseOrderProduct(id uuid.UUID, p PurchaseOrderProduct) (err error) {
	result := o.db.Model(&p).Where("id = ?", id).Updates(&p)
	return result.Error
}
