package database

import (
	"messify/internal/domain/users"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&users.User{},
	)
}