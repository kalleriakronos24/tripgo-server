package dto

import (
	"github.com/google/uuid"
)

type InsertOperatingActivity struct {
	TaxInvoiceNumber      string    `json:"taxInvoiceNumber,omitempty" binding:"required"`
	DeliveryReceiptNumber string    `json:"deliveryReceiptNumber,omitempty" binding:"required"`
	Status                string    `json:"status,omitempty" binding:"required"`
	ClientID              uuid.UUID `json:"clientId" binding:"required"`
	CreatedBy             uuid.UUID
}

type UpdateOperatingActivity struct {
	ID                    uuid.UUID
	TaxInvoiceNumber      string    `json:"taxInvoiceNumber,omitempty" binding:"required"`
	DeliveryReceiptNumber string    `json:"deliveryReceiptNumber,omitempty" binding:"required"`
	Status                string    `json:"status,omitempty" binding:"required"`
	ClientID              uuid.UUID `json:"clientId" binding:"required"`
	UpdatedBy             uuid.UUID
}
