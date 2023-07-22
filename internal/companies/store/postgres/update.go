package postgres

import (
	"context"
	goErrors "errors"

	"github.com/fakovacic/companies-service/internal/companies"
	"github.com/fakovacic/companies-service/internal/companies/errors"
	"github.com/lib/pq"
)

func (s *store) Update(ctx context.Context, id string, model *companies.Company) error {
	_, err := s.db.ExecContext(ctx, `UPDATE companies 
		SET 
			name = $1, 
			description = $2, 
			employees_amount = $3, 
			registered = $4, 
			type = $5, 
			updated_at = $6
		WHERE id = $7`,
		model.Name,
		model.Description,
		model.EmployeesAmount,
		model.Registered,
		model.Type,
		model.UpdatedAt,
		id,
	)
	if err != nil {
		var pErr *pq.Error

		ok := goErrors.As(err, &pErr)
		if ok {
			err = errors.Wrap(err, " database error: %s", pErr.Code.Class().Name())
		}

		return errors.Wrap(err, "update company")
	}

	return nil
}
