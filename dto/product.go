package dto

import (
	"github.com/google/uuid"
)

type InsertProductValidator struct {
	Name      string  `json:"name,omitempty" validate:"required"`
	UnitPrice float64 `json:"unitPrice,omitempty" validate:"required"`
	Packaging string  `json:"packaging,omitempty" validate:"required"`
	Stock     int8    `json:"stock,omitempty" validate:"required"`
	Note      string  `json:"note,omitempty"`
	CompanyID string  `json:"companyId" validate:"required,uuid4"`
	CreatedBy uuid.UUID
}

type UpdateProductValidator struct {
	ID        uuid.UUID
	Name      string  `json:"name,omitempty" validate:"required"`
	UnitPrice float64 `json:"unitPrice,omitempty" validate:"required"`
	Packaging string  `json:"packaging,omitempty" validate:"required"`
	Stock     int8    `json:"stock,omitempty" validate:"required"`
	Note      string  `json:"note,omitempty"`
	CompanyID string  `json:"companyId" validate:"required,uuid4"`
	UpdatedBy uuid.UUID
}

type InsertProduct struct {
	Name      string
	UnitPrice float64
	Packaging string
	Stock     int8
	Note      string
	CompanyID uuid.UUID
	CreatedBy uuid.UUID
}

type UpdateProduct struct {
	ID        uuid.UUID
	Name      string
	UnitPrice float64
	Packaging string
	Stock     int8
	Note      string
	CompanyID uuid.UUID
	UpdatedBy uuid.UUID
}
