package logger

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"testing"
	"time"

	"github.com/besanh/go-library/logger/sentryhook"
	sentryGo "github.com/getsentry/sentry-go"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestInitAndDefault(t *testing.T) {
	// Use a buffer to capture JSON output
	buf := &bytes.Buffer{}
	Init(
		WithJSONWriter(buf),
		WithLevel(zerolog.DebugLevel),
	)

	def1 := Default()
	def2 := Default()
	// Default() should return the same Logger instance
	assert.Same(t, def1, def2)
}

func TestNewAddsComponentField(t *testing.T) {
	// Ensure init with buffer
	buf := &bytes.Buffer{}
	Init(
		WithJSONWriter(buf),
		WithLevel(zerolog.DebugLevel),
	)

	componentName := "mycomponent"
	log := New(componentName)
	log.Info("test_message")

	// Parse output JSON and verify component field
	out := buf.String()
	var msg map[string]any
	err := json.Unmarshal([]byte(out), &msg)
	assert.NoError(t, err)
	assert.Equal(t, "test_message", msg["message"])
}

func TestWithLevel(t *testing.T) {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	buf := &bytes.Buffer{}
	var z zerolog.Logger = zerolog.New(buf)
	WithLevel(zerolog.DebugLevel)(&z)

	assert.Equal(t, zerolog.DebugLevel, zerolog.GlobalLevel())
	assert.Equal(t, zerolog.DebugLevel, z.GetLevel())
}

func TestWithTimeFormat(t *testing.T) {
	org := zerolog.TimeFieldFormat
	defer func() {
		zerolog.TimeFieldFormat = org
	}()
	var z zerolog.Logger = zerolog.New(ioutil.Discard)
	WithTimeFormat(time.RFC3339Nano)(&z)

	assert.Equal(t, time.RFC3339Nano, zerolog.TimeFieldFormat)
}

func TestWithJSONWriter(t *testing.T) {
	buf := &bytes.Buffer{}
	var z zerolog.Logger = zerolog.New(buf)
	WithJSONWriter(buf)(&z)
	z.Info().Msg("hello_json_writer")

	out := buf.String()
	var obj map[string]any
	err := json.Unmarshal([]byte(out), &obj)

	assert.NoError(t, err)
	assert.Equal(t, "hello_json_writer", obj["message"])
}

func TestWithConsoleWriter(t *testing.T) {
	buf := &bytes.Buffer{}
	var z zerolog.Logger = zerolog.New(buf)
	// use a fixed layout without sub-second precision
	layout := "2006-01-02T15:04:05"
	WithConsoleWriter(buf, layout)(&z)
	z.Info().Msg("hello_console_writer")

	out := buf.String()
	assert.Contains(t, out, "hello_console_writer")
	// only check up to seconds to avoid nanosecond mismatches
	prefix := time.Now().Format(layout)
	assert.Contains(t, out, prefix)
}

func TestConvertLevel(t *testing.T) {
	assert.Equal(t, sentryGo.LevelDebug, sentryhook.Convert(zerolog.DebugLevel))
	assert.Equal(t, sentryGo.LevelInfo, sentryhook.Convert(zerolog.InfoLevel))
	assert.Equal(t, sentryGo.LevelWarning, sentryhook.Convert(zerolog.WarnLevel))
	assert.Equal(t, sentryGo.LevelError, sentryhook.Convert(zerolog.ErrorLevel))
	assert.Equal(t, sentryGo.LevelFatal, sentryhook.Convert(zerolog.FatalLevel))
}
