package dto

import "time"

type ResponseTenant struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ResponseGetTenants struct {
	Tenants []ResponseTenant `json:"tenants"`
	Total   int64            `json:"total"`
}