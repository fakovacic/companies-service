package http

import (
	"time"

	"github.com/fakovacic/companies-service/internal/companies"
)

func New(c *companies.Config, service companies.Service) *Handler {
	return &Handler{
		config:  c,
		service: service,
	}
}

type Handler struct {
	config  *companies.Config
	service companies.Service
}

type ErrorMessage struct {
	Message string `json:"message"`
}

type Company struct {
	ID              string                `json:"id"`
	Name            string                `json:"name"`
	Description     string                `json:"description"`
	EmployeesAmount uint32                `json:"employeesAmount"`
	Registered      bool                  `json:"registered"`
	Type            companies.CompanyType `json:"type"`
	CreatedAt       time.Time             `json:"createdAt"`
	UpdatedAt       time.Time             `json:"updatedAt"`
}
