package dto

type RequestCreateSubject struct {
	TenantID    uint   `json:"tenant_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"omitempty"`
	Status      string `json:"status" validate:"omitempty"`
}

type RequestUpdateSubject struct {
	ID          uint   `json:"id" validate:"required"`
	TenantID    uint   `json:"tenant_id" validate:"required"`
	Name        string `json:"name" validate:"omitempty"`
	Description string `json:"description" validate:"omitempty"`
	Status      string `json:"status" validate:"omitempty"`
}