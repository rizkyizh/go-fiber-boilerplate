package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	appRoutes "github.com/rizkyizh/go-fiber-boilerplate/app/routes"
	"github.com/rizkyizh/go-fiber-boilerplate/app/controllers"
	_ "github.com/rizkyizh/go-fiber-boilerplate/docs"
)

func SetupRoutesApp(app *fiber.App) {
	appRoutes.AuthRoutes(app.Group("/auth"))
	appRoutes.UserRoutes(app.Group("/users"))

	healthController := controllers.NewHealthController()
	app.Get("/health", healthController.HealthCheck)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World 🌍🚀")
	})

	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("Not Found")
	})
}
