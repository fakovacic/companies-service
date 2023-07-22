package middleware

import (
	"context"

	"github.com/fakovacic/companies-service/internal/companies"
)

func NewLoggingMiddleware(next companies.Service, config *companies.Config) companies.Service {
	m := loggingMiddleware{
		next:    next,
		config:  config,
		service: "companies",
	}

	return &m
}

type loggingMiddleware struct {
	next    companies.Service
	config  *companies.Config
	service string
}

func (m *loggingMiddleware) Get(ctx context.Context, input string) (*companies.Company, error) {
	reqID := companies.GetCtxStringVal(ctx, companies.ContextKeyRequestID)

	m.config.Log.Info().
		Str("reqID", reqID).
		Str("service", m.service).
		Str("method", "Get").
		Interface("input", input).
		Msg("service request")

	model, err := m.next.Get(ctx, input)

	m.config.Log.Info().
		Str("reqID", reqID).
		Str("service", m.service).
		Str("method", "Get").
		Interface("model", model).
		Err(err).
		Msg("service response")

	return model, err
}

func (m *loggingMiddleware) Create(ctx context.Context, input *companies.Company) (*companies.Company, error) {
	reqID := companies.GetCtxStringVal(ctx, companies.ContextKeyRequestID)

	m.config.Log.Info().
		Str("reqID", reqID).
		Str("service", m.service).
		Str("method", "Create").
		Interface("input", input).
		Msg("service request")

	model, err := m.next.Create(ctx, input)

	m.config.Log.Info().
		Str("reqID", reqID).
		Str("service", m.service).
		Str("method", "Create").
		Interface("model", model).
		Err(err).
		Msg("service response")

	return model, err
}

func (m *loggingMiddleware) Update(ctx context.Context, id string, input *companies.Company, fields []string) (*companies.Company, error) {
	reqID := companies.GetCtxStringVal(ctx, companies.ContextKeyRequestID)

	m.config.Log.Info().
		Str("reqID", reqID).
		Str("service", m.service).
		Str("method", "Update").
		Str("id", id).
		Interface("input", input).
		Interface("fields", fields).
		Msg("service request")

	model, err := m.next.Update(ctx, id, input, fields)

	m.config.Log.Info().
		Str("reqID", reqID).
		Str("service", m.service).
		Str("method", "Update").
		Interface("model", model).
		Err(err).
		Msg("service response")

	return model, err
}

func (m *loggingMiddleware) Delete(ctx context.Context, id string) error {
	reqID := companies.GetCtxStringVal(ctx, companies.ContextKeyRequestID)

	m.config.Log.Info().
		Str("reqID", reqID).
		Str("service", m.service).
		Str("method", "Delete").
		Str("id", id).
		Msg("service request")

	err := m.next.Delete(ctx, id)

	m.config.Log.Info().
		Str("reqID", reqID).
		Str("service", m.service).
		Str("method", "Delete").
		Err(err).
		Msg("service response")

	return err
}
