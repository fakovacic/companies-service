package notifier

import (
	"github.com/fakovacic/companies-service/internal/companies"
	"github.com/segmentio/kafka-go"
)

func New(conn *kafka.Conn) companies.Notifier {
	return &notifier{
		conn: conn,
	}
}

type notifier struct {
	conn *kafka.Conn
}
