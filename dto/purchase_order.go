package dto

import (
	"github.com/google/uuid"
	"mime/multipart"
	"time"
)

type InsertPurchaseOrder struct {
	Number              string    `json:"number,omitempty" form:"number" binding:"required"`
	Type                string    `json:"type,omitempty" form:"type" binding:"required"`
	Recipient           string    `json:"recipient,omitempty" form:"recipient" binding:"required"`
	RecipientEmail      string    `json:"recipientEmail,omitempty" form:"recipientEmail" binding:"required"`
	Date                time.Time `json:"date,omitempty" form:"date" binding:"required"`
	OperatingActivityID uuid.UUID `json:"operatingActivityId" form:"operatingActivityId" binding:"required"`
	Document            *multipart.FileHeader
	CreatedBy           uuid.UUID
}

type InsertFormDataPurchaseOrder struct {
	Number              string    `json:"number,omitempty" form:"number" binding:"required"`
	Type                string    `json:"type,omitempty" form:"type" binding:"required"`
	Recipient           string    `json:"recipient,omitempty" form:"recipient" binding:"required"`
	RecipientEmail      string    `json:"recipientEmail,omitempty" form:"recipientEmail" binding:"required"`
	Date                time.Time `json:"date,omitempty" form:"date" binding:"required"`
	OperatingActivityID string    `json:"operatingActivityId" form:"operatingActivityId" binding:"required"`
	Document            *multipart.FileHeader
	CreatedBy           uuid.UUID
}

type UpdatePurchaseOrder struct {
	ID                  uuid.UUID
	Number              string    `json:"number,omitempty" binding:"required"`
	Type                string    `json:"type,omitempty" binding:"required"`
	Recipient           string    `json:"recipient,omitempty" binding:"required"`
	RecipientEmail      string    `json:"recipientEmail,omitempty" binding:"required"`
	Date                time.Time `json:"date,omitempty" binding:"required"`
	OperatingActivityID uuid.UUID `json:"operatingActivityId" binding:"required"`
	Document            *multipart.FileHeader
	UpdatedBy           uuid.UUID
}

type UpdateFormDataPurchaseOrder struct {
	ID                  uuid.UUID
	Number              string    `json:"number,omitempty" form:"number" binding:"required"`
	Type                string    `json:"type,omitempty" form:"type" binding:"required"`
	Recipient           string    `json:"recipient,omitempty" form:"recipient" binding:"required"`
	RecipientEmail      string    `json:"recipientEmail,omitempty" form:"recipientEmail" binding:"required"`
	Date                time.Time `json:"date,omitempty" form:"date" binding:"required"`
	OperatingActivityID string    `json:"operatingActivityId" form:"operatingActivityId" binding:"required"`
	Document            *multipart.FileHeader
	UpdatedBy           uuid.UUID
}
