package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/rizkyizh/go-fiber-boilerplate/config"
	"github.com/rizkyizh/go-fiber-boilerplate/database"
	"github.com/rizkyizh/go-fiber-boilerplate/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	config.LoadConfig()
	database.ConnectDB()

	app := fiber.New()

	routes.SetupRoutesApp(app)

	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
