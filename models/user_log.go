package models

import (
	"github.com/google/uuid"
	masterModels "github.com/kalleriakronos24/booklap-be/models/master"
	"github.com/kalleriakronos24/booklap-be/types"
)

// type userLogOrm struct {
// 	db *gorm.DB
// }

type UserLog struct {
	ID          uint   `gorm:"primaryKey" json:"-"`
	Action      string `json:"-"`
	Description string
	Kind        string // common / error

	UserID uuid.UUID
	User   masterModels.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID"`

	types.DefaultModelProperty
}
