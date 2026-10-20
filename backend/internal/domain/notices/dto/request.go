package dto

type RequestCreateNotice struct {
	TenantID    uint   `json:"tenant_id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	UserID      uint   `json:"user_id" validate:"required"`
}

type RequestUpdateNotice struct {
	ID          uint   `json:"id" validate:"required"`
	Title       string `json:"title" validate:"omitempty"`
	Description string `json:"description" validate:"omitempty"`
	Status      string `json:"status" validate:"omitempty"`
}