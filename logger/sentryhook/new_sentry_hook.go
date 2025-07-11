package sentryhook

import (
	"time"

	"github.com/getsentry/sentry-go"
	"go.uber.org/zap/zapcore"
)

type SentryCore struct {
	zapcore.Core
	MinLevel zapcore.Level
}

func (s *SentryCore) With(fields []zapcore.Field) zapcore.Core {
	return &SentryCore{
		Core:     s.Core.With(fields),
		MinLevel: s.MinLevel,
	}
}

func (s *SentryCore) Check(entry zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if entry.Level >= s.MinLevel {
		return ce.AddCore(entry, s)
	}
	return ce
}

func (s *SentryCore) Write(entry zapcore.Entry, fields []zapcore.Field) error {
	if entry.Level >= s.MinLevel {
		logToSentry(entry, fields)
	}
	return s.Core.Write(entry, fields)
}

func logToSentry(entry zapcore.Entry, fields []zapcore.Field) {
	event := &sentry.Event{
		Message: entry.Message,
		Level:   ConvertLevel(entry.Level),
		Extra:   make(map[string]any),
	}

	event.Extra["caller"] = entry.Caller.String()
	event.Extra["timestamp"] = entry.Time.Format(time.RFC3339)

	for _, f := range fields {
		if f.Key != "" {
			event.Extra[f.Key] = f.Interface
		}
	}

	sentry.CaptureEvent(event)
	sentry.Flush(2 * time.Second)
}
