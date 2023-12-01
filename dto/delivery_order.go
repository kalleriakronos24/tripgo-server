package dto

import (
	"github.com/google/uuid"
	"mime/multipart"
	"time"
)

type InsertDeliveryOrder struct {
	Number              string    `json:"number,omitempty" binding:"required"`
	ContactPerson       string    `json:"contactPerson,omitempty" binding:"required"`
	PhoneNumber         string    `json:"phoneNumber,omitempty" binding:"required"`
	Address             string    `json:"address,omitempty" binding:"required"`
	Note                string    `json:"note,omitempty" binding:"required"`
	Date                time.Time `json:"date,omitempty" binding:"required"`
	Status              string    `json:"status,omitempty" binding:"required"`
	OperatingActivityID uuid.UUID `json:"operatingActivityId" binding:"required"`
	Document            *multipart.FileHeader
	CreatedBy           uuid.UUID
}

type InsertFormDataDeliveryOrder struct {
	Number              string    `json:"number,omitempty" form:"number" binding:"required"`
	ContactPerson       string    `json:"contactPerson,omitempty" form:"contactPerson" binding:"required"`
	PhoneNumber         string    `json:"phoneNumber,omitempty" form:"phoneNumber" binding:"required"`
	Address             string    `json:"address,omitempty" form:"address" binding:"required"`
	Note                string    `json:"note,omitempty" form:"note" binding:"required"`
	Date                time.Time `json:"date,omitempty" form:"date" binding:"required"`
	Status              string    `json:"status,omitempty" form:"status" binding:"required"`
	OperatingActivityID string    `json:"operatingActivityId" form:"operatingActivityId" binding:"required"`
	Document            *multipart.FileHeader
	CreatedBy           uuid.UUID
}

type UpdateDeliveryOrder struct {
	ID                  uuid.UUID
	Number              string    `json:"number,omitempty" binding:"required"`
	ContactPerson       string    `json:"contactPerson,omitempty" binding:"required"`
	PhoneNumber         string    `json:"phoneNumber,omitempty" binding:"required"`
	Address             string    `json:"address,omitempty" binding:"required"`
	Note                string    `json:"note,omitempty" binding:"required"`
	Date                time.Time `json:"date,omitempty" binding:"required"`
	Status              string    `json:"status,omitempty" binding:"required"`
	OperatingActivityID uuid.UUID `json:"operatingActivityId" binding:"required"`
	Document            *multipart.FileHeader
	UpdatedBy           uuid.UUID
}

type UpdateFormDataDeliveryOrder struct {
	ID                  uuid.UUID
	Number              string    `json:"number,omitempty" form:"number" binding:"required"`
	ContactPerson       string    `json:"contactPerson,omitempty" form:"contactPerson" binding:"required"`
	PhoneNumber         string    `json:"phoneNumber,omitempty" form:"phoneNumber" binding:"required"`
	Address             string    `json:"address,omitempty" form:"address" binding:"required"`
	Note                string    `json:"note,omitempty" form:"note" binding:"required"`
	Date                time.Time `json:"date,omitempty" form:"date" binding:"required"`
	Status              string    `json:"status,omitempty" form:"status" binding:"required"`
	OperatingActivityID string    `json:"operatingActivityId" form:"operatingActivityId" binding:"required"`
	Document            *multipart.FileHeader
	UpdatedBy           uuid.UUID
}
