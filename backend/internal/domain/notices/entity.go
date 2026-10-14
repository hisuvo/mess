package notices

import (
	"messify/internal/domain/users"

	"gorm.io/gorm"
)

type NoticeStatus string

const (
	NoticeStatusActive   NoticeStatus = "ACTIVE"
	NoticeStatusArchived NoticeStatus = "ARCHIVED"
)

type Notice struct {
	gorm.Model

	TenantID    uint         `gorm:"index;not null" json:"tenant_id"`
	Title       string       `gorm:"type:varchar(250);not null" json:"title"`
	Description string       `gorm:"type:text" json:"description"`
	UserID      uint         `gorm:"index;not null" json:"user_id"`
	User        *users.User  `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Status      NoticeStatus `gorm:"type:varchar(30);not null;default:'ACTIVE';index" json:"status"`
}