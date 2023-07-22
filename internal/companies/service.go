package companies

import (
	"context"
	"time"

	"github.com/google/uuid"
)

//go:generate moq -out ./mocks/service.go -pkg mocks  . Service
type Service interface {
	Get(context.Context, string) (*Company, error)
	Create(context.Context, *Company) (*Company, error)
	Update(context.Context, string, *Company, []string) (*Company, error)
	Delete(ctx context.Context, id string) error
}

//go:generate moq -out ./mocks/store.go -pkg mocks  . Store
type Store interface {
	Get(context.Context, string) (*Company, error)
	Create(context.Context, *Company) error
	Update(context.Context, string, *Company) error
	Delete(context.Context, string) error
}

func New(c *Config, store Store, timeFunc func() time.Time, uuidFunc func() uuid.UUID) Service {
	return &service{
		config:   c,
		store:    store,
		timeFunc: timeFunc,
		uuidFunc: uuidFunc,
	}
}

type service struct {
	config   *Config
	store    Store
	timeFunc func() time.Time
	uuidFunc func() uuid.UUID
}
