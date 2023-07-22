package postgres

import (
	"context"
	goErrors "errors"

	"github.com/fakovacic/companies-service/internal/companies"
	"github.com/fakovacic/companies-service/internal/companies/errors"
	"github.com/lib/pq"
)

func (s *store) Create(ctx context.Context, model *companies.Company) error {
	_, err := s.db.ExecContext(ctx, `
		INSERT INTO companies (id, name, description, employees_amount, registered, type, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		model.ID,
		model.Name,
		model.Description,
		model.EmployeesAmount,
		model.Registered,
		model.Type,
		model.CreatedAt,
		model.UpdatedAt,
	)
	if err != nil {
		var pErr *pq.Error

		ok := goErrors.As(err, &pErr)
		if ok {
			err = errors.Wrap(err, " database error: %s", pErr.Code.Class().Name())
		}

		return errors.Wrap(err, "create company")
	}

	return nil
}
