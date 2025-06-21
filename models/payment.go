package models

import (
	"github.com/google/uuid"
	"github.com/kalleriakronos24/khaimal-group/models/master"
	"github.com/kalleriakronos24/khaimal-group/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PaymentOrm struct {
	db *gorm.DB
}

type Payment struct {
	ID             uuid.UUID `gorm:"index:id,unique;type:uuid;default:gen_random_uuid();" json:"id"`
	Amount         float64   `gorm:"not null;default:0" json:"amount,omitempty" binding:"required"`
	PI             string    `gorm:"not null;default:NULL" json:"pi,omitempty" binding:"required"`
	PaymentMethod  string    `gorm:"not null;default:NULL" json:"paymentMethod,omitempty" binding:"required"`
	CardLastNumber string    `gorm:"not null;default:NULL" json:"cardLastNumber,omitempty" binding:"required"`
	Currency       string    `gorm:"not null;default:NULL" json:"currency,omitempty" binding:"required"`
	ReceiptURL     string    `gorm:"not null;default:NULL" json:"receiptUrl,omitempty" binding:"required"`
	Status         string    `gorm:"not null;default:NULL" json:"status,omitempty" binding:"required"`

	CustomerID        uuid.UUID `json:"customerId,omitempty" gorm:"type:uuid;default:NULL"`
	BookingTransferID uuid.UUID `json:"bookingTransferId,omitempty" gorm:"type:uuid;default:NULL"`
	DriverID          uuid.UUID `json:"driverId,omitempty" gorm:"type:uuid;default:NULL"`

	Driver          *master.Driver   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:DriverID;references:ID" json:"driver,omitempty"`
	Customer        *master.Customer `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:CustomerID;references:ID" json:"customer,omitempty"`
	BookingTransfer *BookingTransfer `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:BookingTransferID;references:ID" json:"bookingTransfer,omitempty"`
	types.DefaultModelProperty
}

type PaymentModelAction interface {
	GetOneByID(id uuid.UUID) (m Payment, err error)
	GetOneByBookigTransferID(id uuid.UUID) (Payment Payment, err error)
	GetOneLastCreatedByCustomerID(id uuid.UUID) (Payment Payment, err error)
	GetOneByPaymentIntentID(id uuid.UUID) (Payment Payment, err error)
	GetOneDetailByID(id uuid.UUID) (Payment *Payment, err error)
	GetOneByEmail(email string) (m Payment, err error)
	InsertPayment(p Payment, tx *gorm.DB) (err error)
	DeletePayment(id uuid.UUID, tx *gorm.DB) (err error)
}

func NewPaymentAction(db *gorm.DB) PaymentModelAction {
	return &PaymentOrm{db}
}

func (o *PaymentOrm) GetOneByBookigTransferID(id uuid.UUID) (Payment Payment, err error) {
	result := o.db.Model(&Payment).
		Where("booking_transfer_id = ?", id).
		Preload(clause.Associations).
		First(&Payment)
	return Payment, result.Error
}

func (o *PaymentOrm) GetOneLastCreatedByCustomerID(id uuid.UUID) (Payment Payment, err error) {
	result := o.db.Model(&Payment).
		Where("customer_id = ?", id).
		Preload(clause.Associations).
		Order("created_at DESC").
		Limit(1).
		First(&Payment)
	return Payment, result.Error
}

func (o *PaymentOrm) GetOneByPaymentIntentID(id uuid.UUID) (Payment Payment, err error) {
	result := o.db.Model(&Payment).
		Where("pi = ?", id).
		Preload(clause.Associations).
		First(&Payment)
	return Payment, result.Error
}

func (o *PaymentOrm) GetOneByID(id uuid.UUID) (Payment Payment, err error) {
	result := o.db.Model(&Payment).
		Where("driver_id = ?", id).
		Preload(clause.Associations).
		First(&Payment)
	return Payment, result.Error
}

func (o *PaymentOrm) GetOneDetailByID(id uuid.UUID) (Payment *Payment, err error) {
	result := o.db.Model(&Payment).
		Where("driver_id = ?", id).
		Preload("Driver", func(db *gorm.DB) *gorm.DB {
			return db.Preload(clause.Associations).First(&master.Driver{})
		}).
		First(&Payment)
	return Payment, result.Error
}

func (o *PaymentOrm) GetOneByEmail(email string) (m Payment, err error) {
	result := o.db.Model(&m).Where("email = ?", email).First(&m)
	return m, result.Error
}

func (o *PaymentOrm) InsertPayment(p Payment, tx *gorm.DB) (err error) {
	result := tx.Model(&p).Create(&p)
	return result.Error
}

func (o *PaymentOrm) DeletePayment(id uuid.UUID, tx *gorm.DB) (err error) {
	result := tx.Model(&Payment{}).Delete(&Payment{}, id)
	return result.Error
}
