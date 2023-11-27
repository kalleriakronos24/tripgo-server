package dto

import (
	"github.com/google/uuid"
	models "gitlab.com/odma1/odma-be/models/master"
)

type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserSignup struct {
	Name      string `json:"name" binding:"required"`
	Address   string `json:"address"`
	Username  string `json:"username" binding:"required"`
	Email     string `gorm:"unique" json:"email,omitempty" binding:"required"`
	Password  string `json:"password" binding:"required"`
	CreatedBy uuid.UUID
}

type UserSignupSuperAdmin struct {
	Name     string `json:"name" binding:"required"`
	Address  string `json:"address"`
	Username string `json:"username" binding:"required"`
	Email    string `gorm:"unique" json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserUpdate struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name,omitempty" binding:"required"`
	Address  string    `json:"address"`
	Username string    `json:"username,omitempty" binding:"required"`
	Email    string    `gorm:"unique" json:"email" binding:"required"`
}

type RetrieveUserInfo struct {
	Name      string `json:"name,omitempty" binding:"required"`
	Address   string `json:"address"`
	Username  string `json:"username,omitempty" binding:"required"`
	Email     string `gorm:"unique" json:"email" binding:"required"`
	Password  string `json:"password,omitempty" binding:"required"`
	CreatedBy models.User
	UpdatedBy models.User
}
