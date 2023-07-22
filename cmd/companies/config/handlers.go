package config

import (
	"github.com/fakovacic/companies-service/internal/companies"
	"github.com/fakovacic/companies-service/internal/companies/handlers/http"
)

func NewHandlers(c *companies.Config, service companies.Service) http.Handler {
	return *http.New(c, service)
}
