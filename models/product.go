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

type productOrm struct {
	db *gorm.DB
}

type Product struct {
	ID          uuid.UUID `gorm:"index:id,unique;type:uuid;default:gen_random_uuid();" json:"id"`
	Name        string    `json:"name,omitempty" gorm:"not null"`
	UnitPrice   float64   `json:"unitPrice,omitempty" gorm:"not null"`
	Packaging   string    `json:"packaging,omitempty" gorm:"not null;"`
	Stock       float64   `json:"stock,omitempty" gorm:"not null;default:0"`
	Note        string    `json:"note,omitempty"`
	VATIncluded *bool     `json:"VATIncluded,omitempty" gorm:"default:true;"`

	CompanyID            uuid.UUID               `json:"companyId" gorm:"type:uuid;not null;default:NULL;"`
	Company              *masterModels.Company   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:CompanyID;references:ID" json:"company"`
	ProductHistory       []*ProductHistory       `json:"productHistory,omitempty"`
	PurchaseOrderProduct []*PurchaseOrderProduct `json:"purchaseOrderProduct,omitempty"`

	ProductCreatedBy uuid.UUID          `json:"createdBy" gorm:"type:uuid;not null;default:NULL;"`
	ProductUpdatedBy uuid.UUID          `json:"updatedBy" gorm:"type:uuid;default:NULL;"`
	CreatedByUser    *masterModels.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ProductCreatedBy;references:ID" json:"createdByUser"`
	UpdatedByUser    *masterModels.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ProductUpdatedBy;references:ID" json:"updatedByUser"`

	types.DefaultModelProperty
}

//

type ProductModelAction interface {
	GetAllProduct(userId uuid.UUID) (m []Product, err error)
	GetAllProductPaginated(c *gin.Context, userId uuid.UUID) (*database.Pagination, error)
	GetOneProductByID(id uuid.UUID) (m Product, err error)
	GetOneProductByName(taxNumber string) (m Product, err error)

	InsertProduct(p Product) (err error)
	UpdateProduct(id uuid.UUID, p Product) (err error)
	DeleteProduct(id uuid.UUID, tx *gorm.DB) (err error)
}

func NewProductAction(db *gorm.DB) ProductModelAction {
	return &productOrm{db}
}

func (o *productOrm) GetAllProduct(userId uuid.UUID) (m []Product, err error) {
	result := o.db.Model(&m).
		Preload("CreatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.
				Select([]string{"ID", "Name", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt"}).
				First(&masterModels.User{}, userId)
		}).
		Preload("UpdatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedAt", "UpdatedAt"})
		}).
		Preload("Company", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedAt", "UpdatedAt"})
		}).
		Preload("ProductHistory").
		Find(&m)
	return m, result.Error
}

func (o *productOrm) GetAllProductPaginated(c *gin.Context, userId uuid.UUID) (*database.Pagination, error) {

	var mArr []*Product
	var pagination database.Pagination

	o.db.
		Scopes(database.Paginator(c, &mArr, []string{"CreatedByUser", "UpdatedByUser", "Company", "ProductHistory"}, &pagination)).
		Where("product_created_by", userId).
		Find(&mArr)
	pagination.Data = &mArr
	return &pagination, nil
}

func (o *productOrm) GetOneProductByID(id uuid.UUID) (m Product, err error) {
	result := o.db.Model(&m).
		Preload("CreatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt"})
		}).
		Preload("UpdatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt"})
		}).
		Preload("Company", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedAt", "UpdatedAt"})
		}).
		First(&m, id)
	return m, result.Error
}

func (o *productOrm) GetOneProductByName(name string) (m Product, err error) {
	result := o.db.Model(&m).Where("lower(name) = ?", strings.ToLower(name)).First(&m)
	return m, result.Error
}

func (o *productOrm) InsertProduct(p Product) (err error) {
	result := o.db.Model(&p).Omit(clause.Associations).Create(&p)
	return result.Error
}

func (o *productOrm) UpdateProduct(id uuid.UUID, p Product) (err error) {
	result := o.db.Model(&p).Where("id = ?", id).Updates(&p)
	return result.Error
}

func (o *productOrm) DeleteProduct(id uuid.UUID, tx *gorm.DB) (err error) {
	result := tx.Model(&Product{}).Delete(&Product{}, id)
	return result.Error
}
