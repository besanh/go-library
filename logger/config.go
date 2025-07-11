package logger

import "go.uber.org/zap/zapcore"

func WithServiceName(name string) Option {
	return func(c *Config) {
		c.ServiceName = name
	}
}

func WithLevel(level zapcore.Level) Option {
	return func(c *Config) {
		c.Level = level
	}
}

func WithHTTPHook(url string) Option {
	return func(c *Config) {
		c.HTTPHookURL = url
	}
}

func WithSentry(dsn string) Option {
	return func(c *Config) {
		c.EnableSentry = true
		c.SentryDSN = dsn
	}
}
