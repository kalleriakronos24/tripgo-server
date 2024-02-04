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
	CreatedBy           uuid.UUID
}

type UpdateDeliveryOrderValidator struct {
	ID                  uuid.UUID
	ContactPerson       string `json:"contactPerson,omitempty" binding:"required" validate:"required"`
	PhoneNumber         string `json:"phoneNumber,omitempty" binding:"required" validate:"required"`
	Address             string `json:"address,omitempty" binding:"required" validate:"required"`
	Note                string `json:"note,omitempty" binding:"required" validate:"required"`
	Date                string `json:"date,omitempty" binding:"required" validate:"required,ISO8601date"`
	Status              string `json:"status,omitempty" binding:"required" validate:"required,oneof=shipping delivery canceled"`
	OperatingActivityID string `json:"operatingActivityId,omitempty" binding:"required" validate:"uuid4"`
	UpdatedBy           uuid.UUID
}

type InsertDeliveryOrderValidator struct {
	ContactPerson       string `json:"contactPerson,omitempty" binding:"required" validate:"required"`
	PhoneNumber         string `json:"phoneNumber,omitempty" binding:"required" validate:"required"`
	Address             string `json:"address,omitempty" binding:"required" validate:"required"`
	Note                string `json:"note,omitempty" binding:"required" validate:"required"`
	Date                string `json:"date,omitempty" binding:"required" validate:"required,ISO8601date"`
	Status              string `json:"status,omitempty" binding:"required" validate:"required,oneof=shipping delivery canceled"`
	OperatingActivityID string `json:"operatingActivityId,omitempty" binding:"required" validate:"uuid4"`
	CreatedBy           uuid.UUID
}
