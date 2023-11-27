package dto

import (
	"github.com/google/uuid"
	models "gitlab.com/odma1/odma-be/models/master"
)

type InsertClient struct {
	Name        string `json:"name" gorm:"not null;default:NULL" `
	PhoneNumber string `json:"phoneNumber" gorm:"not null;default:NULL"`
	Email       string `gorm:"unique;not null;default:NULL" json:"email"`
	Address     string `json:"address,omitempty" gorm:"default:NULL"`
	CreatedBy   uuid.UUID
}

type UpdateClient struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	Name        string    `json:"name" gorm:"not null;default:NULL" `
	PhoneNumber string    `json:"phoneNumber" gorm:"not null;default:NULL"`
	Email       string    `gorm:"unique;not null;default:NULL" json:"email"`
	Address     string    `json:"address,omitempty" gorm:"default:NULL"`
	UpdatedBy   uuid.UUID
}

type RetrieveClient struct {
	Name        string `json:"name" gorm:"not null;default:NULL" `
	PhoneNumber string `json:"phoneNumber" gorm:"not null;default:NULL"`
	Email       string `gorm:"unique;not null;default:NULL" json:"email"`
	Address     string `json:"address,omitempty" gorm:"default:NULL"`
	UpdatedBy   models.User
	CreatedBy   models.User
}
