package companies

import (
	"context"

	"github.com/fakovacic/companies-service/internal/companies/errors"
	"github.com/google/uuid"
)

func (s *service) Get(ctx context.Context, id string) (*Company, error) {
	val, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.BadRequest("id invalid format")
	}

	if val == uuid.Nil {
		return nil, errors.BadRequest("id invalid format")
	}

	model, err := s.store.Get(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "get company")
	}

	return model, nil
}
