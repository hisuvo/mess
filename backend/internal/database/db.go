package database

import (
	"errors"
	"fmt"
	"messify/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(cnfg *config.Config) (*gorm.DB, error) {

	dsn := cnfg.DatabaseURL

	if dsn == "" {
		fmt.Println("Database connection string is empty")
		return nil, errors.New("Database connection string is empty")
	}
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})
	
	if err != nil {
		fmt.Println("Database connection error: ", err.Error())
	}

	Migrate(db)

	// Owner seed function call here
	if err := Seed(db, cnfg); err != nil {
		fmt.Println("Owner seed failed")
	}
	
	fmt.Println("Owner seed completed successfully")

	fmt.Println("Database connected successfully")
	return db, nil
}