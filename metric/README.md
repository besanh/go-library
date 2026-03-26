# Metric Package

The `metric` package provides a standardized telemetry wrapper using the `github.com/prometheus/client_golang` library. It offers pre-configured metrics for HTTP traffic analysis, simplifying observability across services.

## Architecture Flow

1. **Initialization**: The package is initialized with a `Config` holding namespace and subsystem metadata.
2. **Registration**: Upon initialization, it instantiates a Prometheus Registry and registers core HTTP metrics (Counters, Histograms, Gauges). Unless disabled, it also mounts the default Go runtime and process collectors.
3. **Observation**: Downstream HTTP handlers or middlewares use the exposed metric instances to increment counts or record durations dynamically.

## Usage

```go
import "github.com/besanh/go-library/metric"
```

### 1. Initialization

Initialize the unified metrics struct with your application's namespace.

```go
cfg := metric.Config{
    Namespace: "my_company",
    Subsystem: "auth_service",
    // RequestBuckets: ... 
}

// Optional functional parameters can be passed to apply custom behavior (e.g. disable default collectors)
metrics := metric.NewMetrics(cfg)
```

### 2. Available Standard Metrics

The `IMetric` interfaces expose the following standardized vectors:

- **`RequestCount`** (`*prometheus.CounterVec`): Tracks total HTTP requests. (Labels: `handler`, `method`, `status`)
- **`RequestDuration`** (`*prometheus.HistogramVec`): Tracks the duration of HTTP requests across configured buckets. (Labels: `handler`)
- **`InFlightGauge`** (`prometheus.Gauge`): Tracks currently active, in-flight HTTP requests.

### 3. Emitting Metrics

Integrate within your HTTP middleware to track operations automatically.

```go
func (m *Metrics) Middleware(handler string, next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Increment in-flight
        m.InFlightGauge.Inc()
        defer m.InFlightGauge.Dec()
        
        timer := prometheus.NewTimer(m.RequestDuration.WithLabelValues(handler))
        defer timer.ObserveDuration()

        // Pass to next
        next.ServeHTTP(w, r)
        
        // Count request completed (assuming custom ResponseWriter to capture status)
        m.RequestCount.WithLabelValues(handler, r.Method, "200").Inc()
    })
}
```