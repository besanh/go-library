package sentryhook

import (
	"encoding/base64"
	"runtime"
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
		Extra: map[string]any{
			"caller":    entry.Caller.String(),
			"file":      entry.Caller.TrimmedPath(),
			"logger":    entry.LoggerName,
			"timestamp": entry.Time.Format(time.RFC3339),
		},
	}

	if fn := runtime.FuncForPC(entry.Caller.PC); fn != nil {
		event.Extra["func"] = fn.Name()
	}

	for _, f := range fields {
		if f.Key == "" {
			continue
		}
		switch f.Type {
		case zapcore.SkipType, zapcore.NamespaceType:
			continue
		case zapcore.StringType:
			event.Extra[f.Key] = f.String
		case zapcore.ErrorType:
			event.Extra[f.Key] = f.Interface
		case zapcore.Int64Type, zapcore.Int32Type, zapcore.Int16Type, zapcore.Int8Type,
			zapcore.Uint64Type, zapcore.Uint32Type, zapcore.Uint16Type, zapcore.Uint8Type:
			event.Extra[f.Key] = f.Integer
		case zapcore.BoolType:
			event.Extra[f.Key] = f.Integer == 1
		case zapcore.Float64Type, zapcore.Float32Type:
			if f.Interface != nil {
				event.Extra[f.Key] = f.Interface
			}
		case zapcore.BinaryType:
			event.Extra[f.Key] = base64.StdEncoding.EncodeToString(f.Interface.([]byte))
		case zapcore.TimeType:
			if t, ok := f.Interface.(time.Time); ok {
				event.Extra[f.Key] = t.Format(time.RFC3339)
			}
		case zapcore.DurationType:
			if d, ok := f.Interface.(time.Duration); ok {
				event.Extra[f.Key] = d.String()
			}
		default:
			if f.Interface != nil {
				event.Extra[f.Key] = f.Interface
			}
		}
	}

	sentry.CaptureEvent(event)
	sentry.Flush(2 * time.Second)
}
