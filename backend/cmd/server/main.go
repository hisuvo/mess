package main

import (
	"fmt"
	"messify/internal/config"
	"messify/internal/database"
	"messify/internal/servers"
)

func main() {
	config, err := config.LoadEnv()
	if err != nil {
		fmt.Println("Env loading error: ", err.Error())
	}
	
	db, err := database.ConnectDB(config.DatabaseURL)
	if err != nil {
		fmt.Println("Database connection error: ", err.Error())
	}

	servers.Srart(db, config)
}