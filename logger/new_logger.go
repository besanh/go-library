package logger

import (
	"os"

	"github.com/besanh/go-library/logger/httpclient"
	"github.com/besanh/go-library/logger/sentryhook"
	"github.com/getsentry/sentry-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Level        string
	EnableSentry bool
	SentryDSN    string
	HTTPHookURL  string
	ServiceName  string
}

type Option func(*Config)

func defaultConfig() Config {
	return Config{
		Level: INFO_LEVEL,
	}
}

type impl struct {
	core *zap.Logger
}

func NewLogger(opts ...Option) (ILogger, error) {
	cfg := defaultConfig()
	for _, opt := range opts {
		opt(&cfg)
	}

	encCfg := zapcore.EncoderConfig{
		MessageKey:   "message",
		LevelKey:     "level",
		TimeKey:      "time",
		NameKey:      "logger",
		CallerKey:    "file",
		FunctionKey:  "func",
		LineEnding:   zapcore.DefaultLineEnding,
		EncodeLevel:  zapcore.LowercaseLevelEncoder,
		EncodeTime:   zapcore.ISO8601TimeEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	}
	encoder := zapcore.NewJSONEncoder(encCfg)
	lv := converLevel(cfg.Level)

	cores := []zapcore.Core{
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), lv),
	}
	if cfg.HTTPHookURL != "" {
		hook := &httpclient.HttpHookWriter{
			URL: cfg.HTTPHookURL,
		}
		cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(hook), lv))
	}
	tee := zapcore.NewTee(cores...)

	var core zapcore.Core = tee

	if cfg.EnableSentry {
		if err := sentry.Init(sentry.ClientOptions{
			Dsn: cfg.SentryDSN,
		}); err != nil {
			return nil, err
		}
		sentryCore := &sentryhook.SentryCore{
			Core:     zapcore.NewNopCore(),
			MinLevel: zapcore.ErrorLevel,
		}

		core = zapcore.NewTee(tee, sentryCore)
	}

	base := zap.New(core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)

	if cfg.ServiceName != "" {
		base = base.Named(cfg.ServiceName)
	}

	return &impl{
		core: base,
	}, nil
}

func converLevel(lv string) zapcore.Level {
	switch lv {
	case DEBUG_LEVEL:
		return zapcore.DebugLevel
	case INFO_LEVEL:
		return zapcore.InfoLevel
	case WARN_LEVEL:
		return zapcore.WarnLevel
	case ERROR_LEVEL:
		return zapcore.ErrorLevel
	case FATAL_LEVEL:
		return zapcore.FatalLevel
	default:
		return zap.ErrorLevel
	}
}
