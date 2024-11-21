package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	"github.com/rizkyizh/go-fiber-boilerplate/app/routes"
	_ "github.com/rizkyizh/go-fiber-boilerplate/docs"
)

func SetupRoutesApp(app *fiber.App) {
	routes.UserRoutes(app.Group("/users"))

	// Default route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World 🌍🚀")
	})

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	// 404 Route
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("Not Found")
	})
}
