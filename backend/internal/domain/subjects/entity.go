package subjects

import "gorm.io/gorm"

type SubjectStatus string

const (
	SubjectStatusActive   SubjectStatus = "ACTIVE"
	SubjectStatusInactive SubjectStatus = "INACTIVE"
)

type Subject struct {
	gorm.Model

	TenantID    uint          `gorm:"index;not null" json:"tenant_id"`
	Name        string        `gorm:"type:varchar(150);not null" json:"name"`
	Description string        `gorm:"type:text" json:"description,omitempty"`
	Status      SubjectStatus `gorm:"type:varchar(30);not null;default:'ACTIVE';index" json:"status"`
}
