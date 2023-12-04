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
	"time"
)

type DeliveryOrderOrm struct {
	db *gorm.DB
}

/**
Delivery Order statuses
- Shipped
- Delivered
- Canceled
- Out for Delivery
- Delivery expected (end of updates)
- Failed delivery attempts
*/

type DeliveryOrder struct {
	ID            uuid.UUID `gorm:"index:id,unique;type:uuid;default:gen_random_uuid();" json:"id"`
	Number        string    `json:"number,omitempty" gorm:"not null;unique"`
	ContactPerson string    `json:"contactPerson,omitempty" gorm:"not null"`
	PhoneNumber   string    `json:"phoneNumber,omitempty" gorm:"not null"`
	Address       string    `json:"address,omitempty" gorm:"not null"`
	Note          string    `json:"note,omitempty"`
	Date          time.Time `json:"date,omitempty" gorm:"not null"`
	Status        string    `json:"status,omitempty" gorm:"not null;default:shipped"`

	DocumentID          uuid.UUID          `json:"documentId,omitempty" gorm:"not null;"`
	Document            *Document          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:DocumentID;references:ID" json:"document"`
	OperatingActivityID uuid.UUID          `json:"operatingActivityId" gorm:"type:uuid;not null;default:NULL;"`
	OperatingActivity   *OperatingActivity `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:OperatingActivityID;references:ID" json:"operatingActivity"`

	DeliveryOrderCreatedBy uuid.UUID          `json:"createdBy" gorm:"type:uuid;not null;default:NULL;"`
	DeliveryOrderUpdatedBy uuid.UUID          `json:"updatedBy" gorm:"type:uuid;default:NULL;"`
	CreatedByUser          *masterModels.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:DeliveryOrderCreatedBy;references:ID" json:"createdByUser"`
	UpdatedByUser          *masterModels.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:DeliveryOrderUpdatedBy;references:ID" json:"updatedByUser"`

	types.DefaultModelProperty
}

type DeliveryOrderModelAction interface {
	GetAllDeliveryOrder(userId uuid.UUID) (m []DeliveryOrder, err error)
	GetAllDeliveryOrderPaginated(c *gin.Context, userId uuid.UUID) (*database.Pagination, error)
	GetOneDeliveryOrderByID(id uuid.UUID) (m DeliveryOrder, err error)
	GetOneDeliveryOrderByNumber(number string) (m DeliveryOrder, err error)
	GetOneDeliveryOrderByOperatingActivityID(id uuid.UUID) (m DeliveryOrder, err error)

	InsertDeliveryOrder(p DeliveryOrder) (err error)
	UpdateDeliveryOrder(id uuid.UUID, p DeliveryOrder) (err error)
}

func NewDeliveryOrderAction(db *gorm.DB) DeliveryOrderModelAction {
	return &DeliveryOrderOrm{db}
}

func (o *DeliveryOrderOrm) GetAllDeliveryOrder(userId uuid.UUID) (m []DeliveryOrder, err error) {
	result := o.db.Model(&m).
		Preload("CreatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.
				Select([]string{"ID", "Name", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt"}).
				First(&masterModels.User{}, userId)
		}).
		Preload("UpdatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedAt", "UpdatedAt"})
		}).
		Preload("Document").
		Preload("OperatingActivity.OperatingActivityProduct").
		Find(&m)
	return m, result.Error
}
func (o *DeliveryOrderOrm) GetAllDeliveryOrderPaginated(c *gin.Context, userId uuid.UUID) (*database.Pagination, error) {

	var mArr []*DeliveryOrder
	var pagination database.Pagination

	o.db.
		Scopes(database.Paginator(c, &mArr, []string{"CreatedByUser", "UpdatedByUser", "Document", "OperatingActivity.OperatingActivityProduct"}, &pagination)).
		Where("delivery_order_created_by", userId).
		Find(&mArr)
	pagination.Data = &mArr
	return &pagination, nil
}

func (o *DeliveryOrderOrm) GetOneDeliveryOrderByID(id uuid.UUID) (m DeliveryOrder, err error) {
	result := o.db.Model(&m).
		Preload("CreatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt"})
		}).
		Preload("UpdatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt"})
		}).
		Preload("Document").
		Preload("OperatingActivity.OperatingActivityProduct").
		First(&m, id)
	return m, result.Error
}

func (o *DeliveryOrderOrm) GetOneDeliveryOrderByNumber(number string) (m DeliveryOrder, err error) {
	result := o.db.Model(&m).Where("lower(number) = ?", strings.ToLower(number)).First(&m)
	return m, result.Error
}

func (o *DeliveryOrderOrm) GetOneDeliveryOrderByOperatingActivityID(id uuid.UUID) (m DeliveryOrder, err error) {
	result := o.db.Model(&m).Where("operating_activity_id = ?", id).First(&m)
	return m, result.Error
}

func (o *DeliveryOrderOrm) InsertDeliveryOrder(p DeliveryOrder) (err error) {
	result := o.db.Model(&p).Omit(clause.Associations).Create(&p)
	return result.Error
}

func (o *DeliveryOrderOrm) UpdateDeliveryOrder(id uuid.UUID, p DeliveryOrder) (err error) {
	result := o.db.Model(&p).Where("id = ?", id).Updates(&p)
	return result.Error
}
