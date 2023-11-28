package models

import (
	"github.com/google/uuid"
	"gitlab.com/odma1/odma-be/models/master"
	"gitlab.com/odma1/odma-be/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type PaymentInstallmentOrm struct {
	db *gorm.DB
}

type PaymentInstallment struct {
	ID       uuid.UUID `gorm:"index:id,unique;type:uuid;default:gen_random_uuid();" json:"id"`
	Currency string    `json:"currency,omitempty" gorm:"not null"`
	Amount   float64   `json:"amount,omitempty" gorm:"not null"`
	Note     string    `json:"note,omitempty" gorm:"not null"`
	Date     time.Time `json:"date,omitempty" gorm:"not null"`

	PaymentID uuid.UUID `json:"paymentId" gorm:"type:uuid;not null;default:NULL;"`
	Payment   *Payment  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:PaymentID;references:ID" json:"payment"`

	PaymentInstallmentCreatedBy uuid.UUID    `json:"createdBy" gorm:"type:uuid;not null;default:NULL;"`
	PaymentInstallmentUpdatedBy uuid.UUID    `json:"updatedBy" gorm:"type:uuid;default:NULL;"`
	CreatedByUser               *master.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:PaymentInstallmentCreatedBy;references:ID" json:"createdByUser"`
	UpdatedByUser               *master.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:PaymentInstallmentUpdatedBy;references:ID" json:"updatedByUser"`

	types.DefaultModelProperty
}

type PaymentInstallmentModelAction interface {
	GetAllPaymentInstallment(userId uuid.UUID) (m []PaymentInstallment, err error)
	GetOnePaymentInstallmentByID(id uuid.UUID) (m PaymentInstallment, err error)
	GetOnePaymentInstallmentByPaymentID(productId uuid.UUID) (m PaymentInstallment, err error)

	InsertPaymentInstallment(p PaymentInstallment) (err error)
	UpdatePaymentInstallment(id uuid.UUID, p PaymentInstallment) (err error)
}

func NewPaymentInstallmentAction(db *gorm.DB) PaymentInstallmentModelAction {
	return &PaymentInstallmentOrm{db}
}

func (o *PaymentInstallmentOrm) GetAllPaymentInstallment(userId uuid.UUID) (m []PaymentInstallment, err error) {
	result := o.db.Model(&m).
		Preload("CreatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.
				Select([]string{"ID", "Name", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt"}).
				First(&master.User{}, userId)
		}).
		Preload("UpdatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedAt", "UpdatedAt"})
		}).
		Preload("Payment").
		Find(&m)
	return m, result.Error
}

func (o *PaymentInstallmentOrm) GetOnePaymentInstallmentByID(id uuid.UUID) (m PaymentInstallment, err error) {
	result := o.db.Model(&m).
		Preload("CreatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt"})
		}).
		Preload("UpdatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt"})
		}).
		Preload("Payment").
		First(&m, id)
	return m, result.Error
}

func (o *PaymentInstallmentOrm) GetOnePaymentInstallmentByPaymentID(paymentId uuid.UUID) (m PaymentInstallment, err error) {
	result := o.db.Model(&m).Where("payment_id = ?", paymentId).First(&m)
	return m, result.Error
}

func (o *PaymentInstallmentOrm) InsertPaymentInstallment(p PaymentInstallment) (err error) {
	result := o.db.Model(&p).Omit(clause.Associations).Create(&p)
	return result.Error
}

func (o *PaymentInstallmentOrm) UpdatePaymentInstallment(id uuid.UUID, p PaymentInstallment) (err error) {
	result := o.db.Model(&p).Where("id = ?", id).Updates(&p)
	return result.Error
}
