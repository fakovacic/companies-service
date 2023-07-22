package middleware

import (
	"github.com/fakovacic/companies-service/internal/companies"
	"github.com/gofiber/fiber/v2"
)

func Logger(config *companies.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		l := config.Log.With().
			Str("url", c.OriginalURL()).
			Str("method", c.Route().Method).
			Logger()
		l.Info().Msg("http request")

		return c.Next()
	}
}
