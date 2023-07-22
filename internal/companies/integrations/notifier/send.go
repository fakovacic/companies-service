package notifier

import "context"

func (s *notifier) Send(_ context.Context) {
	s.config.Log.Info().Msg("send notification")
}
