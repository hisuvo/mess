package dto

import "time"

type ResponseRole struct {
	ID        uint      `json:"id"`
	TenantID  uint      `json:"tenant_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ResponseGetRoles struct {
	Roles []ResponseRole `json:"roles"`
	Total int64          `json:"total"`
}