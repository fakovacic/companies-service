package companies

import "context"

//go:generate moq -out ./mocks/notifier.go -pkg mocks  . Notifier
type Notifier interface {
	Send(context.Context, string, string) error
}
