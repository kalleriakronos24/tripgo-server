package dto

type DriverSignUpValidator struct {
	Email       string `json:"email,omitempty" form:"email" validate:"required,email"`
	Password    string `json:"password,omitempty" form:"password" validate:"required"`
	Name        string `json:"name,omitempty" form:"name" validate:"required"`
	Phone       string `json:"phone,omitempty" form:"phone" validate:"required,e164"`
	PlateNumber string `json:"plateNumber,omitempty" form:"plateNumber" validate:"required"`
	DriverType  string `json:"driverType,omitempty" form:"driverType" validate:"required"`
	//LicensePhoto *multipart.FileHeader
}

type DriverSignup struct {
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required"`
	Name       string `json:"name" validate:"required"`
	Phone      string `json:"phone" validate:"required,e164"`
	DriverType string `json:"driverType" validate:"required"`
}
