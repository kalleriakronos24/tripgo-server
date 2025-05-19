package dto

import (
	"time"

	"github.com/google/uuid"
)

type InsertBookingTransfer struct {
	CustomerID        uuid.UUID `json:"id"`
	AdultSeater       int       `json:"adultSeater,omitempty"`
	ChildSeater       int       `json:"childSeater,omitempty"`
	FromLatCoordinate float32   `json:"fromLatCoordinate,omitempty"`
	FromLngCoordinate float32   `json:"fromLngCoordinate,omitempty"`
	ToLatCoordinate   float32   `json:"toLatCoordinate,omitempty"`
	ToLngCoordinate   float32   `json:"toLngCoordinate,omitempty"`
	FromLocation      string    `json:"fromLocation,omitempty"`
	ToLocation        string    `json:"toLocation,omitempty"`
	TotalDistance     float32   `json:"totalDistance,omitempty"`
	PassengerNotes    string    `json:"passengerNotes,omitempty"`
	PickUpDate        time.Time `json:"pickUpDate,omitempty"`
	CarModelID        uuid.UUID `json:"carModelId"`
	Price             float32   `json:"price,omitempty"`
	GrandTotal        float32   `json:"grandTotal,omitempty"`
	AddPickupPoint    int       `json:"addPickupPoint,omitempty"`
	AddDropoffPoint   int       `json:"addDropoffPoint,omitempty"`
	RefferalCode      string    `json:"refferalCode,omitempty"`
	PaymentOption     string    `json:"paymentOption,omitempty"`
	PI                string    `json:"pi,omitempty"`
	Currency          string    `json:"currency,omitempty"`
}

type UpdateBookingTransfer struct {
	ID                uuid.UUID
	AdultSeater       int       `json:"adultSeater,omitempty" validate:"required"`
	ChildSeater       int       `json:"childSeater,omitempty" validate:"required"`
	FromLatCoordinate float32   `json:"fromLatCoordinate,omitempty" validate:"required"`
	FromLngCoordinate float32   `json:"fromLngCoordinate,omitempty" validate:"required"`
	ToLatCoordinate   float32   `json:"toLatCoordinate,omitempty" validate:"required"`
	ToLngCoordinate   float32   `json:"toLngCoordinate,omitempty" validate:"required"`
	FromLocation      string    `json:"fromLocation,omitempty" validate:"required"`
	ToLocation        string    `json:"toLocation,omitempty" validate:"required"`
	TotalDistance     float32   `json:"totalDistance,omitempty"`
	PassengerNotes    string    `json:"passengerNotes,omitempty" validate:"required"`
	PickUpDate        time.Time `json:"pickUpDate,omitempty" validate:"required"`
	CarModelID        uuid.UUID `json:"carModelId"`
	Price             float32   `json:"price,omitempty" validate:"required"`
	GrandTotal        float32   `json:"grandTotal,omitempty"`
	AddPickupPoint    int       `json:"addPickupPoint,omitempty"`
	AddDropoffPoint   int       `json:"addDropoffPoint,omitempty"`
	RefferalCode      string    `json:"refferalCode,omitempty"`
	PaymentOption     string    `json:"paymentOption,omitempty"`
	PI                string    `json:"pi,omitempty"`
}
