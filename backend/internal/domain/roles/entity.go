package roles

import (
	"messify/internal/domain/permssions"

	"gorm.io/gorm"
)

var RoleStatus string

type Role struct {
	gorm.Model

	TenantID    *uint                    `gorm:"index" json:"tenant_id,omitempty"` // here tenant_id is mass_id
	Name        string                   `gorm:"type:varchar(50);not null;index" json:"name"`
	Description string                   `gorm:"type:text" json:"description,omitempty"`
	Status		string					 `gorm:"size:20;default:ACTIVE" json:"status"`
	Permissions []permssions.Permission `gorm:"many2many:role_permissions;" json:"permissions,omitempty"`
}

