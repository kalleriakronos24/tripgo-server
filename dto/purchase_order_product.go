package dto

import (
	"github.com/google/uuid"
)

type InsertPurchaseOrderProduct struct {
	Quantity        float32   `json:"quantity,omitempty" binding:"required"`
	VATRate         int16     `json:"vatRate,omitempty" binding:"required"`
	SubTotal        float64   `json:"subTotal,omitempty" binding:"required"`
	GrandTotal      float64   `json:"grandTotal,omitempty" binding:"required"`
	PurchaseOrderID uuid.UUID `json:"purchaseOrderId" binding:"required"`
	ClientID        uuid.UUID `json:"clientId" binding:"required"`
	ProductID       uuid.UUID `json:"productId" binding:"required"`
	CreatedBy       uuid.UUID
}

type UpdatePurchaseOrderProduct struct {
	ID              uuid.UUID
	Quantity        float32   `json:"quantity,omitempty" binding:"required"`
	VATRate         int16     `json:"vatRate,omitempty" binding:"required"`
	SubTotal        float64   `json:"subTotal,omitempty" binding:"required"`
	GrandTotal      float64   `json:"grandTotal,omitempty" binding:"required"`
	PurchaseOrderID uuid.UUID `json:"purchaseOrderId" binding:"required"`
	ClientID        uuid.UUID `json:"clientId" binding:"required"`
	ProductID       uuid.UUID `json:"productId" binding:"required"`
	UpdatedBy       uuid.UUID
}
