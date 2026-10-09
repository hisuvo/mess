package database

import (
	"errors"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(dsn string) (*gorm.DB, error) {
	if dsn == "" {
		fmt.Println("Database connection string is empty")
		return nil, errors.New("Database connection string is empty")
	}
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Database connection error: ", err.Error())
	}

	Migrate(db)
	fmt.Println("Database connected successfully")
	return db, nil
}