package http

import (
	goErrors "errors"
	"net/http"

	"github.com/fakovacic/companies-service/internal/companies/errors"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Get() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		if id == "" {
			c.Status(http.StatusInternalServerError)

			return c.JSON(ErrorMessage{
				Message: "id is required",
			})
		}

		res, err := h.service.Get(c.Context(), id)
		if err != nil {
			var internalError *errors.Error

			if goErrors.As(err, &internalError) {
				c.Status(internalError.HTTPStatusCode())

				return c.JSON(ErrorMessage{
					Message: internalError.Error(),
				})
			}

			c.Status(http.StatusInternalServerError)

			return c.JSON(ErrorMessage{
				Message: err.Error(),
			})
		}

		return c.JSON(Company{
			ID:              res.ID,
			Name:            res.Name,
			Description:     res.Description,
			EmployeesAmount: res.EmployeesAmount,
			Registered:      res.Registered,
			Type:            res.Type,
			CreatedAt:       res.CreatedAt,
			UpdatedAt:       res.UpdatedAt,
		})
	}
}
