package database

import (
	"messify/internal/config"
	"messify/internal/domain/users"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB, cnfg *config.Config) error {
	var count int64

	db.Model(&users.User{}).Count(&count)

	if count > 0 {
		return nil
	}

	owner := users.User{
		Name: cnfg.OwnerName,
		Email: cnfg.OwnerEmail,
		Password: cnfg.OwnerPassword,
		Phone: cnfg.OwnerPhone,
	}

	return db.Model(&users.User{}).Create(&owner).Error
}