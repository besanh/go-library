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
	Level        zapcore.Level
	EnableSentry bool
	SentryDSN    string
	HTTPHookURL  string
	ServiceName  string
}

type Option func(*Config)

func defaultConfig() Config {
	return Config{
		Level: zapcore.InfoLevel,
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

	// Output to stdout + optional HTTP hook
	cores := []zapcore.Core{
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), cfg.Level),
	}
	if cfg.HTTPHookURL != "" {
		hook := &httpclient.HttpHookWriter{
			URL: cfg.HTTPHookURL,
		}
		cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(hook), cfg.Level))
	}
	tee := zapcore.NewTee(cores...)

	var core zapcore.Core = tee

	if cfg.EnableSentry {
		if err := sentry.Init(sentry.ClientOptions{
			Dsn: cfg.SentryDSN,
		}); err != nil {
			return nil, err
		}
		core = &sentryhook.SentryCore{
			Core:     tee,
			MinLevel: zapcore.ErrorLevel,
		}
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
