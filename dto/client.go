package dto

import (
	"github.com/google/uuid"
)

type InsertClientValidator struct {
	Name        string `json:"name" binding:"required" validate:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required" validate:"required,e164"`
	Email       string `json:"email" binding:"required" validate:"required,email"`
	Address     string `json:"address,omitempty"`
	CreatedBy   uuid.UUID
	CompanyID   string `json:"companyId" validate:"required,uuid4"`
}

type UpdateClientValidator struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name" binding:"required"  validate:"required"`
	PhoneNumber string    `json:"phoneNumber" binding:"required"  validate:"required,e164"`
	Email       string    `json:"email" binding:"required"  validate:"required,email"`
	Address     string    `json:"address,omitempty"`
	UpdatedBy   uuid.UUID
	CompanyID   string `json:"companyId" binding:"required"  validate:"required,uuid4"`
}

type InsertClient struct {
	Name        string
	PhoneNumber string
	Email       string
	Address     string
	CreatedBy   uuid.UUID
	CompanyID   uuid.UUID
}

type UpdateClient struct {
	ID          uuid.UUID
	Name        string
	PhoneNumber string
	Email       string
	Address     string
	UpdatedBy   uuid.UUID
	CompanyID   uuid.UUID
}
