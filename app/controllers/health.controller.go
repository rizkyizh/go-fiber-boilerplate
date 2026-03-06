package controllers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/rizkyizh/go-fiber-boilerplate/database"
	"github.com/rizkyizh/go-fiber-boilerplate/utils"
)

type HealthController struct{}

func NewHealthController() *HealthController {
	return &HealthController{}
}

// HealthCheck godoc
// @Summary Health check
// @Description Returns application and database health status
// @Tags Health
// @Produce json
// @Success 200 {object} utils.ResponseData
// @Failure 503 {object} utils.ErrorResponse
// @Router /health [get]
func (ctrl *HealthController) HealthCheck(c *fiber.Ctx) error {
	h := &utils.ResponseHandler{}

	sqlDB, err := database.DB.DB()
	if err != nil {
		return h.ServiceUnavailable(c, []string{"database connection error"})
	}

	if err := sqlDB.Ping(); err != nil {
		return h.ServiceUnavailable(c, []string{"database ping failed: " + err.Error()})
	}

	return h.Ok(c, fiber.Map{
		"status":   "ok",
		"database": "connected",
	}, "healthy", nil)
}
