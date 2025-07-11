package logger

func WithServiceName(name string) Option {
	return func(c *Config) {
		c.ServiceName = name
	}
}

func WithLevel(level string) Option {
	return func(c *Config) {
		switch level {
		case DEBUG_LEVEL, INFO_LEVEL, WARN_LEVEL, ERROR_LEVEL, FATAL_LEVEL:
			c.Level = level
			return
		default:
			c.Level = INFO_LEVEL
			return
		}
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
