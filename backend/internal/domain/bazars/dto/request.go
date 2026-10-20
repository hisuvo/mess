package dto

import "time"

type RequestCreateBazar struct {
	TenantID uint   `json:"tenant_id" validate:"required"`
	UserID   uint   `json:"user_id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Date     time.Time `json:"date" validate:"required"`
	Total    float64   `json:"total" validate:"required"`
	IsActive bool      `json:"is_active" validate:"required"`
}

type RequestUpdateBazar struct {
	ID       uint      `json:"id" validate:"required"`
	Name     string    `json:"name" validate:"omitempty"`
	Date     time.Time `json:"date" validate:"omitempty"`
	Total    float64   `json:"total" validate:"omitempty"`
	IsActive bool      `json:"is_active" validate:"omitempty"`
}