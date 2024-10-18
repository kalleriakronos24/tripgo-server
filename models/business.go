package models

import (
	"github.com/google/uuid"
	"github.com/kalleriakronos24/khaimal-group/types"
)

type Business struct {
	ID           uuid.UUID `gorm:"index:id,unique;type:uuid;default:gen_random_uuid();" json:"id"`
	Name         string    `json:"name" gorm:"not null" binding:"required"`
	Location     string    `gorm:"location" json:",omitempty" binding:"required"`
	Address      string    `json:"address,omitempty" binding:"required" gorm:"not null"`
	WorkingHours string    `json:"workingHours,omitempty" binding:"required" gorm:"not null;"`
	ProfileImage string    `json:"profileImage,omitempty" gorm:"default:NULL"`
	Phone        string    `json:"phone,omitempty" gorm:"default:NULL"`
	Status       string    `json:"status,omitempty" binding:"required" gorm:"not null;default:active;"`

	CreatedBy uuid.UUID `json:"createdBy,omitempty" gorm:"type:uuid;default:NULL"`
	UpdatedBy uuid.UUID `json:"updatedBy,omitempty" gorm:"type:uuid;default:NULL"`

	types.DefaultModelProperty
}
