package dto

import "time"

type RequestCreatePayment struct {
	TenantID      uint      `json:"tenant_id" validate:"required"`
	UserID        *uint     `json:"user_id" validate:"omitempty"`
	Amount        float64   `json:"amount" validate:"required"`
	PaymentType   string    `json:"payment_type" validate:"required"`
	PaymentMethod string    `json:"payment_method" validate:"required"`
	PaymentDate   time.Time `json:"payment_date" validate:"required"`
	TransactionID string    `json:"transaction_id" validate:"omitempty"`
	Status        string    `json:"status" validate:"omitempty"`
	Description   string    `json:"description" validate:"omitempty"`
}

type RequestUpdatePayment struct {
	ID            uint      `json:"id" validate:"required"`
	Amount        float64   `json:"amount" validate:"omitempty"`
	PaymentType   string    `json:"payment_type" validate:"omitempty"`
	PaymentMethod string    `json:"payment_method" validate:"omitempty"`
	PaymentDate   time.Time `json:"payment_date" validate:"omitempty"`
	TransactionID string    `json:"transaction_id" validate:"omitempty"`
	Status        string    `json:"status" validate:"omitempty"`
	Description   string    `json:"description" validate:"omitempty"`
}
