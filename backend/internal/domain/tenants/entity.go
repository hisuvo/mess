package tenants

import "gorm.io/gorm"

type TenantStatus string

const (
	TenantStatusActive    TenantStatus = "ACTIVE"
	TenantStatusInactive  TenantStatus = "INACTIVE"
	TenantStatusSuspended TenantStatus = "SUSPENDED"
)

type Tenant struct {
	gorm.Model

	Name    string       `gorm:"type:varchar(100);not null" json:"name"`
	Address string       `gorm:"type:text" json:"address,omitempty"`
	Phone   string       `gorm:"type:varchar(20)" json:"phone,omitempty"`
	Email   string       `gorm:"type:varchar(100);uniqueIndex" json:"email,omitempty"`
	Status  TenantStatus `gorm:"type:varchar(30);not null;default:'ACTIVE';index" json:"status"`
}