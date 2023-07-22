package notifier

import (
	"github.com/fakovacic/companies-service/internal/companies"
)

func New(c *companies.Config) companies.Notifier {
	return &notifier{
		config: c,
	}
}

type notifier struct {
	config *companies.Config
}
