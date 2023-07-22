package middleware

import (
	"github.com/fakovacic/companies-service/internal/companies"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func ReqID(config *companies.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals(companies.ContextKeyRequestID, uuid.New().String())

		return c.Next()
	}
}
