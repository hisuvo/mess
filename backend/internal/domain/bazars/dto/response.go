package dto

import (
	"messify/internal/domain/users"
	"time"
)

type ResponseBazar struct {
	ID       uint      `json:"id"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
	DeleteAt time.Time `json:"delete_at"`

	Name     string    `json:"name"`
	Date     time.Time `json:"date"`
	Total    float64   `json:"total"`
	IsActive bool      `json:"is_active"`

	TenantID uint `json:"tenant_id"`

	UserID uint `json:"user_id"`
	User users.User `json:"user"`
}