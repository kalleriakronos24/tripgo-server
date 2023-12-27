package dto

import (
	"github.com/google/uuid"
	models "gitlab.com/odma1/odma-be/models/master"
)

type UserLogin struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserSignup struct {
	Name      string `json:"name" validate:"required"`
	Address   string `json:"address"`
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email,omitempty" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	Role      string `json:"role" validate:"required"`
	CompanyID string `json:"companyId" binding:"required" validate:"required,uuid4"`
	CreatedBy uuid.UUID
}

type UserSignupSuperAdmin struct {
	Name     string `json:"name" validate:"required"`
	Address  string `json:"address"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserUpdate struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name,omitempty" validate:"required"`
	Address   string    `json:"address"`
	Username  string    `json:"username,omitempty" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Role      string    `json:"role" validate:"required"`
	Status    string    `json:"status" validate:"required"`
	CompanyID string    `json:"companyId" binding:"required" validate:"required,uuid4"`
}

type RetrieveUserInfo struct {
	Name      string `json:"name,omitempty" validate:"required"`
	Address   string `json:"address"`
	Username  string `json:"username,omitempty" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password,omitempty" validate:"required"`
	CreatedBy models.User
	UpdatedBy models.User
}
