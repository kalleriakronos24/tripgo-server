package dto

import "github.com/google/uuid"

type UpdateDriverTopupHistory struct {
	DriverTopupID uuid.UUID `json:"driverTopupId" validate:"required"`
	Amount        float64   `json:"amount" validate:"required"`
	Status        string    `json:"status" validate:"required"`
	DriverID      uuid.UUID
}
