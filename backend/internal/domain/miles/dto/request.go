package dto

import "time"

type RequestCreateMile struct {
	TenantID  uint      `json:"tenant_id" validate:"required"`
	UserID    uint      `json:"user_id" validate:"required"`
	Date      time.Time `json:"date" validate:"required"`
	Breakfast float64   `json:"breakfast" validate:"required"`
	Lunch     float64   `json:"lunch" validate:"required"`
	Dinner    float64   `json:"dinner" validate:"required"`
	Note      string    `json:"note" validate:"omitempty"`
}

type RequestUpdateMile struct {
	ID        uint    `json:"id" validate:"required"`
	Breakfast float64 `json:"breakfast" validate:"omitempty"`
	Lunch     float64 `json:"lunch" validate:"omitempty"`
	Dinner    float64 `json:"dinner" validate:"omitempty"`
	Note      string    `json:"note" validate:"omitempty"`
}