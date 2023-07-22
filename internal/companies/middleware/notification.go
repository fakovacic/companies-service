package middleware

import (
	"context"

	"github.com/fakovacic/companies-service/internal/companies"
)

func NewNotificationMiddleware(next companies.Service, config *companies.Config, notifier companies.Notifier) companies.Service {
	m := notificationMiddleware{
		next:     next,
		config:   config,
		notifier: notifier,
	}

	return &m
}

type notificationMiddleware struct {
	next     companies.Service
	config   *companies.Config
	notifier companies.Notifier
}

func (m *notificationMiddleware) Get(ctx context.Context, input string) (*companies.Company, error) {
	return m.next.Get(ctx, input)
}

func (m *notificationMiddleware) Create(ctx context.Context, input *companies.Company) (*companies.Company, error) {
	return m.next.Create(ctx, input)
}

func (m *notificationMiddleware) Update(ctx context.Context, id string, input *companies.Company, fields []string) (*companies.Company, error) {
	reqID := companies.GetCtxStringVal(ctx, companies.ContextKeyRequestID)

	model, err := m.next.Update(ctx, id, input, fields)

	if err == nil {
		eventErr := m.notifier.Send(ctx, "companies", "updated")
		if eventErr != nil {
			m.config.Log.Error().Str("reqID", reqID).Err(eventErr).Msg("failed to send event")
		}

		if eventErr == nil {
			m.config.Log.Info().Str("reqID", reqID).Msg("event sent")
		}
	}

	return model, err
}

func (m *notificationMiddleware) Delete(ctx context.Context, id string) error {
	return m.next.Delete(ctx, id)
}
