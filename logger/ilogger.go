package logger

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

type ILogger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
	Fatal(msg string, args ...any)

	With(args ...any) ILogger
}

type loggerImpl struct {
	z zerolog.Logger
}

func NewLogger(z zerolog.Logger) ILogger {
	return &loggerImpl{z: z}
}

func (l *loggerImpl) With(kv ...any) ILogger {
	// Initialize child logger
	e := l.z
	// Iterate pairs
	for i := 0; i+1 < len(kv); i += 2 {
		key := fmt.Sprint(kv[i])
		val := kv[i+1]
		e = e.With().Interface(key, val).Logger()
	}
	return &loggerImpl{z: e}
}

func (l *loggerImpl) log(level zerolog.Level, msg string, kv ...any) {
	// Create event based on level
	var e *zerolog.Event
	switch level {
	case zerolog.DebugLevel:
		e = l.z.Debug()
	case zerolog.InfoLevel:
		e = l.z.Info()
	case zerolog.WarnLevel:
		e = l.z.Warn()
	case zerolog.ErrorLevel:
		e = l.z.Error()
	case zerolog.FatalLevel:
		e = l.z.Fatal()
	default:
		e = l.z.Info()
	}
	// Attach fields
	for i := 0; i+1 < len(kv); i += 2 {
		key := fmt.Sprint(kv[i])
		val := kv[i+1]
		if errVal, ok := val.(error); ok {
			e = e.Err(errVal)
		} else {
			e = e.Interface(key, val)
		}
	}
	// If odd trailing arg, log as error field
	if len(kv)%2 == 1 {
		if errVal, ok := kv[len(kv)-1].(error); ok {
			e = e.Err(errVal)
		} else {
			e = e.Interface("args", kv[len(kv)-1])
		}
	}
	// Emit message
	e.Msg(msg)
}

func (l *loggerImpl) Debug(msg string, args ...any) {
	l.log(zerolog.DebugLevel, msg, args...)
}

func (l *loggerImpl) Info(msg string, args ...any) {
	l.log(zerolog.InfoLevel, msg, args...)
}

func (l *loggerImpl) Warn(msg string, args ...any) {
	l.log(zerolog.WarnLevel, msg, args...)
}

func (l *loggerImpl) Error(msg string, args ...any) {
	l.log(zerolog.ErrorLevel, msg, args...)
}

func (l *loggerImpl) Fatal(msg string, args ...any) {
	l.log(zerolog.FatalLevel, msg, args...)
	os.Exit(1)
}
