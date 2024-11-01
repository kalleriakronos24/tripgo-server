package dto

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type InsertFormCarManagement struct {
	Name              string                `form:"name" binding:"required" validate:"required"`
	PlateNumber       string                `form:"plateNumber" binding:"required" validate:"required"`
	FrontCarPhoto     *multipart.FileHeader `swaggerignore:"true"`
	LicensePhoto      string                `form:"licensePhoto"`
	CarManagementType string                `form:"carManagementType" binding:"required" validate:"required"`
	CarModelID        string                `form:"carModelId" binding:"required" validate:"required,uuid4"`
	DriverID          uuid.UUID
}

type InsertCarManagement struct {
	Name              string
	PlateNumber       string
	FrontCarPhoto     *multipart.FileHeader
	LicensePhoto      string
	CarManagementType string
	CarModelID        uuid.UUID
	DriverID          uuid.UUID
}

type UpdateFormCarManagement struct {
	Name          string                `form:"name" binding:"required" validate:"required"`
	PlateNumber   string                `form:"plateNumber" binding:"required" validate:"required"`
	FrontCarPhoto *multipart.FileHeader `swaggerignore:"true"`
	LicensePhoto  string                `form:"licensePhoto"`
	CarModelID    string                `form:"carModelId" binding:"required" validate:"required,uuid4"`
}

type UpdateCarManagement struct {
	Name          string
	PlateNumber   string
	FrontCarPhoto *multipart.FileHeader
	LicensePhoto  string
	CarModelID    uuid.UUID
}
