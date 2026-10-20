package dto

type RequestCreateRole struct {
	TenantID uint   `json:"tenant_id" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

type RequestUpdateRole struct {
	ID       uint   `json:"id" validate:"required"`
	TenantID uint   `json:"tenant_id" validate:"required"`
	Name     string `json:"name" validate:"omitempty"`
}
