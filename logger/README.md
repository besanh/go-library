# Logger Package

The `logger` package provides a standardized, high-performance structured logging implementation built on top of `go.uber.org/zap`. It also provides modular extensions for error tracking via Sentry and pushing logs over HTTP.

## Architecture Flow

1. **Core Processing**: The logger constructs a customized `zap.Logger` using a unified `Config` and the functional options pattern. 
2. **Output Multiplexing**: Log entries are piped through a `zapcore.Tee` structure:
   - **`stdout`**: Standard console output equipped with a JSON format encoder, caller skipping, and stack traces on error.
   - **HTTP Hook (Optional)**: If configured, log entries are broadcasted via HTTP POST to the specified `HTTPHookURL`.
   - **Sentry Hook (Optional)**: If `EnableSentry` is enabled, entries at the `ERROR` or `FATAL` level are caught and forwarded directly to the configured Sentry DSN.

## Usage

```go
import "github.com/besanh/go-library/logger"
```

### 1. Basic Initialization

Initialize the logger with standard console output at `INFO` level.

```go
log, err := logger.NewLogger(
    func(c *logger.Config) {
        c.Level = logger.INFO_LEVEL
        c.ServiceName = "my-service"
    },
)
if err != nil {
    panic(err)
}

log.Info("Service started")
```

### 2. Integration with Sentry

Configure the logger to broadcast `ERROR` level occurrences directly to Sentry.

```go
log, err := logger.NewLogger(
    func(c *logger.Config) {
        c.Level = logger.DEBUG_LEVEL
        c.ServiceName = "user-service"
        c.EnableSentry = true
        c.SentryDSN = "https://example@sentry.io/123"
    },
)

// This will be logged to stdout AND captured as an event in Sentry
log.Errorf("Failed to connect to database: %v", dbErr)
```

### 3. Integration with HTTP Hooks

Configure the logger to pipe its JSON output to an external HTTP sink.

```go
log, err := logger.NewLogger(
    func(c *logger.Config) {
        c.Level = logger.INFO_LEVEL
        c.HTTPHookURL = "https://my-log-collector.internal/ingest"
    },
)
// Logs will now be JSON-encoded and POSTed to the HTTPHookURL
```
