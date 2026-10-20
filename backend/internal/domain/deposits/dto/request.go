package dto

import (
	"time"
)

type RequestCreateDeposit struct {
	TenantID      uint          `json:"tenant_id" validate:"required"`
	UserID        uint          `json:"user_id" validate:"required"`
	Amount        float64       `json:"amount" validate:"required"`
	PaymentDate   time.Time     `json:"payment_date" validate:"required"`
	Method        string        `json:"method" validate:"required"`
	TransactionID string        `json:"transaction_id" validate:"required"`
	Note          string        `json:"note" validate:"omitempty"`
}

type RequestUpdateDeposit struct {
	ID            uint   `json:"id" validate:"required"`
	Amount        float64 `json:"amount" validate:"omitempty"`
	PaymentDate   time.Time `json:"payment_date" validate:"omitempty"`
	Method        string `json:"method" validate:"omitempty"`
	TransactionID string `json:"transaction_id" validate:"omitempty"`
	Note          string `json:"note" validate:"omitempty"`
}