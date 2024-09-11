package dto

/*
Name         string    `json:"name" gorm:"not null" binding:"required"`

		BasePrice    int       `json:"basePrice,omitempty" gorm:"not null" binding:"required"`
		LuggageCount int       `json:"luggageCount,omitempty" gorm:"not null" binding:"required"`
		PersonCount  int       `json:"personCount,omitempty" gorm:"not null" binding:"required"`
		Price        int       `json:"price,omitempty" gorm:"not null" binding:"required"`
		PriceType    string    `json:"priceType,omitempty" gorm:"not null" binding:"required"`
		ImagePath    string    `json:"imagePath,omitempty" gorm:"not null" binding:"required"`
		Status       string    `json:"status,omitempty" binding:"required" gorm:"not null;default:active;"`

		 {
	        name: 'Economy',
	        imageUrl: '/images/transport-images/economy.png',
	        isActive: false,
	        activeClass: 'border-2 border-[#cdb23c]',
	        luggageCount: 3,
	        personCount: 3,
	        defaultCurrency: 'RM',
	        price: 0,
	    },
*/
type CreateCarModel struct {
	Name            string `json:"name" validate:"required"`
	ImageUrl        string `json:"imageUrl" validate:"required"`
	IsActive        string `json:"isActive" validate:"required"`
	LuggageCount    string `json:"luggageCount" validate:"required"`
	PersonCount     string `json:"personCount" validate:"required"`
	DefaultCurrency string `json:"defaultCurrency" validate:"required"`
	Price           string `json:"price" validate:"required"`
}
