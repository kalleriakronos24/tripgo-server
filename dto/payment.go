package dto

import (
	"github.com/google/uuid"
	"time"
)

type InsertPayment struct {
	Currency       string    `json:"currency,omitempty" binding:"required"`
	Amount         float64   `json:"amount,omitempty" binding:"required"`
	Term           string    `json:"term,omitempty" binding:"required"`
	ARAging        int16     `json:"araging,omitempty" binding:"required"`
	Remarks        string    `json:"remarks,omitempty" binding:"required"`
	DueDate        time.Time `json:"dueDate,omitempty" binding:"required"`
	SettlementDate time.Time `json:"settlementDate,omitempty" binding:"required"`

	OperatingActivityID uuid.UUID `json:"operatingActivityId" binding:"required"`
	CreatedBy           uuid.UUID
}

type UpdatePayment struct {
	ID             uuid.UUID
	Currency       string    `json:"currency,omitempty" binding:"required"`
	Amount         float64   `json:"amount,omitempty" binding:"required"`
	Term           string    `json:"term,omitempty" binding:"required"`
	ARAging        int16     `json:"araging,omitempty" binding:"required"`
	Remarks        string    `json:"remarks,omitempty" binding:"required"`
	DueDate        time.Time `json:"dueDate,omitempty" binding:"required"`
	SettlementDate time.Time `json:"settlementDate,omitempty" binding:"required"`

	OperatingActivityID uuid.UUID `json:"operatingActivityId" binding:"required"`
	UpdatedBy           uuid.UUID
}
