package config

import (
	"log"
	"os"
)

type Config struct {
	DB_URL              string
	JWT_SECRET          string
	JWT_REFRESH_SECRET  string
	JWT_ACCESS_EXPIRY   string
	JWT_REFRESH_EXPIRY  string
}

var AppConfig Config

func LoadConfig() {
	AppConfig = Config{
		DB_URL:             os.Getenv("DATABASE_URL"),
		JWT_SECRET:         getEnvOrDefault("JWT_SECRET", "change-me-in-production-access"),
		JWT_REFRESH_SECRET: getEnvOrDefault("JWT_REFRESH_SECRET", "change-me-in-production-refresh"),
		JWT_ACCESS_EXPIRY:  getEnvOrDefault("JWT_ACCESS_EXPIRY", "15m"),
		JWT_REFRESH_EXPIRY: getEnvOrDefault("JWT_REFRESH_EXPIRY", "168h"),
	}

	if AppConfig.DB_URL == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}
}

func getEnvOrDefault(key, defaultVal string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultVal
}
