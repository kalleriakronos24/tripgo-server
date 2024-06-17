package dto

import (
	"github.com/google/uuid"
)

type UserLogin struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// PUBLIC
type UserSignup struct {
	Name      string `json:"name" validate:"required"`
	Address   string `json:"address"`
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email,omitempty" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	CreatedBy uuid.UUID
}

// TENANT
type OwnerSignup struct {
	Name      string `json:"name" validate:"required"`
	Address   string `json:"address"`
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email,omitempty" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	CreatedBy uuid.UUID
}
type AdminSignup struct {
	Name      string `json:"name" validate:"required"`
	Address   string `json:"address"`
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email,omitempty" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	CreatedBy uuid.UUID
}

// INTERNAL
type InternalSignUp struct {
	Name         string `json:"name" validate:"required"`
	Address      string `json:"address"`
	Username     string `json:"username" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required"`
	Phone        string `json:"phone"`
	ProfileImage string `json:"profileImage"`
}

type UserUpdate struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name,omitempty" validate:"required"`
	Address      string    `json:"address"`
	Username     string    `json:"username,omitempty" validate:"required"`
	Email        string    `json:"email" validate:"required,email"`
	Phone        string    `json:"phone"`
	ProfileImage string    `json:"profileImage"`
	Status       string    `json:"status" validate:"required"`
}

type RetrieveUserInfo struct {
	Name     string `json:"name,omitempty" validate:"required"`
	Address  string `json:"address"`
	Username string `json:"username,omitempty" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
	Role     string `json:"role" validate:"required"`
}
