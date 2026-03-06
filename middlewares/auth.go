package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"

	"github.com/rizkyizh/go-fiber-boilerplate/config"
	"github.com/rizkyizh/go-fiber-boilerplate/utils"
)

// AuthRequired validates the Bearer JWT access token and stores claims in c.Locals("claims").
func AuthRequired() fiber.Handler {
	return func(c *fiber.Ctx) error {
		h := &utils.ResponseHandler{}

		authHeader := c.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return h.Unauthorized(c, []string{"missing or invalid authorization header"})
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := utils.ValidateToken(tokenStr, config.AppConfig.JWT_SECRET)
		if err != nil {
			return h.Unauthorized(c, []string{"invalid or expired token"})
		}

		c.Locals("claims", claims)
		return c.Next()
	}
}
