package dto

import "mime/multipart"

type InsertFormCompany struct {
	Name               string                   `form:"name" json:"name,omitempty" binding:"required" validate:"required"`
	PhoneNumber        string                   `form:"phoneNumber" json:"phoneNumber,omitempty" binding:"required" validate:"required"`
	Email              string                   `form:"email" json:"email,omitempty" binding:"required" validate:"required,email"`
	CompanyName        string                   `form:"companyName" json:"companyName,omitempty" binding:"required" validate:"required"`
	CompanyAddress     string                   `form:"companyAddress" json:"companyAddress,omitempty" binding:"required" validate:"required"`
	CompanyCountry     string                   `form:"companyCountry" json:"companyCountry,omitempty" binding:"required" validate:"required"`
	CompanyNumber      string                   `form:"companyNumber" json:"companyNumber,omitempty" binding:"required" validate:"required"`
	CompanyCertificate *multipart.FileHeader    `form:"companyCertificate" swaggerignore:"true"`
	Drivers            []*DriverSignUpValidator `json:"drivers,omitempty" binding:"required" validate:"required,dive,required"`
	Transports         []*InsertCarManagement   `json:"transports,omitempty" binding:"required" validate:"required,dive,required"`
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
}
