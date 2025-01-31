package dto

type CredentialSignInDto struct {
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
	DeviceToken string `json:"deviceToken"`
	Type        string `json:"type"`
}
