package models

import (
	"github.com/google/uuid"
	masterModels "gitlab.com/odma1/odma-be/models/master"
	"gitlab.com/odma1/odma-be/types"
	"gorm.io/gorm"
)

type userLogOrm struct {
	db *gorm.DB
}

type UserLog struct {
	ID          uint   `gorm:"primaryKey" json:"-"`
	Action      string `json:"-"`
	Description string
	Kind        string // common / error

	UserID uuid.UUID
	User   masterModels.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID"`

	types.DefaultModelProperty
}
