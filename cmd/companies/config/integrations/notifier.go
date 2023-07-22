package integrations

import (
	"github.com/fakovacic/companies-service/internal/companies"
	"github.com/fakovacic/companies-service/internal/companies/integrations/notifier"
)

func NewNotifier(cfg *companies.Config) companies.Notifier {
	return notifier.New(cfg)
}
