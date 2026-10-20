package dto

import (
	"messify/internal/domain/roles"
	"time"
)

type ResponsePermission struct {
	ID        uint         `json:"id"`
	TenantID  uint         `json:"tenant_id"`
	Name      string       `json:"name"`
	RoleID    uint         `json:"role_id"`
	Role      *roles.Role  `json:"role,omitempty"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}