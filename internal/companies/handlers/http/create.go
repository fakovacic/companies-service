package http

import (
	goErrors "errors"
	"net/http"

	"github.com/fakovacic/companies-service/internal/companies"
	"github.com/fakovacic/companies-service/internal/companies/errors"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

type CreateRequest struct {
	Name            string                `json:"name"`
	Description     string                `json:"description,omitempty"`
	EmployeesAmount uint32                `json:"employeesAmount"`
	Registered      bool                  `json:"registered"`
	Type            companies.CompanyType `json:"type"`
}

func (h *Handler) Create() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var (
			json = jsoniter.ConfigCompatibleWithStandardLibrary
			req  CreateRequest
		)

		err := json.Unmarshal(c.Body(), &req)
		if err != nil {
			c.Status(http.StatusInternalServerError)

			return c.JSON(ErrorMessage{
				Message: err.Error(),
			})
		}

		res, err := h.service.Create(c.Context(), &companies.Company{
			Name:            req.Name,
			Description:     req.Description,
			EmployeesAmount: req.EmployeesAmount,
			Registered:      req.Registered,
			Type:            req.Type,
		})
		if err != nil {
			var internalError errors.Error

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

		c.Status(http.StatusCreated)

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
