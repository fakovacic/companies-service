package postgres

import (
	"context"
	"database/sql"
	goErrors "errors"

	"github.com/fakovacic/companies-service/internal/companies"
	"github.com/fakovacic/companies-service/internal/companies/errors"
	"github.com/lib/pq"
)

func (s *store) Get(ctx context.Context, id string) (*companies.Company, error) {
	var usr companies.Company

	err := s.db.QueryRowContext(ctx,
		`SELECT 
			id, 
			name,
			description,
			employees_amount,
			registered,
			type,
			created_at,
			updated_at
		 FROM companies 
		 WHERE id=$1`,
		id,
	).Scan(
		&usr.ID,
		&usr.Name,
		&usr.Description,
		&usr.EmployeesAmount,
		&usr.Registered,
		&usr.Type,
		&usr.CreatedAt,
		&usr.UpdatedAt,
	)
	if err != nil {
		if goErrors.Is(err, sql.ErrNoRows) {
			return nil, errors.NotFound("not found company")
		}

		var pErr *pq.Error

		ok := goErrors.As(err, &pErr)
		if ok {
			err = errors.Wrap(err, " database error: %s", pErr.Code.Class().Name())
		}

		return nil, errors.Wrap(err, "get company")
	}

	return &usr, nil
}
