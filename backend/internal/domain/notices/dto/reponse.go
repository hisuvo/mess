package dto

import "time"

type ResponseNotice struct {
	ID        uint      `json:"id"`
	TenantID  uint      `json:"tenant_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserID      uint      `json:"user_id"`
	Status      string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}