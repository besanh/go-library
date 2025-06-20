package metric

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// IMetric exposes only the methods you want public.
type IMetric interface {
	// Handler returns an http.Handler to serve the /metrics endpoint.
	Handler() http.Handler

	// Instrument wraps an http.Handler, tracking count, duration, and in-flight.
	Instrument(handlerName string, next http.Handler) http.Handler

	// RegisterCustom allows adding a collector at runtime.
	RegisterCustom(c prometheus.Collector) error
}

// Handler implements IMetric.Handler
func (m *Metrics) Handler() http.Handler {
	return promhttp.HandlerFor(m.registry, promhttp.HandlerOpts{})
}

// Instrument implements IMetric.Instrument
func (m *Metrics) Instrument(handlerName string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.InFlightGauge.Inc()
		defer m.InFlightGauge.Dec()

		sw := &statusWriter{ResponseWriter: w}
		t := prometheus.NewTimer(m.RequestDuration.WithLabelValues(handlerName))
		defer t.ObserveDuration()

		next.ServeHTTP(sw, r)

		status := sw.Status()
		m.RequestCount.WithLabelValues(handlerName, r.Method, status).Inc()
	})
}

// RegisterCustom implements IMetric.RegisterCustom
func (m *Metrics) RegisterCustom(c prometheus.Collector) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	typ := fmt.Sprintf("%T", c)
	existing, _ := m.registry.Gather()
	for _, mf := range existing {
		if mf.GetName() == typ {
			return fmt.Errorf("collector already registered: %s", typ)
		}
	}
	return m.registry.Register(c)
}
