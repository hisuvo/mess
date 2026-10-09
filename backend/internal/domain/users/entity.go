package users

import "gorm.io/gorm"

type UserStatus string
var (
	UserStatusActive UserStatus = "ACTIVE"
	UserStatusInactive UserStatus = "INACTIVE"
	StatusBlocked  UserStatus = "BLOCKED"
)
 
type UserRole string
var (
	UserRoleAdmin UserRole = "ADMIN"
	UserRoleManager UserRole = "MANAGER"
	UserRoleMember UserRole = "MEMBER"
)

type User struct {
    gorm.Model // Provides ID, CreatedAt, UpdatedAt, DeletedAt

    Name     string `gorm:"type:varchar(50);not null"`
    Email    string `gorm:"type:varchar(50);uniqueIndex;not null"`
    Phone    string `gorm:"type:varchar(20);not null"`
    Password string `gorm:"not null"`
    ImageURL string `gorm:"type:text"`
    Address  string `gorm:"type:text"`
    Role     UserRole `gorm:"type:varchar(30);not null;default:'MEMBER';index"`
    Status   UserStatus `gorm:"type:varchar(30);not null;default:'ACTIVE';index"`
}