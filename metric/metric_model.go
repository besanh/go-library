package metric

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

// Metrics holds all registered Prometheus collectors and provides methods
// to interact with them.
// Metrics holds Prometheus metrics and configuration.
type Metrics struct {
	// internal registry (can be custom or default)
	registry *prometheus.Registry

	// HTTP request counter: labels = handler, method, status
	RequestCount *prometheus.CounterVec

	// HTTP request duration: labels = handler
	RequestDuration *prometheus.HistogramVec

	// Gauge for the number of in-flight requests
	InFlightGauge prometheus.Gauge

	// Protect concurrent custom collector registration
	mu sync.Mutex

	// Options flags
	disableDefault bool
}

// Config contains naming and bucket settings.
type Config struct {
	Namespace      string
	Subsystem      string
	RequestBuckets []float64
}

// Option configures Metrics behavior.
type Option func(*Metrics)
