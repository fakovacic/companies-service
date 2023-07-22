package config

import (
	"os"

	"github.com/fakovacic/companies-service/internal/companies"
	"github.com/fakovacic/companies-service/internal/companies/errors"
)

func NewConfig() (*companies.Config, error) {
	env := os.Getenv("ENV")
	if env == "" {
		return nil, errors.New("ENV is empty")
	}

	return companies.NewConfig(env), nil
}
