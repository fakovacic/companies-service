package companies

import (
	"context"

	"github.com/fakovacic/companies-service/internal/companies/errors"
	"github.com/google/uuid"
)

func (s *service) Delete(ctx context.Context, id string) error {
	val, err := uuid.Parse(id)
	if err != nil {
		return errors.BadRequest("id invalid format")
	}

	if val == uuid.Nil {
		return errors.BadRequest("id invalid format")
	}

	err = s.store.Delete(ctx, id)
	if err != nil {
		return errors.Wrap(err, "delete company")
	}

	return nil
}
