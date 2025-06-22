package sentryhook

import (
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/rs/zerolog"
)

func (i *sentryHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	if level < zerolog.ErrorLevel {
		return
	}
	sentry.WithScope(func(scope *sentry.Scope) {
		scope.SetLevel(Convert(level))
		sentry.CaptureMessage(msg)
	})
	sentry.Flush(2 * time.Second)
}

func Convert(l zerolog.Level) sentry.Level {
	switch l {
	case zerolog.DebugLevel:
		return sentry.LevelDebug
	case zerolog.InfoLevel:
		return sentry.LevelInfo
	case zerolog.WarnLevel:
		return sentry.LevelWarning
	case zerolog.ErrorLevel:
		return sentry.LevelError
	case zerolog.FatalLevel:
		return sentry.LevelFatal
	default:
		return sentry.LevelError
	}
}
