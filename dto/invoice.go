package dto

import (
	"github.com/google/uuid"
	"time"
)

type InsertInvoice struct {
	Number              string    `json:"number,omitempty" binding:"required"`
	Type                string    `json:"type,omitempty" binding:"required"`
	Date                time.Time `json:"date,omitempty" binding:"required"`
	OperatingActivityID uuid.UUID `json:"operatingActivityId,omitempty" binding:"required"`
	CreatedBy           uuid.UUID
}

type UpdateInvoice struct {
	ID                  uuid.UUID
	Number              string    `json:"number,omitempty" binding:"required"`
	Type                string    `json:"type,omitempty" binding:"required"`
	Date                time.Time `json:"date,omitempty" binding:"required"`
	OperatingActivityID uuid.UUID `json:"operatingActivityId" binding:"required"`
	UpdatedBy           uuid.UUID
}
