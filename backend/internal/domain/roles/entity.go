package roles

import (
	"messify/internal/domain/permssions"

	"gorm.io/gorm"
)

type Role struct {
	gorm.Model

	TenantID    *uint                    `gorm:"index" json:"tenant_id,omitempty"`
	Name        string                   `gorm:"type:varchar(50);not null;index" json:"name"`
	Description string                   `gorm:"type:text" json:"description,omitempty"`
	Permissions []permssions.Permission `gorm:"many2many:role_permissions;" json:"permissions,omitempty"`
}
