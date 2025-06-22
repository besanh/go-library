package sentryhook

import (
	"testing"

	sentrypkg "github.com/getsentry/sentry-go"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

// TestNewSentryHook_ReturnsNonNil ensures the constructor returns a non-nil Hook.
func TestNewSentryHook_ReturnsNonNil(t *testing.T) {
	hook := NewSentryHook("https://example.com")
	assert.NotNil(t, hook)
}

// TestNewSentryHook_VariadicOptions ensures the constructor accepts variadic SentryOption.
func TestNewSentryHook_VariadicOptions(t *testing.T) {
	before := func(o *sentrypkg.ClientOptions) {
		o.Dsn = ""
		o.BeforeSend = func(event *sentrypkg.Event, hint *sentrypkg.EventHint) *sentrypkg.Event {
			return event
		}
	}
	hook := NewSentryHook("https://example.com", before)
	assert.NotNil(t, hook)
}

// TestNewSentryHook_ImplementsZerologHook verifies the returned object implements zerolog.Hook.
func TestNewSentryHook_ImplementsZerologHook(t *testing.T) {
	hook := NewSentryHook("dsn")
	// Type assertion
	assert.True(t, hook == zerolog.Hook(hook))
}
