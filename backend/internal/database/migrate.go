package database

import (
	"messify/internal/domain/bazars"
	"messify/internal/domain/deposits"
	"messify/internal/domain/meals"
	"messify/internal/domain/notices"
	"messify/internal/domain/payments"
	"messify/internal/domain/permssions"
	"messify/internal/domain/roles"
	"messify/internal/domain/subjects"
	"messify/internal/domain/tenants"
	"messify/internal/domain/users"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&tenants.Tenant{},
		&permssions.Permission{},
		&roles.Role{},
		&users.User{},
		&bazars.Bazar{},
		&meals.Meal{},
		&deposits.Deposit{},
		&payments.Payment{},
		&notices.Notice{},
		&subjects.Subject{},
	)
}