package dto

type CustomerSignup struct {
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required"`
	Name         string `json:"name" validate:"required"`
	Phone        string `json:"phone"`
	RefferalCode string `json:"refferalCode"`
}
