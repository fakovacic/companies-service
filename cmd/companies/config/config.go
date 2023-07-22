package config

import (
	"os"

	"github.com/fakovacic/companies-service/internal/companies"
	"github.com/fakovacic/companies-service/internal/companies/errors"
)

func NewConfig() (*companies.Config, error) {
	env := os.Getenv("ENV")
	if env == "" {
		return nil, errors.New("env is empty")
	}

	signingKey := os.Getenv("JWT_SIGNING_KEY")
	if signingKey == "" {
		return nil, errors.New("jwt signing key is empty")
	}

	return companies.NewConfig(env, signingKey), nil
}
