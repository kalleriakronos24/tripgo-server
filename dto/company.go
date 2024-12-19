package dto

import "mime/multipart"

type InsertFormCompany struct {
	Name               string                     `json:"name,omitempty" binding:"required" validate:"required"`
	PhoneNumber        string                     `json:"phoneNumber,omitempty" binding:"required" validate:"required"`
	Email              string                     `json:"email,omitempty" binding:"required" validate:"required,email"`
	CompanyName        string                     `json:"companyName,omitempty" binding:"required" validate:"required"`
	CompanyAddress     string                     `json:"companyAddress,omitempty" binding:"required" validate:"required"`
	CompanyCountry     string                     `json:"companyCountry,omitempty" binding:"required" validate:"required"`
	CompanyNumber      string                     `json:"companyNumber,omitempty" binding:"required" validate:"required"`
	CompanyCertificate *multipart.FileHeader      `form:"companyCertificate" swaggerignore:"true"`
	Drivers            []*DriverSignUpValidator   `form:"drivers,omitempty" binding:"required" validate:"required,dive,required"`
	Transports         []*InsertFormCarManagement `form:"transports,omitempty" binding:"required" validate:"required,dive,required"`
}

type InsertCompany struct {
	Name               string
	PhoneNumber        string
	Email              string
	CompanyName        string
	CompanyAddress     string
	CompanyCountry     string
	CompanyNumber      string
	CompanyCertificate *multipart.FileHeader
	Drivers            []*DriverSignup
	Transports         []*InsertCarManagement
}
