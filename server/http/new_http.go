package http

import (
	"io"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Environment string
	writer      io.Writer
}

type impl struct {
	mode           string
	router         *gin.Engine
	beforeShutdown func()
}

func New(cfg Config, opts ...Option) IHTTP {
	gin.SetMode(gin.DebugMode)
	for _, opt := range opts {
		opt(&cfg)
	}
	if strings.TrimSpace(cfg.Environment) != "" {
		gin.SetMode(cfg.Environment)
	}

	// Log Writers
	var wrs []io.Writer
	if cfg.Environment != gin.ReleaseMode {
		wrs = append(wrs, os.Stdout)
	}
	if cfg.writer != nil {
		wrs = append(wrs, cfg.writer)
	}
	gin.DefaultWriter = io.MultiWriter(wrs...)

	// Create router
	router := gin.Default()

	return &impl{
		router: router,
	}
}
