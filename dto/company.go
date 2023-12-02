package dto

import "github.com/google/uuid"

type InsertCompany struct {
	Name              string `json:"name,omitempty" binding:"required" validate:"required"`
	PhoneNumber       string `json:"phoneNumber,omitempty" binding:"required" validate:"required"`
	Email             string `json:"email,omitempty" binding:"required" validate:"required,email"`
	Address           string `json:"address"`
	PICName           string `json:"PICName,omitempty" binding:"required" validate:"required"`
	PICDesignation    string `json:"PICDesignation" binding:"required" validate:"required"`
	BankAccountName   string `json:"bankAccountName,omitempty" binding:"required" validate:"required"`
	BankAccountNumber int    `json:"bankAccountNumber,omitempty" binding:"required" validate:"required"`
	CreatedBy         uuid.UUID
}

type UpdateCompany struct {
	ID                uuid.UUID `json:"id"`
	Name              string    `json:"name" binding:"required" validate:"required"`
	PhoneNumber       string    `json:"phoneNumber" binding:"required" validate:"required"`
	Email             string    `json:"email" binding:"required" validate:"required,email"`
	Address           string    `json:"address"`
	PICName           string    `json:"PICName" binding:"required" validate:"required"`
	PICDesignation    string    `json:"PICDesignation" binding:"required" validate:"required"`
	BankAccountName   string    `json:"bankAccountName" binding:"required" validate:"required"`
	BankAccountNumber int       `json:"bankAccountNumber" binding:"required" validate:"required"`
	UpdatedBy         uuid.UUID
}
