package dto

import (
	"github.com/google/uuid"
	"mime/multipart"
	"time"
)

type InsertFormDataDeliveryOrder struct {
	ContactPerson       string `json:"contactPerson,omitempty" form:"contactPerson" binding:"required"`
	PhoneNumber         string `json:"phoneNumber,omitempty" form:"phoneNumber" binding:"required"`
	Address             string `json:"address,omitempty" form:"address" binding:"required"`
	Note                string `json:"note,omitempty" form:"note" binding:"required"`
	Date                string `json:"date,omitempty" form:"date" binding:"required"`
	Status              string `json:"status,omitempty" form:"status" binding:"required"`
	OperatingActivityID string `json:"operatingActivityId" form:"operatingActivityId" binding:"required"`
	Document            *multipart.FileHeader
	CreatedBy           uuid.UUID
}

type UpdateFormDataDeliveryOrder struct {
	ID                  uuid.UUID
	ContactPerson       string `json:"contactPerson,omitempty" form:"contactPerson" binding:"required"`
	PhoneNumber         string `json:"phoneNumber,omitempty" form:"phoneNumber" binding:"required"`
	Address             string `json:"address,omitempty" form:"address" binding:"required"`
	Note                string `json:"note,omitempty" form:"note" binding:"required"`
	Date                string `json:"date,omitempty" form:"date" binding:"required"`
	Status              string `json:"status,omitempty" form:"status" binding:"required"`
	OperatingActivityID string `json:"operatingActivityId" form:"operatingActivityId" binding:"required"`
	Document            *multipart.FileHeader
	UpdatedBy           uuid.UUID
}

type UpdateDeliveryOrder struct {
	ID                  uuid.UUID
	ContactPerson       string
	PhoneNumber         string
	Address             string
	Note                string
	Date                time.Time
	Status              string
	OperatingActivityID uuid.UUID
	Document            *multipart.FileHeader
	UpdatedBy           uuid.UUID
}

type InsertDeliveryOrder struct {
	ContactPerson       string
	PhoneNumber         string
	Address             string
	Note                string
	Date                time.Time
	Status              string
	OperatingActivityID uuid.UUID
	Document            *multipart.FileHeader
	CreatedBy           uuid.UUID
}
