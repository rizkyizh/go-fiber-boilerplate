package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
)

// SetupHelmet adds security-related HTTP headers to every response.
func SetupHelmet(app *fiber.App) {
	app.Use(helmet.New())
}
