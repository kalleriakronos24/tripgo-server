package dto

import (
	"github.com/google/uuid"
	"mime/multipart"
	"time"
)

type InsertFormDataPurchaseOrder struct {
	Number              string `form:"number" binding:"required" validate:"required"`
	Type                string `form:"type" binding:"required" validate:"required,oneof=in out"`
	Recipient           string `form:"recipient" binding:"required" validate:"required"`
	RecipientEmail      string `form:"recipientEmail" binding:"required" validate:"required,email"`
	Date                string `form:"date" binding:"required" validate:"required,ISO8601date"`
	OperatingActivityID string `form:"operatingActivityId" binding:"required" validate:"required,uuid4"`
	Document            *multipart.FileHeader
	CreatedBy           uuid.UUID
}

type UpdateFormDataPurchaseOrder struct {
	ID                  uuid.UUID
	Number              string `form:"number" binding:"required" validate:"required"`
	Type                string `form:"type" binding:"required" validate:"required,oneof=in out"`
	Recipient           string `form:"recipient" binding:"required" validate:"required"`
	RecipientEmail      string `form:"recipientEmail" binding:"required" validate:"required,email"`
	Date                string `form:"date" binding:"required" validate:"required,ISO8601date"`
	OperatingActivityID string `form:"operatingActivityId" binding:"required" validate:"required,uuid4"`
	Document            *multipart.FileHeader
	UpdatedBy           uuid.UUID
}

type InsertPurchaseOrder struct {
	Number              string
	Type                string
	Recipient           string
	RecipientEmail      string
	Date                time.Time
	OperatingActivityID uuid.UUID
	Document            *multipart.FileHeader
	CreatedBy           uuid.UUID
}

type UpdatePurchaseOrder struct {
	ID                  uuid.UUID
	Number              string
	Type                string
	Recipient           string
	RecipientEmail      string
	Date                time.Time
	OperatingActivityID uuid.UUID
	Document            *multipart.FileHeader
	UpdatedBy           uuid.UUID
}
