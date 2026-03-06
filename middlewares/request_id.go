package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

// SetupRequestID injects a unique X-Request-ID header into every request/response.
func SetupRequestID(app *fiber.App) {
	app.Use(requestid.New())
}
