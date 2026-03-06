package middlewares

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	if os.Getenv("APP_ENV") == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})
	}
}

// SetupLogger adds structured JSON request logging via zerolog.
func SetupLogger(app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		log.Info().
			Str("requestId", c.GetRespHeader("X-Request-Id")).
			Str("method", c.Method()).
			Str("path", c.Path()).
			Int("status", c.Response().StatusCode()).
			Dur("latency", time.Since(start)).
			Str("ip", c.IP()).
			Msg("request")

		return err
	})
}
