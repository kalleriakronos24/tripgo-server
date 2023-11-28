package models

import (
	"github.com/google/uuid"
	"gitlab.com/odma1/odma-be/models/master"
	"gitlab.com/odma1/odma-be/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type PaymentOrm struct {
	db *gorm.DB
}

type Payment struct {
	ID             uuid.UUID `gorm:"index:id,unique;type:uuid;default:gen_random_uuid();" json:"id"`
	Currency       string    `json:"currency,omitempty" gorm:"not null"`
	Amount         float64   `json:"amount,omitempty" gorm:"not null"`
	Term           string    `json:"term,omitempty" gorm:"not null"`
	ARAging        int16     `json:"araging,omitempty" gorm:"not null;default:0"`
	Remarks        string    `json:"remarks,omitempty" gorm:"not null"`
	DueDate        time.Time `json:"dueDate,omitempty" gorm:"not null"`
	SettlementDate time.Time `json:"settlementDate,omitempty" gorm:"not null"`

	OperatingActivityID uuid.UUID          `json:"operatingActivityId" gorm:"type:uuid;not null;default:NULL;"`
	OperatingActivity   *OperatingActivity `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:OperatingActivityID;references:ID" json:"operatingActivity"`

	PaymentCreatedBy uuid.UUID    `json:"createdBy" gorm:"type:uuid;not null;default:NULL;"`
	PaymentUpdatedBy uuid.UUID    `json:"updatedBy" gorm:"type:uuid;default:NULL;"`
	CreatedByUser    *master.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:PaymentCreatedBy;references:ID" json:"createdByUser"`
	UpdatedByUser    *master.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:PaymentUpdatedBy;references:ID" json:"updatedByUser"`

	types.DefaultModelProperty
}

type PaymentModelAction interface {
	GetAllPayment(userId uuid.UUID) (m []Payment, err error)
	GetOnePaymentByID(id uuid.UUID) (m Payment, err error)
	GetOnePaymentByOperatingActivityID(productId uuid.UUID) (m Payment, err error)

	InsertPayment(p Payment) (err error)
	UpdatePayment(id uuid.UUID, p Payment) (err error)
}

func NewPaymentAction(db *gorm.DB) PaymentModelAction {
	return &PaymentOrm{db}
}

func (o *PaymentOrm) GetAllPayment(userId uuid.UUID) (m []Payment, err error) {
	result := o.db.Model(&m).
		Preload("CreatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.
				Select([]string{"ID", "Name", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt"}).
				First(&master.User{}, userId)
		}).
		Preload("UpdatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedAt", "UpdatedAt"})
		}).
		Preload("OperatingActivity").
		Find(&m)
	return m, result.Error
}

func (o *PaymentOrm) GetOnePaymentByID(id uuid.UUID) (m Payment, err error) {
	result := o.db.Model(&m).
		Preload("CreatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt"})
		}).
		Preload("UpdatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt"})
		}).
		Preload("OperatingActivity").
		First(&m, id)
	return m, result.Error
}

func (o *PaymentOrm) GetOnePaymentByOperatingActivityID(operatingActivityId uuid.UUID) (m Payment, err error) {
	result := o.db.Model(&m).Where("operating_activity_id = ?", operatingActivityId).First(&m)
	return m, result.Error
}

func (o *PaymentOrm) InsertPayment(p Payment) (err error) {
	result := o.db.Model(&p).Omit(clause.Associations).Create(&p)
	return result.Error
}

func (o *PaymentOrm) UpdatePayment(id uuid.UUID, p Payment) (err error) {
	result := o.db.Model(&p).Where("id = ?", id).Updates(&p)
	return result.Error
}
