package postgres

import (
	"context"
	goErrors "errors"

	"github.com/fakovacic/companies-service/internal/companies/errors"
	"github.com/lib/pq"
)

func (s *store) Delete(ctx context.Context, id string) error {
	_, err := s.db.ExecContext(ctx, `DELETE FROM companies WHERE id = $1`, id)
	if err != nil {
		var pErr *pq.Error

		ok := goErrors.As(err, &pErr)
		if ok {
			err = errors.Wrap(err, " database error: %s", pErr.Code.Class().Name())
		}

		return errors.Wrap(err, "delete company")
	}

	return nil
}
