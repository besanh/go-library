package metric

import "github.com/prometheus/client_golang/prometheus"

// WithRegistry injects a custom Prometheus registry.
func WithRegistry(reg *prometheus.Registry) Option {
	return func(m *Metrics) {
		m.registry = reg
	}
}

// WithoutDefaultCollectors disables Go and process collectors.
func WithoutDefaultCollectors() Option {
	return func(m *Metrics) {
		m.disableDefault = true
	}
}

// WithCustomCollectors pre-registers provided collectors.
func WithCustomCollectors(cs ...prometheus.Collector) Option {
	return func(m *Metrics) {
		for _, c := range cs {
			m.registry.MustRegister(c)
		}
	}
}
