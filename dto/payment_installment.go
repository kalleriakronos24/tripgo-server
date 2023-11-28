package dto

import (
	"github.com/google/uuid"
	"time"
)

type InsertPaymentInstallment struct {
	Currency  string    `json:"currency,omitempty" binding:"required"`
	Amount    float64   `json:"amount,omitempty" binding:"required"`
	Note      string    `json:"note,omitempty" binding:"required"`
	Date      time.Time `json:"date,omitempty" binding:"required"`
	PaymentID uuid.UUID `json:"paymentId" binding:"required"`
	CreatedBy uuid.UUID
}

type UpdatePaymentInstallment struct {
	ID        uuid.UUID
	Currency  string    `json:"currency,omitempty" binding:"required"`
	Amount    float64   `json:"amount,omitempty" binding:"required"`
	Note      string    `json:"note,omitempty" binding:"required"`
	Date      time.Time `json:"date,omitempty" binding:"required"`
	PaymentID uuid.UUID `json:"paymentId" binding:"required"`
	UpdatedBy uuid.UUID
}
