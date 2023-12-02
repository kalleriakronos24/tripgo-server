package dto

import (
	"github.com/google/uuid"
	"time"
)

type InsertInvoiceValidator struct {
	Number              string `json:"number,omitempty" binding:"required" validate:"required"`
	Type                string `json:"type,omitempty" binding:"required" validate:"required,oneof=proforma invoice"`
	Date                string `json:"date,omitempty" binding:"required" validate:"required,ISO8601date"`
	OperatingActivityID string `json:"operatingActivityId,omitempty" binding:"required" validate:"uuid4"`
	CreatedBy           uuid.UUID
}

type UpdateInvoiceValidator struct {
	ID                  uuid.UUID
	Number              string `json:"number,omitempty" binding:"required" validate:"required"`
	Type                string `json:"type,omitempty" binding:"required" validate:"required,oneof=proforma invoice"`
	Date                string `json:"date,omitempty" binding:"required" validate:"required,ISO8601date"`
	OperatingActivityID string `json:"operatingActivityId,omitempty" binding:"required" validate:"uuid4"`
	UpdatedBy           uuid.UUID
}

type InsertInvoice struct {
	Number              string
	Type                string
	Date                time.Time
	OperatingActivityID uuid.UUID
	CreatedBy           uuid.UUID
}

type UpdateInvoice struct {
	ID                  uuid.UUID
	Number              string
	Type                string
	Date                time.Time
	OperatingActivityID uuid.UUID
	UpdatedBy           uuid.UUID
}
