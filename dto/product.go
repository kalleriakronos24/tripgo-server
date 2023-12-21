package dto

import (
	"github.com/google/uuid"
)

type InsertProductValidator struct {
	Name        string  `json:"name,omitempty" validate:"required"`
	UnitPrice   float64 `json:"unitPrice,omitempty" validate:"required"`
	Packaging   string  `json:"packaging,omitempty" validate:"required"`
	Stock       float64 `json:"stock,omitempty" validate:"required"`
	Note        string  `json:"note,omitempty"`
	VATIncluded bool    `json:"vatIncluded,omitempty" validate:"exist"`
	CompanyID   string  `json:"companyId,omitempty"`
	CreatedBy   uuid.UUID
}

type UpdateProductValidator struct {
	ID          uuid.UUID
	Name        string  `json:"name,omitempty" validate:"required"`
	UnitPrice   float64 `json:"unitPrice,omitempty" validate:"required"`
	Packaging   string  `json:"packaging,omitempty" validate:"required"`
	Stock       float64 `json:"stock,omitempty" validate:"required"`
	Note        string  `json:"note,omitempty"`
	VATIncluded bool    `json:"vatIncluded,omitempty" validate:"exist"`
	CompanyID   string  `json:"companyId,omitempty"`
	UpdatedBy   uuid.UUID
}

type InsertProduct struct {
	Name        string
	UnitPrice   float64
	Packaging   string
	Stock       float64
	Note        string
	VATIncluded bool
	CompanyID   uuid.UUID
	CreatedBy   uuid.UUID
}

type UpdateProduct struct {
	ID          uuid.UUID
	Name        string
	UnitPrice   float64
	Packaging   string
	Stock       float64
	Note        string
	VATIncluded bool
	CompanyID   uuid.UUID
	UpdatedBy   uuid.UUID
}
