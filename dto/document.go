package dto

import (
	"github.com/google/uuid"
)

type InsertDocument struct {
	Base64       string `json:"base64,omitempty" binding:"required"`
	Path         string `json:"path,omitempty"  binding:"required"`
	AbsolutePath string `json:"absolutePath,omitempty" binding:"required"`
	FileName     string `json:"filename,omitempty"  binding:"required"`
	Extension    string `json:"extension,omitempty"  binding:"required"`
	Location     string `json:"location,omitempty"  binding:"required"`
	CreatedBy    uuid.UUID
}

type UpdateDocument struct {
	ID           uuid.UUID
	Base64       string `json:"base64,omitempty" binding:"required"`
	Path         string `json:"path,omitempty"  binding:"required"`
	AbsolutePath string `json:"absolutePath,omitempty" binding:"required"`
	FileName     string `json:"filename,omitempty"  binding:"required"`
	Extension    string `json:"extension,omitempty"  binding:"required"`
	Location     string `json:"location,omitempty"  binding:"required"`
	UpdatedBy    uuid.UUID
}
