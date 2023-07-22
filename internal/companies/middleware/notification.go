package middleware

import (
	"context"

	"github.com/fakovacic/companies-service/internal/companies"
)

type notificationMiddleware struct {
	next     companies.Service
	notifier companies.Notifier
}

func NewNotificationMiddleware(next companies.Service, notifier companies.Notifier) companies.Service {
	m := notificationMiddleware{
		next:     next,
		notifier: notifier,
	}

	return &m
}

func (m *notificationMiddleware) Get(ctx context.Context, input string) (*companies.Company, error) {
	return m.next.Get(ctx, input)
}

func (m *notificationMiddleware) Create(ctx context.Context, input *companies.Company) (*companies.Company, error) {
	return m.next.Create(ctx, input)
}

func (m *notificationMiddleware) Update(ctx context.Context, id string, input *companies.Company, fields []string) (*companies.Company, error) {
	model, err := m.next.Update(ctx, id, input, fields)

	if err == nil {
		m.notifier.Send(ctx)
	}

	return model, err
}

func (m *notificationMiddleware) Delete(ctx context.Context, id string) error {
	return m.next.Delete(ctx, id)
}
