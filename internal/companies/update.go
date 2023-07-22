package companies

import (
	"context"

	"github.com/fakovacic/companies-service/internal/companies/errors"
	"github.com/google/uuid"
)

func (s *service) Update(ctx context.Context, id string, m *Company, fields []string) (*Company, error) {
	val, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.BadRequest("id invalid format")
	}

	if val == uuid.Nil {
		return nil, errors.BadRequest("id invalid format")
	}

	if len(fields) == 0 {
		return nil, errors.BadRequest("update fields empty")
	}

	for i := range fields {
		val, ok := CompanyUpdateFields[fields[i]]
		if !ok {
			return nil, errors.BadRequest("field '%s' not exist", fields[i])
		}

		if !val {
			return nil, errors.BadRequest("field '%s' cannot be updated", fields[i])
		}
	}

	model, err := s.store.Get(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "get company")
	}

	for i := range fields {
		switch fields[i] {
		case CompanyField.Name:
			model.Name = m.Name
		case CompanyField.Description:
			model.Description = m.Description
		case CompanyField.EmployeesAmount:
			model.EmployeesAmount = m.EmployeesAmount
		case CompanyField.Registered:
			model.Registered = m.Registered
		case CompanyField.Type:
			model.Type = m.Type
		}
	}

	err = model.Validate()
	if err != nil {
		return nil, errors.Wrap(err, "validation")
	}

	model.UpdatedAt = s.timeFunc()

	err = s.store.Update(ctx, id, model)
	if err != nil {
		return nil, errors.Wrap(err, "update company")
	}

	return model, nil
}
