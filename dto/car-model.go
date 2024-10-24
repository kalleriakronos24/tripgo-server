package dto

type CreateCarModel struct {
	Name            string `json:"name" validate:"required"`
	ImageUrl        string `json:"imageUrl" validate:"required"`
	IsActive        string `json:"isActive" validate:"required"`
	LuggageCount    string `json:"luggageCount" validate:"required"`
	PersonCount     string `json:"personCount" validate:"required"`
	DefaultCurrency string `json:"defaultCurrency" validate:"required"`
	Price           string `json:"price" validate:"required"`
}
