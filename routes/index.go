package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/rizkyizh/go-fiber-boilerplate/app/routes"
)

func SetupRoutesApp(app *fiber.App) {
	routes.UserRoutes(app.Group("/users"))

	// Default route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World 🌍🚀")
	})

	// 404 Route
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("Not Found")
	})
}
