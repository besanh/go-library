package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

// NewMetrics builds and registers metrics according to Config and Options.
func NewMetrics(cfg Config, opts ...Option) IMetric {
	// choose buckets
	buckets := cfg.RequestBuckets
	if len(buckets) == 0 {
		buckets = prometheus.DefBuckets
	}

	// initialize registry if not provided
	effectiveReg := prometheus.NewRegistry()
	m := &Metrics{
		registry: effectiveReg,
		RequestCount: prometheus.NewCounterVec(
			prometheus.CounterOpts{Namespace: cfg.Namespace, Subsystem: cfg.Subsystem, Name: "request_count", Help: "Total HTTP requests."},
			[]string{"handler", "method", "status"},
		),
		RequestDuration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{Namespace: cfg.Namespace, Subsystem: cfg.Subsystem, Name: "request_duration_seconds", Help: "Duration of HTTP requests.", Buckets: buckets},
			[]string{"handler"},
		),
		InFlightGauge: prometheus.NewGauge(prometheus.GaugeOpts{Namespace: cfg.Namespace, Subsystem: cfg.Subsystem, Name: "in_flight_requests", Help: "In-flight HTTP requests."}),
	}

	// apply functional options
	for _, opt := range opts {
		opt(m)
	}

	// register core metrics
	m.registry.MustRegister(m.RequestCount, m.RequestDuration, m.InFlightGauge)

	// default collectors
	if !m.disableDefault {
		m.registry.MustRegister(collectors.NewGoCollector(), collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
	}

	return m
}
