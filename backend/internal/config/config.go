package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	DatabaseURL string
	JwtSecret string
	JwtIssuer string
	JwtAudience string
	JwtExpiresIn int
	AdminName string
	AdminEmail string
	AdminPassword string
	ApiKey string
}

func GetStringEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		return fmt.Sprintf("environment variable %s is not set", key)
	}
	return value
}

func GetIntEnv(key string) (int, error) {
	value := os.Getenv(key)
	if value == "" {
		return 0, fmt.Errorf("environment variable %s is not set", key)
	}
	val, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("environment variable %s is not an integer", key)
	}
	return val, nil
}

func LoadEnv() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Here env file load")
	}

	return &Config{
		Port: GetStringEnv("PORT"),
		DatabaseURL: GetStringEnv("DATABASE_URL"),
		JwtSecret: GetStringEnv("JWT_SECRET"),
		JwtIssuer: GetStringEnv("JWT_ISSUER"),
		JwtAudience: GetStringEnv("JWT_AUDIENCE"),
		// JwtExpiresIn: jwtExpiresIn,
		AdminEmail: GetStringEnv("ADMIN_EMAIL"),
		AdminPassword: GetStringEnv("ADMIN_PASSWORD"),
		AdminName: GetStringEnv("ADMIN_NAME"),
		ApiKey: GetStringEnv("API_KEY"),
	}, nil
}