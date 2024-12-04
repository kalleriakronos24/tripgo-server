package dto

import "github.com/google/uuid"

type ValidatorUpdateDriverTopupHistory struct {
	DriverTopupID string  `json:"driverTopupId" validate:"required,uuid4"`
	DriverID      string  `json:"driverId" validate:"required,uuid4"`
	Amount        float64 `json:"amount" validate:"required"`
	Status        string  `json:"status" validate:"required"`
}

type UpdateDriverTopupHistory struct {
	DriverTopupID uuid.UUID
	Amount        float64
	Status        string
	DriverID      uuid.UUID
}
