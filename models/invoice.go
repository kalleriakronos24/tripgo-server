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

type InvoiceOrm struct {
	db *gorm.DB
}
type Invoice struct {
	ID     uuid.UUID `gorm:"index:id,unique;type:uuid;default:gen_random_uuid();" json:"id"`
	Number string    `json:"number,omitempty" gorm:"not null;unique"`
	Type   string    `json:"type,omitempty" gorm:"not null;default:proforma"`
	Date   time.Time `json:"date,omitempty" gorm:"not null"`

	OperatingActivityID uuid.UUID          `json:"operatingActivityId" gorm:"type:uuid;not null;default:NULL;"`
	OperatingActivity   *OperatingActivity `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:OperatingActivityID;references:ID" json:"operatingActivity"`

	InvoiceCreatedBy uuid.UUID          `json:"createdBy" gorm:"type:uuid;not null;default:NULL;"`
	InvoiceUpdatedBy uuid.UUID          `json:"updatedBy" gorm:"type:uuid;default:NULL;"`
	CreatedByUser    *masterModels.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:InvoiceCreatedBy;references:ID" json:"createdByUser"`
	UpdatedByUser    *masterModels.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:InvoiceUpdatedBy;references:ID" json:"updatedByUser"`

	types.DefaultModelProperty
}

type InvoiceModelAction interface {
	GetAllInvoice(userId uuid.UUID) (m []Invoice, err error)
	GetAllInvoicePaginated(c *gin.Context, userId uuid.UUID) (*database.Pagination, error)
	GetOneInvoiceByID(id uuid.UUID) (m Invoice, err error)
	GetOneInvoiceByNumber(number string) (m Invoice, err error)
	GetOneInvoiceByOperatingActivityID(id uuid.UUID) (m Invoice, err error)

	InsertInvoice(p Invoice) (err error)
	UpdateInvoice(id uuid.UUID, p Invoice) (err error)
}

func NewInvoiceAction(db *gorm.DB) InvoiceModelAction {
	return &InvoiceOrm{db}
}

func (o *InvoiceOrm) GetAllInvoice(userId uuid.UUID) (m []Invoice, err error) {
	result := o.db.Model(&m).
		Preload("CreatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.
				Select([]string{"ID", "Name", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt"}).
				First(&masterModels.User{}, userId)
		}).
		Preload("UpdatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedAt", "UpdatedAt"})
		}).
		Preload("OperatingActivity.OperatingActivityProduct").
		Find(&m)
	return m, result.Error
}

func (o *InvoiceOrm) GetAllInvoicePaginated(c *gin.Context, userId uuid.UUID) (*database.Pagination, error) {

	var mArr []*Invoice
	var pagination database.Pagination

	o.db.
		Scopes(database.Paginator(c, &mArr, []string{"CreatedByUser", "UpdatedByUser", "OperatingActivity.OperatingActivityProduct"}, &pagination)).
		Where("invoice_created_by", userId).
		Find(&mArr)
	pagination.Data = &mArr
	return &pagination, nil
}

func (o *InvoiceOrm) GetOneInvoiceByID(id uuid.UUID) (m Invoice, err error) {
	result := o.db.Model(&m).
		Preload("CreatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt"})
		}).
		Preload("UpdatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt"})
		}).
		Preload("OperatingActivity.OperatingActivityProduct").
		First(&m, id)
	return m, result.Error
}

func (o *InvoiceOrm) GetOneInvoiceByNumber(number string) (m Invoice, err error) {
	result := o.db.Model(&m).Where("lower(number) = ?", strings.ToLower(number)).First(&m)
	return m, result.Error
}

func (o *InvoiceOrm) GetOneInvoiceByOperatingActivityID(id uuid.UUID) (m Invoice, err error) {
	result := o.db.Model(&m).Where("operating_activity_id = ?", id).First(&m)
	return m, result.Error
}

func (o *InvoiceOrm) InsertInvoice(p Invoice) (err error) {
	result := o.db.Model(&p).Omit(clause.Associations).Create(&p)
	return result.Error
}

func (o *InvoiceOrm) UpdateInvoice(id uuid.UUID, p Invoice) (err error) {
	result := o.db.Model(&p).Where("id = ?", id).Updates(&p)
	return result.Error
}
