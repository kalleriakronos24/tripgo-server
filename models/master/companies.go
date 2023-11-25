package models

import (
	"github.com/google/uuid"
	"gitlab.com/odma1/odma-be/types"
	"gorm.io/gorm"
)

type companyOrm struct {
	db *gorm.DB
}

type Company struct {
	ID                uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Name              string    `json:"name,omitempty"`
	PhoneNumber       string    `json:"phoneNumber,omitempty"`
	Email             string    `gorm:"unique" json:"email,omitempty"`
	Address           string    `json:"address"`
	PICName           string    `json:"PICName,omitempty"`
	PICDesignation    string    `json:"PICDesignation"`
	BankAccountName   string    `json:"bankAccountName,omitempty"`
	BankAccountNumber string    `json:"bankAccountNumber,omitempty"`
	CreatedBy         uuid.UUID `json:"createdBy,omitempty"`

	User User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:CreatedBy"`
	types.DefaultModelProperty
}
