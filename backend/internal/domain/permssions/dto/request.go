package dto

type RequestCreatePermission struct {
	TenantID uint   `json:"tenant_id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	RoleID   uint   `json:"role_id" validate:"required"`
}

type RequestUpdatePermission struct {
	ID       uint   `json:"id" validate:"required"`
	TenantID uint   `json:"tenant_id" validate:"required"`
	Name     string `json:"name" validate:"omitempty"`
	RoleID   uint   `json:"role_id" validate:"omitempty"`
}