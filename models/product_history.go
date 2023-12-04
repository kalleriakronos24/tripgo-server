package models

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	database "gitlab.com/odma1/odma-be/db"
	"gitlab.com/odma1/odma-be/models/master"
	"gitlab.com/odma1/odma-be/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductHistoryOrm struct {
	db *gorm.DB
}

type ProductHistory struct {
	ID       uuid.UUID `gorm:"index:id,unique;type:uuid;default:gen_random_uuid();" json:"id"`
	Status   string    `json:"status,omitempty" gorm:"not null"`
	Quantity int8      `json:"quantity,omitempty" gorm:"not null;default:0"`

	ProductID           uuid.UUID          `json:"productId" gorm:"type:uuid;not null;default:NULL;"`
	Product             *Product           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ProductID;references:ID" json:"product"`
	OperatingActivityID uuid.UUID          `json:"operatingActivityId" gorm:"type:uuid;not null;default:NULL;"`
	OperatingActivity   *OperatingActivity `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:OperatingActivityID;references:ID" json:"operatingActivity"`

	ProductHistoryCreatedBy uuid.UUID    `json:"createdBy" gorm:"type:uuid;not null;default:NULL;"`
	ProductHistoryUpdatedBy uuid.UUID    `json:"updatedBy" gorm:"type:uuid;default:NULL;"`
	CreatedByUser           *master.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ProductHistoryCreatedBy;references:ID" json:"createdByUser"`
	UpdatedByUser           *master.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ProductHistoryUpdatedBy;references:ID" json:"updatedByUser"`

	types.DefaultModelProperty
}

type ProductHistoryModelAction interface {
	GetAllProductHistory(userId uuid.UUID) (m []ProductHistory, err error)
	GetAllProductHistoryPaginated(c *gin.Context, userId uuid.UUID) (*database.Pagination, error)
	GetOneProductHistoryByID(id uuid.UUID) (m ProductHistory, err error)
	GetOneProductHistoryByProductID(productId uuid.UUID) (m ProductHistory, err error)

	InsertProductHistory(p ProductHistory) (err error)
	UpdateProductHistory(id uuid.UUID, p ProductHistory) (err error)
}

func NewProductHistoryAction(db *gorm.DB) ProductHistoryModelAction {
	return &ProductHistoryOrm{db}
}

func (o *ProductHistoryOrm) GetAllProductHistory(userId uuid.UUID) (m []ProductHistory, err error) {
	result := o.db.Model(&m).
		Preload("CreatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.
				Select([]string{"ID", "Name", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt"}).
				First(&master.User{}, userId)
		}).
		Preload("UpdatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedAt", "UpdatedAt"})
		}).
		Preload("Product").
		Preload("OperatingActivity").
		Find(&m)
	return m, result.Error
}

func (o *ProductHistoryOrm) GetAllProductHistoryPaginated(c *gin.Context, userId uuid.UUID) (*database.Pagination, error) {

	var mArr []*ProductHistory
	var pagination database.Pagination

	o.db.
		Scopes(database.Paginator(c, &mArr, []string{"CreatedByUser", "UpdatedByUser", "Product", "OperatingActivity"}, &pagination)).
		Where("product_history_created_by", userId).
		Find(&mArr)
	pagination.Data = &mArr
	return &pagination, nil
}

func (o *ProductHistoryOrm) GetOneProductHistoryByID(id uuid.UUID) (m ProductHistory, err error) {
	result := o.db.Model(&m).
		Preload("CreatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt"})
		}).
		Preload("UpdatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt"})
		}).
		Preload("Product").
		Preload("OperatingActivity").
		First(&m, id)
	return m, result.Error
}

func (o *ProductHistoryOrm) GetOneProductHistoryByProductID(productId uuid.UUID) (m ProductHistory, err error) {
	result := o.db.Model(&m).Where("product_id = ?", productId).First(&m)
	return m, result.Error
}

func (o *ProductHistoryOrm) InsertProductHistory(p ProductHistory) (err error) {
	result := o.db.Model(&p).Omit(clause.Associations).Create(&p)
	return result.Error
}

func (o *ProductHistoryOrm) UpdateProductHistory(id uuid.UUID, p ProductHistory) (err error) {
	result := o.db.Model(&p).Where("id = ?", id).Updates(&p)
	return result.Error
}
