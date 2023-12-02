package dto

import (
	"github.com/google/uuid"
)

type InsertOperatingActivityProductValidator struct {
	Quantity            float32 `json:"quantity,omitempty" validate:"required"`
	VATRate             int16   `json:"vatRate,omitempty" validate:"required"`
	SubTotal            float64 `json:"subTotal,omitempty" validate:"required"`
	GrandTotal          float64 `json:"grandTotal,omitempty" validate:"required"`
	OperatingActivityID string  `json:"operatingActivityId" validate:"required,uuid4"`
	ProductID           string  `json:"productId" validate:"required,uuid4"`
	CreatedBy           uuid.UUID
}

type UpdateOperatingActivityProductValidator struct {
	ID                  uuid.UUID
	Quantity            float32 `json:"quantity,omitempty" validate:"required"`
	VATRate             int16   `json:"vatRate,omitempty" validate:"required"`
	SubTotal            float64 `json:"subTotal,omitempty" validate:"required"`
	GrandTotal          float64 `json:"grandTotal,omitempty" validate:"required"`
	OperatingActivityID string  `json:"operatingActivityId" validate:"required,uuid44"`
	ProductID           string  `json:"productId" validate:"required,uuid4"`
	UpdatedBy           uuid.UUID
}

type InsertOperatingActivityProduct struct {
	Quantity            float32
	VATRate             int16
	SubTotal            float64
	GrandTotal          float64
	OperatingActivityID uuid.UUID
	ProductID           uuid.UUID
	CreatedBy           uuid.UUID
}

type UpdateOperatingActivityProduct struct {
	ID                  uuid.UUID
	Quantity            float32
	VATRate             int16
	SubTotal            float64
	GrandTotal          float64
	OperatingActivityID uuid.UUID
	ProductID           uuid.UUID
	UpdatedBy           uuid.UUID
}
