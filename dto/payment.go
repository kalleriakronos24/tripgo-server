package dto

import (
	"github.com/google/uuid"
	"time"
)

type InsertPaymentValidator struct {
	Currency            string  `json:"currency,omitempty" binding:"required" validate:"required"`
	Amount              float64 `json:"amount,omitempty" binding:"required" validate:"required"`
	Term                string  `json:"term,omitempty" binding:"required" validate:"required"`
	ARAging             int16   `json:"araging,omitempty" binding:"required" validate:"required"`
	Remarks             string  `json:"remarks,omitempty" binding:"required" validate:"required"`
	DueDate             string  `json:"dueDate,omitempty" binding:"required" validate:"required,ISO8601date"`
	SettlementDate      string  `json:"settlementDate,omitempty" binding:"required" validate:"required,ISO8601date"`
	OperatingActivityID string  `json:"operatingActivityId" binding:"required" validate:"required,uuid4"`
	CreatedBy           uuid.UUID
}

type UpdatePaymentValidator struct {
	ID                  uuid.UUID
	Currency            string  `json:"currency,omitempty" binding:"required" validate:"required"`
	Amount              float64 `json:"amount,omitempty" binding:"required" validate:"required"`
	Term                string  `json:"term,omitempty" binding:"required" validate:"required"`
	ARAging             int16   `json:"araging,omitempty" binding:"required" validate:"required"`
	Remarks             string  `json:"remarks,omitempty" binding:"required" validate:"required"`
	DueDate             string  `json:"dueDate,omitempty" binding:"required" validate:"required,ISO8601date"`
	SettlementDate      string  `json:"settlementDate,omitempty" binding:"required" validate:"required,ISO8601date"`
	OperatingActivityID string  `json:"operatingActivityId" binding:"required" validate:"required,uuid4"`
	UpdatedBy           uuid.UUID
}

type InsertPayment struct {
	Currency            string
	Amount              float64
	Term                string
	ARAging             int16
	Remarks             string
	DueDate             time.Time
	SettlementDate      time.Time
	OperatingActivityID uuid.UUID
	CreatedBy           uuid.UUID
}

type UpdatePayment struct {
	ID                  uuid.UUID
	Currency            string
	Amount              float64
	Term                string
	ARAging             int16
	Remarks             string
	DueDate             time.Time
	SettlementDate      time.Time
	OperatingActivityID uuid.UUID
	UpdatedBy           uuid.UUID
}
