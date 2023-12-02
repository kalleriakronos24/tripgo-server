package dto

import (
	"github.com/google/uuid"
	"time"
)

type InsertPaymentInstallmentValidator struct {
	Currency  string  `json:"currency,omitempty" binding:"required" validate:"required"`
	Amount    float64 `json:"amount,omitempty" binding:"required" validate:"required"`
	Note      string  `json:"note,omitempty" binding:"required" validate:"required"`
	Date      string  `json:"date,omitempty" binding:"required" validate:"required,ISO8601date"`
	PaymentID string  `json:"paymentId" binding:"required" validate:"required,uuid4"`
	CreatedBy uuid.UUID
}

type UpdatePaymentInstallmentValidator struct {
	ID        uuid.UUID
	Currency  string  `json:"currency,omitempty" binding:"required" validate:"required"`
	Amount    float64 `json:"amount,omitempty" binding:"required" validate:"required"`
	Note      string  `json:"note,omitempty" binding:"required" validate:"required"`
	Date      string  `json:"date,omitempty" binding:"required" validate:"required,ISO8601date"`
	PaymentID string  `json:"paymentId" binding:"required" validate:"required,uuid4"`
	UpdatedBy uuid.UUID
}

type InsertPaymentInstallment struct {
	Currency  string
	Amount    float64
	Note      string
	Date      time.Time
	PaymentID uuid.UUID
	CreatedBy uuid.UUID
}

type UpdatePaymentInstallment struct {
	ID        uuid.UUID
	Currency  string
	Amount    float64
	Note      string
	Date      time.Time
	PaymentID uuid.UUID
	UpdatedBy uuid.UUID
}
