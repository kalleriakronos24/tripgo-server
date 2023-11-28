package dto

import (
	"github.com/google/uuid"
	"time"
)

type InsertQuotation struct {
	Number              string    `json:"number,omitempty" binding:"required"`
	FrancoArea          string    `json:"francoArea,omitempty" binding:"required"`
	PaymentTerm         string    `json:"paymentTerm,omitempty" binding:"required"`
	SendAfter           string    `json:"sendAfter,omitempty" binding:"required"`
	Date                time.Time `json:"date,omitempty" binding:"required"`
	OperatingActivityID uuid.UUID `json:"operatingActivityId" binding:"required"`
	CreatedBy           uuid.UUID
}

type UpdateQuotation struct {
	ID                  uuid.UUID
	Number              string    `json:"number,omitempty" binding:"required"`
	FrancoArea          string    `json:"francoArea,omitempty" binding:"required"`
	PaymentTerm         string    `json:"paymentTerm,omitempty" binding:"required"`
	SendAfter           string    `json:"sendAfter,omitempty" binding:"required"`
	Date                time.Time `json:"date,omitempty" binding:"required"`
	OperatingActivityID uuid.UUID `json:"operatingActivityId" binding:"required"`
	UpdatedBy           uuid.UUID
}
