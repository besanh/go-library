package logger

import "github.com/rs/zerolog"

type loggerImpl struct{ z zerolog.Logger }

func (l *loggerImpl) With(fields ...Field) ILogger {
	ev := l.z.With()
	for _, f := range fields {
		ev = ev.Interface(f.Key, f.Value)
	}
	return &loggerImpl{ev.Logger()}
}

func (l *loggerImpl) Debug(msg string, fields ...Field) {
	ev := l.z.Debug()
	for _, f := range fields {
		ev = ev.Interface(f.Key, f.Value)
	}
	ev.Msg(msg)
}

func (l *loggerImpl) Info(msg string, fields ...Field) {
	ev := l.z.Info()
	for _, f := range fields {
		ev = ev.Interface(f.Key, f.Value)
	}
	ev.Msg(msg)
}

func (l *loggerImpl) Warn(msg string, fields ...Field) {
	ev := l.z.Warn()
	for _, f := range fields {
		ev = ev.Interface(f.Key, f.Value)
	}
	ev.Msg(msg)
}

func (l *loggerImpl) Error(msg string, fields ...Field) {
	ev := l.z.Error()
	for _, f := range fields {
		ev = ev.Interface(f.Key, f.Value)
	}
	ev.Msg(msg)
}

func (l *loggerImpl) Fatal(msg string, fields ...Field) {
	ev := l.z.Fatal()
	for _, f := range fields {
		ev = ev.Interface(f.Key, f.Value)
	}
	ev.Msg(msg)
}
