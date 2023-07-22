package notifier

import "context"

func (s *notifier) Send(ctx context.Context) {
	s.config.Log.Info().Msg("send notification")
}
