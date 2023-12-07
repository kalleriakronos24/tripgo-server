package dto

import (
	"github.com/google/uuid"
)

type InsertProductHistoryValidator struct {
	Status              string  `json:"status,omitempty" binding:"required" validate:"required,oneof=increase decrease"`
	Quantity            float64 `json:"quantity,omitempty" binding:"required" validate:"required"`
	ProductID           string  `json:"productId" binding:"required" validate:"required,uuid4"`
	OperatingActivityID string  `json:"operatingActivityId"`
	CreatedBy           uuid.UUID
}

type UpdateProductHistoryValidator struct {
	ID                  uuid.UUID
	Status              string  `json:"status,omitempty" binding:"required" validate:"required,oneof=increase decrease"`
	Quantity            float64 `json:"quantity,omitempty" binding:"required" validate:"required"`
	ProductID           string  `json:"productId" binding:"required" validate:"required,uuid4"`
	OperatingActivityID string  `json:"operatingActivityId"`
	UpdatedBy           uuid.UUID
}

type InsertProductHistory struct {
	Status              string
	Quantity            float64
	ProductID           uuid.UUID
	OperatingActivityID uuid.UUID
	CreatedBy           uuid.UUID
}

type UpdateProductHistory struct {
	ID                  uuid.UUID
	Status              string
	Quantity            float64
	ProductID           uuid.UUID
	OperatingActivityID uuid.UUID
	UpdatedBy           uuid.UUID
}
