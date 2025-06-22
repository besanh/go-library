package httpclient

import (
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

// TestNewHTTPHook_ReturnsNonNil ensures the constructor returns a non-nil Hook.
func TestNewHTTPHook_ReturnsNonNil(t *testing.T) {
	hook := NewHTTPHook("http://example.com/log", nil)
	assert.NotNil(t, hook)
}

// TestNewHTTPHook_DefaultClient ensures a nil client uses http.DefaultClient without panic.
func TestNewHTTPHook_DefaultClient(t *testing.T) {
	hook := NewHTTPHook("http://example.com/log", nil)
	// Should not panic when Run is called
	assert.NotPanics(t, func() {
		hook.Run(nil, zerolog.InfoLevel, "test_message")
	})
}

// TestNewHTTPHook_ImplementsZerologHook verifies the returned object implements zerolog.Hook.
func TestNewHTTPHook_ImplementsZerologHook(t *testing.T) {
	hook := NewHTTPHook("http://example.com/log", nil)
	assert.True(t, hook == zerolog.Hook(hook))
}
