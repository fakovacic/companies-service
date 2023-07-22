package notifier

import (
	"context"

	"github.com/fakovacic/companies-service/internal/companies/errors"
	jsoniter "github.com/json-iterator/go"
	"github.com/segmentio/kafka-go"
)

type Event struct {
	Service string `json:"service"`
	Action  string `json:"action"`
}

func (n *notifier) Send(_ context.Context, service, action string) error {
	json := jsoniter.ConfigCompatibleWithStandardLibrary

	event := Event{
		Service: service,
		Action:  action,
	}

	bytes, err := json.Marshal(event)
	if err != nil {
		return errors.Wrap(err, "failed to marshal messages")
	}

	_, err = n.conn.WriteMessages(
		kafka.Message{
			Value: bytes,
		},
	)
	if err != nil {
		return errors.Wrap(err, "failed to write messages")
	}

	return nil
}
