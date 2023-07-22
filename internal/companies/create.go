package companies

import (
	"context"

	"github.com/fakovacic/companies-service/internal/companies/errors"
)

func (s *service) Create(ctx context.Context, m *Company) (*Company, error) {
	err := m.Validate()
	if err != nil {
		return nil, errors.Wrap(err, "validation")
	}

	m.ID = s.uuidFunc().String()
	m.CreatedAt = s.timeFunc()
	m.UpdatedAt = s.timeFunc()

	err = s.store.Create(ctx, m)
	if err != nil {
		return nil, errors.Wrap(err, "create company")
	}

	return m, nil
}
