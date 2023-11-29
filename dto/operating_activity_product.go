package dto

import (
	"github.com/google/uuid"
)

type InsertOperatingActivityProduct struct {
	Quantity            float32   `json:"quantity,omitempty" binding:"required"`
	VATRate             int16     `json:"vatRate,omitempty" binding:"required"`
	SubTotal            float64   `json:"subTotal,omitempty" binding:"required"`
	GrandTotal          float64   `json:"grandTotal,omitempty" binding:"required"`
	OperatingActivityID uuid.UUID `json:"operatingActivityId" binding:"required"`
	ProductID           uuid.UUID `json:"productId" binding:"required"`
	CreatedBy           uuid.UUID
}

type UpdateOperatingActivityProduct struct {
	ID                  uuid.UUID
	Quantity            float32   `json:"quantity,omitempty" binding:"required"`
	VATRate             int16     `json:"vatRate,omitempty" binding:"required"`
	SubTotal            float64   `json:"subTotal,omitempty" binding:"required"`
	GrandTotal          float64   `json:"grandTotal,omitempty" binding:"required"`
	OperatingActivityID uuid.UUID `json:"operatingActivityId" binding:"required"`
	ProductID           uuid.UUID `json:"productId" binding:"required"`
	UpdatedBy           uuid.UUID
}
