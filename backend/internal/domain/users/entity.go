package users

import "gorm.io/gorm"

type UserStatus string

var (
	UserStatusActive   UserStatus = "ACTIVE"
	UserStatusInactive UserStatus = "INACTIVE"
	StatusBlocked      UserStatus = "BLOCKED"
)

type UserRole string

var (
	UserRoleAdmin   UserRole = "ADMIN"
	UserRoleManager UserRole = "MANAGER"
	UserRoleMember  UserRole = "MEMBER"
)

type User struct {
	gorm.Model // Provides ID, CreatedAt, UpdatedAt, DeletedAt

	TenantID *uint      `gorm:"index" json:"tenant_id,omitempty"`
	RoleID   *uint      `gorm:"index" json:"role_id,omitempty"`
	Name     string     `gorm:"type:varchar(50);not null" json:"name"`
	Email    string     `gorm:"type:varchar(50);uniqueIndex;not null" json:"email"`
	Phone    string     `gorm:"type:varchar(20);not null" json:"phone"`
	Password string     `gorm:"not null" json:"-"`
	ImageURL string     `gorm:"type:text" json:"image_url,omitempty"`
	Address  string     `gorm:"type:text" json:"address,omitempty"`
	Role     UserRole   `gorm:"type:varchar(30);not null;default:'MEMBER';index" json:"role"`
	Status   UserStatus `gorm:"type:varchar(30);not null;default:'ACTIVE';index" json:"status"`
}