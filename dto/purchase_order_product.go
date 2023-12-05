package dto

import (
	"github.com/google/uuid"
)

type InsertPurchaseOrderProductValidator struct {
	Quantity        float64 `json:"quantity,omitempty" binding:"required" validate:"required"`
	VATRate         int16   `json:"vatRate,omitempty" binding:"required" validate:"required"`
	SubTotal        float64 `json:"subTotal,omitempty" binding:"required" validate:"required"`
	GrandTotal      float64 `json:"grandTotal,omitempty" binding:"required" validate:"required"`
	PurchaseOrderID string  `json:"purchaseOrderId" binding:"required" validate:"required,uuid"`
	ClientID        string  `json:"clientId" binding:"required" validate:"required,uuid"`
	ProductID       string  `json:"productId" binding:"required" validate:"required,uuid"`
	CreatedBy       uuid.UUID
}

type UpdatePurchaseOrderProductValidator struct {
	ID              uuid.UUID
	Quantity        float64 `json:"quantity,omitempty" binding:"required" validate:"required"`
	VATRate         int16   `json:"vatRate,omitempty" binding:"required" validate:"required"`
	SubTotal        float64 `json:"subTotal,omitempty" binding:"required" validate:"required"`
	GrandTotal      float64 `json:"grandTotal,omitempty" binding:"required" validate:"required"`
	PurchaseOrderID string  `json:"purchaseOrderId" binding:"required" validate:"required,uuid"`
	ClientID        string  `json:"clientId" binding:"required" validate:"required,uuid4"`
	ProductID       string  `json:"productId" binding:"required" validate:"required,uuid4"`
	UpdatedBy       uuid.UUID
}

type InsertPurchaseOrderProduct struct {
	Quantity        float64
	VATRate         int16
	SubTotal        float64
	GrandTotal      float64
	PurchaseOrderID uuid.UUID
	ClientID        uuid.UUID
	ProductID       uuid.UUID
	CreatedBy       uuid.UUID
}

type UpdatePurchaseOrderProduct struct {
	ID              uuid.UUID
	Quantity        float64
	VATRate         int16
	SubTotal        float64
	GrandTotal      float64
	PurchaseOrderID uuid.UUID
	ClientID        uuid.UUID
	ProductID       uuid.UUID
	UpdatedBy       uuid.UUID
}
