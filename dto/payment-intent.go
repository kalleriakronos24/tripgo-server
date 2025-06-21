package dto

type CreatePaymentIntentValidator struct {
	Amount   int64  `json:"amount" binding:"required"  validate:"required"`
	Currency string `json:"currency" binding:"required"  validate:"required"`
}

type PaymentIntent struct {
	Amount   int64
	Currency string
}
