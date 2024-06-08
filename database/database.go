package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/rizkyizh/go-fiber-boilerplate/app/models"
	"github.com/rizkyizh/go-fiber-boilerplate/config"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dsn := config.AppConfig.DB_URL
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	log.Println("Database Connected successfully")

	// TODO: always drop table just for development
	err = DB.Migrator().DropTable(&models.User{})
	if err != nil {
		panic("failed to drop table")
	}

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Error AutoMigrate database: %v", err)
	}
}
