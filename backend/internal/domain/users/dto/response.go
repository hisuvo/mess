package dto

import (
	"time"
)

type UserResponse struct {
	ID        uint      `json:"id"`
	TenantID  *uint     `json:"tenant_id,omitempty"`
	RoleID    *uint     `json:"role_id,omitempty"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	ImageURL  string    `json:"image_url"`
	Address   string    `json:"address"`
	Role      string    `json:"role"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}