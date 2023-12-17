package dto

import (
	"github.com/google/uuid"
	"time"
)

type InsertQuotationValidator struct {
	FrancoArea          string `json:"francoArea,omitempty" binding:"required" validate:"required"`
	PaymentTerm         string `json:"paymentTerm,omitempty" binding:"required" validate:"required"`
	SendAfter           string `json:"sendAfter,omitempty" binding:"required" validate:"required"`
	Date                string `json:"date,omitempty" binding:"required" validate:"required,ISO8601date"`
	OperatingActivityID string `json:"operatingActivityId,omitempty"`
	CreatedBy           uuid.UUID
}

type UpdateQuotationValidator struct {
	ID                  uuid.UUID
	FrancoArea          string `json:"francoArea,omitempty" binding:"required" validate:"required"`
	PaymentTerm         string `json:"paymentTerm,omitempty" binding:"required" validate:"required"`
	SendAfter           string `json:"sendAfter,omitempty" binding:"required" validate:"required"`
	Date                string `json:"date,omitempty" binding:"required" validate:"required,ISO8601date"`
	OperatingActivityID string `json:"operatingActivityId,omitempty"`
	UpdatedBy           uuid.UUID
}

type InsertQuotation struct {
	FrancoArea          string
	PaymentTerm         string
	SendAfter           string
	Date                time.Time
	OperatingActivityID uuid.UUID
	CreatedBy           uuid.UUID
}

type UpdateQuotation struct {
	ID                  uuid.UUID
	FrancoArea          string
	PaymentTerm         string
	SendAfter           string
	Date                time.Time
	OperatingActivityID uuid.UUID
	UpdatedBy           uuid.UUID
}
