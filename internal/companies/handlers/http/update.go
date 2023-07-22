package http

import (
	goErrors "errors"
	"net/http"

	"github.com/fakovacic/companies-service/internal/companies"
	"github.com/fakovacic/companies-service/internal/companies/errors"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

type UpdateRequest struct {
	Fields []string      `json:"fields"`
	Data   UpdateCompany `json:"data"`
}

type UpdateCompany struct {
	Name            string                `json:"name,omitempty"`
	Description     string                `json:"description,omitempty"`
	EmployeesAmount uint32                `json:"employeesAmount,omitempty"`
	Registered      bool                  `json:"registered,omitempty"`
	Type            companies.CompanyType `json:"type,omitempty"`
}

func (h *Handler) Update() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		if id == "" {
			c.Status(http.StatusInternalServerError)

			return c.JSON(ErrorMessage{
				Message: "id is required",
			})
		}

		var (
			json = jsoniter.ConfigCompatibleWithStandardLibrary
			req  UpdateRequest
		)

		err := json.Unmarshal(c.Body(), &req)
		if err != nil {
			c.Status(http.StatusInternalServerError)

			return c.JSON(ErrorMessage{
				Message: err.Error(),
			})
		}

		res, err := h.service.Update(c.Context(), id, &companies.Company{
			Name:            req.Data.Name,
			Description:     req.Data.Description,
			EmployeesAmount: req.Data.EmployeesAmount,
			Registered:      req.Data.Registered,
			Type:            req.Data.Type,
		},
			req.Fields,
		)
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

		c.Status(http.StatusOK)

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
