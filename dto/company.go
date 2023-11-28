package dto

import "github.com/google/uuid"

type InsertCompany struct {
	Name              string `json:"name,omitempty" binding:"required"`
	PhoneNumber       string `json:"phoneNumber,omitempty" binding:"required"`
	Email             string `json:"email,omitempty" binding:"required"`
	Address           string `json:"address"`
	PICName           string `json:"PICName,omitempty" binding:"required"`
	PICDesignation    string `json:"PICDesignation" binding:"required"`
	BankAccountName   string `json:"bankAccountName,omitempty" binding:"required"`
	BankAccountNumber int    `json:"bankAccountNumber,omitempty" binding:"required"`
	CreatedBy         uuid.UUID
}

type UpdateCompany struct {
	ID                uuid.UUID `json:"id"`
	Name              string    `json:"name" binding:"required"`
	PhoneNumber       string    `json:"phoneNumber" binding:"required"`
	Email             string    `json:"email" binding:"required"`
	Address           string    `json:"address"`
	PICName           string    `json:"PICName" binding:"required"`
	PICDesignation    string    `json:"PICDesignation" binding:"required"`
	BankAccountName   string    `json:"bankAccountName" binding:"required"`
	BankAccountNumber int       `json:"bankAccountNumber" binding:"required"`
	UpdatedBy         uuid.UUID
}
