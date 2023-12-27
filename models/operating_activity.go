package models

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	database "gitlab.com/odma1/odma-be/db"
	masterModels "gitlab.com/odma1/odma-be/models/master"
	"gitlab.com/odma1/odma-be/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
)

type operatingActivityOrm struct {
	db *gorm.DB
}

type DynamicQueryOperatingActivityParam struct {
	*OperatingActivity
}

type OperatingActivity struct {
	ID                    uuid.UUID            `gorm:"index:id,unique;type:uuid;default:gen_random_uuid();" json:"id"`
	TaxInvoiceNumber      string               `json:"taxInvoiceNumber,omitempty" gorm:"not null;unique"`
	DeliveryReceiptNumber string               `json:"deliveryReceiptNumber,omitempty" gorm:"not null;unique"`
	Status                string               `json:"status,omitempty" gorm:"not null;default:ongoing"`
	ClientID              uuid.UUID            `json:"clientId" gorm:"type:uuid;not null;default:NULL;"`
	Client                *masterModels.Client `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:ClientID;references:ID" json:"client"`

	ProductHistory           []*ProductHistory           `json:"productHistory,omitempty"`
	Quotation                []*Quotation                `json:"quotation,omitempty"`
	Payment                  []*Payment                  `json:"payment,omitempty"`
	OperatingActivityProduct []*OperatingActivityProduct `json:"operatingActivityProduct,omitempty"`
	Invoice                  []*Invoice                  `json:"invoice,omitempty"`
	DeliveryOrder            []*DeliveryOrder            `json:"deliveryOrder,omitempty"`
	PurchaseOrder            []*PurchaseOrder            `json:"purchaseOrder,omitempty"`

	OperatingActivityCreatedBy uuid.UUID          `json:"createdBy" gorm:"type:uuid;not null;default:NULL;"`
	OperatingActivityUpdatedBy uuid.UUID          `json:"updatedBy" gorm:"type:uuid;default:NULL;"`
	CreatedByUser              *masterModels.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:OperatingActivityCreatedBy;references:ID" json:"createdByUser"`
	UpdatedByUser              *masterModels.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:OperatingActivityUpdatedBy;references:ID" json:"updatedByUser"`

	types.DefaultModelProperty
}

type OperatingActivityModelAction interface {
	GetAllOperatingActivity(userId uuid.UUID) (m []OperatingActivity, err error)
	GetAllOperatingActivityPaginated(c *gin.Context, userId uuid.UUID) (*database.Pagination, error)
	GetOneOperatingActivityByID(id uuid.UUID) (m OperatingActivity, err error)
	GetOneOperatingActivityByTaxNumber(taxNumber string) (m OperatingActivity, err error)
	GetOneOperatingActivityByClientID(clientId uuid.UUID) (m []OperatingActivity, err error)

	InsertOperatingActivity(p OperatingActivity) (err error)
	UpdateOperatingActivity(id uuid.UUID, p OperatingActivity) (err error)
	DeleteOperatingActivityByClientID(id uuid.UUID, tx *gorm.DB) (err error)
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
		Preload("Payment").
		Preload("OperatingActivityProduct").
		Find(&m)
	return m, result.Error
}

func (o *operatingActivityOrm) GetAllOperatingActivityPaginated(c *gin.Context, userId uuid.UUID) (*database.Pagination, error) {

	var mArr []*OperatingActivity
	var pagination database.Pagination

	o.db.
		Scopes(database.Paginator(c, &mArr, []string{
			"CreatedByUser",
			"UpdatedByUser",
			"Client",
			"Payment",
			"OperatingActivityProduct"}, &pagination)).
		Where("operating_activity_created_by", userId).
		Find(&mArr)
	pagination.Data = &mArr
	return &pagination, nil
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

func (o *operatingActivityOrm) GetOneOperatingActivityByClientID(clientId uuid.UUID) (m []OperatingActivity, err error) {
	result := o.db.Model(&m).Where("client_id = ?", clientId).Find(&m)
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

func (o *operatingActivityOrm) DeleteOperatingActivityByClientID(id uuid.UUID, tx *gorm.DB) (err error) {
	result := tx.Model(&OperatingActivity{}).Where("client_id = ?", id).Delete(&OperatingActivity{})
	return result.Error
}
