package dto

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type InsertFormCarManagement struct {
	Name        string `form:"name" json:"name" binding:"required" validate:"required"`
	PlateNumber string `form:"plateNumber" json:"plateNumber" binding:"required" validate:"required"`
	// FrontCarPhoto     *multipart.FileHeader `json:"frontCarPhoto" swaggerignore:"true"`
	// LicensePhoto      *multipart.FileHeader `json:"licensePhoto" swaggerignore:"true"`
	// RoadTaxPhoto      *multipart.FileHeader `json:"roadTaxPhoto" swaggerignore:"true"`
	// VEPPhoto          *multipart.FileHeader `json:"vepPhoto" swaggerignore:"true"`
	CarManagementType string `form:"carManagementType" json:"carManagementType" binding:"required" validate:"required"`
	CarModelID        string `form:"carModelId" json:"carModelId" binding:"required" validate:"required,uuid4"`
	// DriverID          uuid.UUID
}

type InsertCarManagement struct {
	Name        string
	PlateNumber string
	// FrontCarPhoto     *multipart.FileHeader
	// LicensePhoto      *multipart.FileHeader
	// RoadTaxPhoto      *multipart.FileHeader
	//  VEPPhoto          *multipart.FileHeader
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
