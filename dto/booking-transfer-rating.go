package dto

import (
	"github.com/google/uuid"
)

type InsertBookingTransferRating struct {
	CustomerID                uuid.UUID `json:"customerId"`
	BookingTransferAssignedID uuid.UUID `json:"bookingTransferAssignedID"`
	CarManagementId           uuid.UUID `json:"carManagementId"`
	Rating                    int       `json:"rating,omitempty"`
}

type UpdateBookingTransferRating struct {
	CustomerID                uuid.UUID `json:"customerId"`
	BookingTransferAssignedID uuid.UUID `json:"bookingTransferAssignedID"`
	CarManagementId           uuid.UUID `json:"carManagementId"`
	Rating                    int       `json:"rating,omitempty"`
}
