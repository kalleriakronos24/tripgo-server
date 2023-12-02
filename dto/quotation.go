package dto

import (
	"github.com/google/uuid"
	"time"
)

type InsertQuotationValidator struct {
	Number              string `json:"number,omitempty" binding:"required" validate:"required"`
	FrancoArea          string `json:"francoArea,omitempty" binding:"required" validate:"required"`
	PaymentTerm         string `json:"paymentTerm,omitempty" binding:"required" validate:"required"`
	SendAfter           string `json:"sendAfter,omitempty" binding:"required" validate:"required"`
	Date                string `json:"date,omitempty" binding:"required" validate:"required,ISO8601date"`
	OperatingActivityID string `json:"operatingActivityId" binding:"required" validate:"required,uuid4"`
	CreatedBy           uuid.UUID
}

type UpdateQuotationValidator struct {
	ID                  uuid.UUID
	Number              string `json:"number,omitempty" binding:"required" validate:"required"`
	FrancoArea          string `json:"francoArea,omitempty" binding:"required" validate:"required"`
	PaymentTerm         string `json:"paymentTerm,omitempty" binding:"required" validate:"required"`
	SendAfter           string `json:"sendAfter,omitempty" binding:"required" validate:"required"`
	Date                string `json:"date,omitempty" binding:"required" validate:"required,ISO8601date"`
	OperatingActivityID string `json:"operatingActivityId" binding:"required" validate:"required,uuid4"`
	UpdatedBy           uuid.UUID
}

type InsertQuotation struct {
	Number              string
	FrancoArea          string
	PaymentTerm         string
	SendAfter           string
	Date                time.Time
	OperatingActivityID uuid.UUID
	CreatedBy           uuid.UUID
}

type UpdateQuotation struct {
	ID                  uuid.UUID
	Number              string
	FrancoArea          string
	PaymentTerm         string
	SendAfter           string
	Date                time.Time
	OperatingActivityID uuid.UUID
	UpdatedBy           uuid.UUID
}
