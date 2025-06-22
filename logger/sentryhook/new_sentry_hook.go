package sentryhook

import (
	"github.com/getsentry/sentry-go"
	"github.com/rs/zerolog"
)

type ISentryHook interface {
	Run(event *zerolog.Event, level sentry.Level, msg string)
}

type SentryOption func(*sentry.ClientOptions)

func NewSentryHook(dsn string, opts ...SentryOption) zerolog.Hook {
	cfg := sentry.ClientOptions{
		Dsn: dsn,
	}
	for _, fn := range opts {
		fn(&cfg)
	}
	_ = sentry.Init(cfg)

	return &sentryHook{}
}
