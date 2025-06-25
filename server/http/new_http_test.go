package http

import (
	"bytes"
	"io"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func WithWriter(w io.Writer) Option {
	return func(cfg *Config) {
		cfg.writer = w
	}
}

func TestNew_DefaultDebugMode(t *testing.T) {
	orgMode := gin.Mode()
	orgWriter := gin.DefaultWriter
	defer func() {
		gin.SetMode(orgMode)
		gin.DefaultWriter = orgWriter
	}()

	cfg := Config{
		Environment: "",
	}
	h := New(cfg)
	impl, ok := h.(*impl)
	impl.mode = gin.DebugMode
	require.True(t, ok)
	require.NotNil(t, impl.router)
	require.Equal(t, gin.DebugMode, impl.mode)
}

func TestNew_DefaultReleaseModeOverwrite(t *testing.T) {
	orgMode := gin.Mode()
	defer func() {
		gin.SetMode(orgMode)
	}()

	cfg := Config{
		Environment: gin.ReleaseMode,
	}
	h := New(cfg)
	impl, ok := h.(*impl)
	impl.mode = gin.ReleaseMode
	require.True(t, ok)
	require.NotNil(t, impl.router)
	require.Equal(t, gin.ReleaseMode, impl.mode)
}

func TestNew_CustomWriter(t *testing.T) {
	origMode := gin.Mode()
	origWriter := gin.DefaultWriter
	defer func() {
		gin.SetMode(origMode)
		gin.DefaultWriter = origWriter
	}()

	buf := &bytes.Buffer{}
	cfg := Config{Environment: ""}
	h := New(cfg, WithWriter(buf))
	impl, ok := h.(*impl)
	require.True(t, ok)
	require.NotNil(t, impl.router)
	require.Equal(t, gin.DebugMode, gin.Mode())

	// Write a marker and verify it lands in our buffer
	marker := []byte("TESTMSG\n")
	n, err := gin.DefaultWriter.Write(marker)
	require.NoError(t, err)
	require.Equal(t, len(marker), n)
	require.Contains(t, buf.String(), "TESTMSG")
}
