package middlewares

import (
	"github.com/gofiber/fiber/v2"

	"github.com/rizkyizh/go-fiber-boilerplate/utils"
)

// RequireRole restricts access to users whose role is in the provided list.
// Must be used after AuthRequired.
func RequireRole(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		h := &utils.ResponseHandler{}

		claims, ok := c.Locals("claims").(*utils.TokenClaims)
		if !ok || claims == nil {
			return h.Unauthorized(c, []string{"unauthenticated"})
		}

		for _, role := range roles {
			if claims.Role == role {
				return c.Next()
			}
		}

		return h.Forbidden(c, []string{"insufficient permissions"})
	}
}
