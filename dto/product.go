package dto

import (
	"github.com/google/uuid"
)

type InsertProduct struct {
	Name      string    `json:"name,omitempty" binding:"required"`
	UnitPrice float64   `json:"unitPrice,omitempty" binding:"required"`
	Packaging string    `json:"packaging,omitempty" binding:"required"`
	Stock     int8      `json:"stock,omitempty" binding:"required"`
	Note      string    `json:"note,omitempty"`
	CompanyID uuid.UUID `json:"companyId" binding:"required"`
	CreatedBy uuid.UUID
}

type UpdateProduct struct {
	ID        uuid.UUID
	Name      string    `json:"name,omitempty" binding:"required"`
	UnitPrice float64   `json:"unitPrice,omitempty" binding:"required"`
	Packaging string    `json:"packaging,omitempty" binding:"required"`
	Stock     int8      `json:"stock,omitempty" binding:"required"`
	Note      string    `json:"note,omitempty"`
	CompanyID uuid.UUID `json:"companyId" binding:"required"`
	UpdatedBy uuid.UUID
}
