package dto

import "github.com/google/uuid"

type InsertCompany struct {
	Name              string `json:"name,omitempty" gorm:"not null"`
	PhoneNumber       string `json:"phoneNumber,omitempty" gorm:"not null"`
	Email             string `gorm:"unique, not null" json:"email,omitempty"`
	Address           string `json:"address"`
	PICName           string `json:"PICName,omitempty" gorm:"not null"`
	PICDesignation    string `json:"PICDesignation" gorm:"not null"`
	BankAccountName   string `json:"bankAccountName,omitempty" gorm:"not null"`
	BankAccountNumber int    `json:"bankAccountNumber,omitempty" gorm:"not null"`
	CreatedBy         uuid.UUID
}

type UpdateCompany struct {
	ID                uuid.UUID `json:"id"`
	Name              string    `json:"name,omitempty" gorm:"not null"`
	PhoneNumber       string    `json:"phoneNumber,omitempty" gorm:"not null"`
	Email             string    `gorm:"unique, not null" json:"email,omitempty"`
	Address           string    `json:"address"`
	PICName           string    `json:"picName,omitempty" gorm:"not null"`
	PICDesignation    string    `json:"picDesignation" gorm:"not null"`
	BankAccountName   string    `json:"bankAccountName,omitempty" gorm:"not null"`
	BankAccountNumber int       `json:"bankAccountNumber,omitempty" gorm:"not null"`
	UpdatedBy         uuid.UUID
}

type RetrieveCompany struct {
	ID                uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()" json:"id"`
	Name              string    `json:"name,omitempty" gorm:"not null"`
	PhoneNumber       string    `json:"phoneNumber,omitempty" gorm:"not null"`
	Email             string    `gorm:"unique, not null" json:"email,omitempty"`
	Address           string    `json:"address"`
	PICName           string    `json:"picName,omitempty" gorm:"not null"`
	PICDesignation    string    `json:"picDesignation" gorm:"not null"`
	BankAccountName   string    `json:"bankAccountName,omitempty" gorm:"not null"`
	BankAccountNumber int       `json:"bankAccountNumber,omitempty" gorm:"not null"`
}
