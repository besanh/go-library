package logger

import (
	"io"
	"os"
	"sync"

	"github.com/besanh/go-library/logger/httpclient"
	"github.com/besanh/go-library/logger/sentryhook"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Field struct {
	Key   string
	Value any
}

type Option func(*zerolog.Logger)

var (
	defaultLogger ILogger
	once          sync.Once
)

func Init(opts ...Option) {
	once.Do(func() {
		z := zerolog.New(os.Stdout).
			With().
			Timestamp().
			Logger()
		for _, opt := range opts {
			opt(&z)
		}
		defaultLogger = &loggerImpl{z}
		log.Logger = z
	})
}

func Default() ILogger {
	if defaultLogger == nil {
		Init()
	}
	return defaultLogger
}

func New(component string) ILogger {
	return Default().With(Field{Key: "component", Value: component})
}

func WithLevel(level zerolog.Level) Option {
	return func(z *zerolog.Logger) {
		zerolog.SetGlobalLevel(level)
		*z = z.Level(level)
	}
}

func WithConsoleWriter(w io.Writer, timeFmt string) Option {
	return func(z *zerolog.Logger) {
		cw := zerolog.ConsoleWriter{Out: w, TimeFormat: timeFmt}
		*z = zerolog.New(cw).With().Timestamp().Logger()
	}
}

func WithJSONWriter(w io.Writer) Option {
	return func(z *zerolog.Logger) {
		*z = zerolog.New(w).With().Timestamp().Logger()
	}
}

func WithTimeFormat(format string) Option {
	return func(z *zerolog.Logger) {
		zerolog.TimeFieldFormat = format
		*z = z.With().
			Timestamp().
			Logger()
	}
}

func WithSentryHook(dsn string, opts ...sentryhook.SentryOption) Option {
	return func(z *zerolog.Logger) {
		hook := sentryhook.NewSentryHook(dsn, opts...)
		*z = z.Hook(hook)
	}
}

func WithHTTPHook(url string, client httpclient.IHTTPClient) Option {
	return func(z *zerolog.Logger) {
		hook := httpclient.NewHTTPHook(url, client)
		*z = z.Hook(hook)
	}
}
