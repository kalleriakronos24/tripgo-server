package dto

import (
	"github.com/google/uuid"
)

type InsertProductHistory struct {
	Status   string `json:"status,omitempty" binding:"required"`
	Quantity int8   `json:"quantity,omitempty" binding:"required"`

	ProductID           uuid.UUID `json:"productId" binding:"required"`
	OperatingActivityID uuid.UUID `json:"operatingActivityId" binding:"required"`
	CreatedBy           uuid.UUID
}

type UpdateProductHistory struct {
	ID       uuid.UUID
	Status   string `json:"status,omitempty" binding:"required"`
	Quantity int8   `json:"quantity,omitempty" binding:"required"`

	ProductID           uuid.UUID `json:"productId" binding:"required"`
	OperatingActivityID uuid.UUID `json:"operatingActivityId" binding:"required"`
	UpdatedBy           uuid.UUID
}
