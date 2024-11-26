package dto

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type InsertFormDriverTopup struct {
	Amount       string                `form:"amount" binding:"required" validate:"required"`
	ReceiptImage *multipart.FileHeader `form:"receiptImage" swaggerignore:"true"`
	DriverID     uuid.UUID
}

type InsertDriverTopup struct {
	Amount       float64
	ReceiptImage *multipart.FileHeader
	DriverID     uuid.UUID
}
