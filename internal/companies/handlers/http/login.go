package http

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

const (
	jwtDuration = time.Hour * 72
)

type LoginResponse struct {
	Token string `json:"token"`
}

func (h *Handler) Login() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		headers := c.GetReqHeaders()

		if headers["X-Auth"] == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		// send request to auth service

		claims := jwt.MapClaims{
			"name":  "Admin",
			"admin": true,
			"exp":   time.Now().Add(jwtDuration).Unix(),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		t, err := token.SignedString([]byte(
			h.config.JWTSigningKey,
		))
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(&LoginResponse{
			Token: t,
		})
	}
}
