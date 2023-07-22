package companies

import (
	"context"
	"os"

	"github.com/rs/zerolog"
)

const (
	ServiceName string = "companies"
)

func NewConfig(env string) *Config {
	return &Config{
		Env: env,
		Log: newLogger(),
	}
}

type Config struct {
	Log zerolog.Logger
	Env string
}

func newLogger() zerolog.Logger {
	zerolog.TimestampFieldName = "timestamp"
	zerolog.LevelFieldName = "level"
	zerolog.MessageFieldName = "msg"
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	return zerolog.New(os.Stdout).
		With().Timestamp().Logger()
}

// Context

type ContextKey string

const (
	ContextKeyRequestID ContextKey = "requestID"
)

func GetCtxStringVal(ctx context.Context, key ContextKey) string {
	ctxValue := ctx.Value(key)

	if ctxValue != nil {
		val, ok := ctxValue.(string)
		if ok {
			return val
		}
	}

	return ""
}
