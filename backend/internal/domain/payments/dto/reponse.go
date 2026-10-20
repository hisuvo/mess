package dto

import (
	"messify/internal/domain/users"
	"time"
)

type ResponsePayment struct {
	ID            uint      `json:"id"`
	TenantID      uint      `json:"tenant_id"`
	UserID        *uint     `json:"user_id,omitempty"`
	User          *users.User `json:"user,omitempty"`
	Amount        float64   `json:"amount"`
	Type          string    `json:"type"`
	Method        string    `json:"method"`
	PaymentDate   time.Time `json:"payment_date"`
	TransactionID string    `json:"transaction_id,omitempty"`
	Status        string    `json:"status"`
	Description   string    `json:"description,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
