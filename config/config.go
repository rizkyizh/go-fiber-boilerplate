package config

import (
	"log"
	"os"
)

type Config struct {
	DB_URL string
}

var AppConfig Config

func LoadConfig() {
	AppConfig = Config{
		DB_URL: os.Getenv("DATABASE_URL"),
	}

	if AppConfig.DB_URL == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}
}
