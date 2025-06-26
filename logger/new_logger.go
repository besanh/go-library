// logger/new_logger.go
package logger

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"sync"

	"github.com/besanh/go-library/logger/httpclient"
	"github.com/besanh/go-library/logger/sentryhook"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Field represents a key/value pair to log.
type Field struct {
	Key   string
	Value any
}

// Option configures the zerolog.Logger during initialization.
type Option func(*zerolog.Logger)

// colorJSONWriter wraps an io.Writer and colorizes JSON keys and values.
type colorJSONWriter struct {
	w io.Writer
}

var (
	// regex patterns to match JSON keys and values
	keyPattern   = regexp.MustCompile(`"(\w+)":`)
	valuePattern = regexp.MustCompile(`: "(.*?)"`)
)

// Write colorizes JSON log entries: keys in cyan, values in yellow.
func (c *colorJSONWriter) Write(p []byte) (n int, err error) {
	s := string(p)
	// color keys
	s = keyPattern.ReplaceAllStringFunc(s, func(m string) string {
		sub := keyPattern.FindStringSubmatch(m)
		return fmt.Sprintf("\x1b[36m\"%s\"\x1b[0m:", sub[1])
	})
	// color values
	s = valuePattern.ReplaceAllStringFunc(s, func(m string) string {
		sub := valuePattern.FindStringSubmatch(m)
		return fmt.Sprintf(": \x1b[33m\"%s\"\x1b[0m", sub[1])
	})
	return c.w.Write([]byte(s))
}

var (
	defaultLogger ILogger
	once          sync.Once
)

// Init initializes the global default logger with colored JSON output.
// Logs remain in JSON format but keys and values are colorized for readability.
func Init(opts ...Option) {
	once.Do(func() {
		// Wrap stdout with JSON colorizer
		w := &colorJSONWriter{w: os.Stdout}

		// Base logger uses JSON writer
		z := zerolog.New(w).
			With().
			Timestamp().
			Logger()

		// Apply additional options (hooks, level)
		for _, opt := range opts {
			opt(&z)
		}

		defaultLogger = &loggerImpl{z}
		log.Logger = z
	})
}

// Default returns the singleton ILogger, initializing it if necessary.
func Default() ILogger {
	if defaultLogger == nil {
		Init()
	}
	return defaultLogger
}

// New creates a child logger with the "component" field set.
func New(component string) ILogger {
	return Default().With(Field{Key: "component", Value: component})
}

// WithLevel returns an Option to set the global log level.
func WithLevel(level zerolog.Level) Option {
	return func(z *zerolog.Logger) {
		zerolog.SetGlobalLevel(level)
		*z = z.Level(level)
	}
}

// WithConsoleWriter returns an Option to switch to console output (colored) instead of JSON.
func WithConsoleWriter(w io.Writer, timeFmt string) Option {
	return func(z *zerolog.Logger) {
		cw := zerolog.ConsoleWriter{Out: w, TimeFormat: timeFmt, NoColor: false}
		*z = zerolog.New(cw).With().Timestamp().Logger()
	}
}

// WithJSONWriter returns an Option to switch to uncolored JSON output.
func WithJSONWriter(w io.Writer) Option {
	return func(z *zerolog.Logger) {
		*z = zerolog.New(w).With().Timestamp().Logger()
	}
}

// WithTimeFormat sets the timestamp format.
func WithTimeFormat(format string) Option {
	return func(z *zerolog.Logger) {
		zerolog.TimeFieldFormat = format
		*z = z.With().Timestamp().Logger()
	}
}

// WithSentryHook adds a Sentry hook for error reporting.
func WithSentryHook(dsn string, opts ...sentryhook.SentryOption) Option {
	return func(z *zerolog.Logger) {
		hook := sentryhook.NewSentryHook(dsn, opts...)
		*z = z.Hook(hook)
	}
}

// WithHTTPHook adds an HTTP hook for log shipping.
func WithHTTPHook(url string, client httpclient.IHTTPClient) Option {
	return func(z *zerolog.Logger) {
		hook := httpclient.NewHTTPHook(url, client)
		*z = z.Hook(hook)
	}
}
