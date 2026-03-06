package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"

	"github.com/rizkyizh/go-fiber-boilerplate/utils"
)

// SetupRateLimiter applies a global rate limit of 60 requests per minute per IP.
func SetupRateLimiter(app *fiber.App) {
	app.Use(limiter.New(limiter.Config{
		Max:        60,
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			h := &utils.ResponseHandler{}
			return h.TooManyRequests(c, []string{"rate limit exceeded, try again later"})
		},
	}))
}
