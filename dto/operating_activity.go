package dto

import (
	"github.com/google/uuid"
)

type InsertOperatingActivityValidator struct {
	TaxInvoiceNumber      string `json:"taxInvoiceNumber,omitempty" binding:"required" validate:"required"`
	DeliveryReceiptNumber string `json:"deliveryReceiptNumber,omitempty" binding:"required" validate:"required"`
	Status                string `json:"status,omitempty" binding:"required" validate:"required,oneof=ongoing done canceled"`
	ClientID              string `json:"clientId" binding:"required" validate:"required,uuid4"`
	CreatedBy             uuid.UUID
}

type UpdateOperatingActivityValidator struct {
	ID                    uuid.UUID
	TaxInvoiceNumber      string `json:"taxInvoiceNumber,omitempty" binding:"required" validate:"required"`
	DeliveryReceiptNumber string `json:"deliveryReceiptNumber,omitempty" binding:"required" validate:"required"`
	Status                string `json:"status,omitempty" binding:"required" validate:"required,oneof=ongoing done canceled"`
	ClientID              string `json:"clientId" binding:"required" validate:"required,uuid4"`
	UpdatedBy             uuid.UUID
}

type InsertOperatingActivity struct {
	TaxInvoiceNumber      string
	DeliveryReceiptNumber string
	Status                string
	ClientID              uuid.UUID
	CreatedBy             uuid.UUID
}

type UpdateOperatingActivity struct {
	ID                    uuid.UUID
	TaxInvoiceNumber      string
	DeliveryReceiptNumber string
	Status                string
	ClientID              uuid.UUID
	UpdatedBy             uuid.UUID
}
