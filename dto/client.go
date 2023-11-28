package dto

import (
	"github.com/google/uuid"
)

type InsertClient struct {
	Name        string `json:"name" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Address     string `json:"address" gorm:"default:NULL"`
	CreatedBy   uuid.UUID
	CompanyID   uuid.UUID `json:"companyId" binding:"required"`
}

type UpdateClient struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name" binding:"required"`
	PhoneNumber string    `json:"phoneNumber" binding:"required"`
	Email       string    `json:"email" binding:"required"`
	Address     string    `json:"address,omitempty"`
	UpdatedBy   uuid.UUID
	CompanyID   uuid.UUID `json:"companyId" binding:"required"`
}
