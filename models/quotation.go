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

type quotationOrm struct {
	db *gorm.DB
}

type Quotation struct {
	ID          uuid.UUID `gorm:"index:id,unique;type:uuid;default:gen_random_uuid();" json:"id"`
	Number      string    `json:"number,omitempty" gorm:"not null;unique"`
	FrancoArea  string    `json:"francoArea,omitempty" gorm:"not null"`
	PaymentTerm string    `json:"paymentTerm,omitempty" gorm:"not null"`
	SendAfter   string    `json:"sendAfter,omitempty" gorm:"not null"`
	Date        time.Time `json:"date,omitempty" gorm:"not null;default:CURRENT_TIMESTAMP"`
	Sequence    string    `json:"sequence,omitempty" gorm:"not null;default:0001"`

	OperatingActivityID uuid.UUID          `json:"operatingActivityId" gorm:"type:uuid;default:NULL"`
	OperatingActivity   *OperatingActivity `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:OperatingActivityID;references:ID" json:"operatingActivity"`

	QuotationCreatedBy uuid.UUID          `json:"createdBy" gorm:"type:uuid;not null;default:NULL"`
	QuotationUpdatedBy uuid.UUID          `json:"updatedBy" gorm:"type:uuid;default:NULL"`
	CreatedByUser      *masterModels.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:QuotationCreatedBy;references:ID" json:"createdByUser"`
	UpdatedByUser      *masterModels.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:QuotationUpdatedBy;references:ID" json:"updatedByUser"`

	types.DefaultModelProperty
}

type QuotationModelAction interface {
	GetAllQuotation(userId uuid.UUID) (m []Quotation, err error)
	GetAllQuotationPaginated(c *gin.Context, userId uuid.UUID) (*database.Pagination, error)
	GetOneQuotationByID(id uuid.UUID) (m Quotation, err error)
	GetOneQuotationByNumber(name string) (m Quotation, err error)

	InsertQuotation(p Quotation) (err error)
	UpdateQuotation(id uuid.UUID, p Quotation) (err error)
	DeleteQuotationByOptActID(id uuid.UUID, tx *gorm.DB) (err error)
}

func NewQuotationAction(db *gorm.DB) QuotationModelAction {
	return &quotationOrm{db}
}

func (o *quotationOrm) GetAllQuotation(userId uuid.UUID) (m []Quotation, err error) {
	result := o.db.Model(&m).
		Preload("CreatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.
				Select([]string{"ID", "Name", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt"}).
				First(&masterModels.User{}, userId)
		}).
		Preload("UpdatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedAt", "UpdatedAt"})
		}).
		Preload("OperatingActivity.Client").
		Find(&m)

	return m, result.Error
}

func (o *quotationOrm) GetAllQuotationPaginated(c *gin.Context, userId uuid.UUID) (*database.Pagination, error) {

	var mArr []*Quotation
	var pagination database.Pagination

	o.db.
		Scopes(database.Paginator(c, &mArr, []string{"CreatedByUser", "UpdatedByUser", "OperatingActivity.Client"}, &pagination)).
		Where("quotation_created_by", userId).
		Find(&mArr)
	pagination.Data = &mArr
	return &pagination, nil
}

func (o *quotationOrm) GetOneQuotationByID(id uuid.UUID) (m Quotation, err error) {
	result := o.db.Model(&m).
		Preload("CreatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt"})
		}).
		Preload("UpdatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt"})
		}).
		Preload("OperatingActivity.Client").
		First(&m, id)
	return m, result.Error
}

func (o *quotationOrm) GetOneQuotationByNumber(number string) (m Quotation, err error) {
	result := o.db.Model(&m).Where("lower(number) = ?", strings.ToLower(number)).First(&m)
	return m, result.Error
}

func (o *quotationOrm) InsertQuotation(p Quotation) (err error) {
	result := o.db.Model(&p).Omit(clause.Associations).Create(&p)
	return result.Error
}

func (o *quotationOrm) UpdateQuotation(id uuid.UUID, p Quotation) (err error) {
	result := o.db.Model(&p).Where("id = ?", id).Updates(&p)
	return result.Error
}

func (o *quotationOrm) DeleteQuotationByOptActID(id uuid.UUID, tx *gorm.DB) (err error) {
	result := tx.Model(&Quotation{}).Where("operating_activity_id = ?", id).Delete(&Quotation{})
	return result.Error
}
