package dto

type RequestCreateTenant struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"omitempty"`
	Phone   string `json:"phone" validate:"omitempty"`
	Email   string `json:"email" validate:"omitempty"`
}

type RequestUpdateTenant struct {
	ID      uint   `json:"id" validate:"required"`
	Name    string `json:"name" validate:"omitempty"`
	Address string `json:"address" validate:"omitempty"`
	Phone   string `json:"phone" validate:"omitempty"`
	Email   string `json:"email" validate:"omitempty"`
	Status  string `json:"status" validate:"omitempty"`
}