package metric

import (
	"fmt"
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func containsPrefix(slice []string, prefix string) bool {
	for _, s := range slice {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}

	return false
}

func TestNewMetrics_DefaultBehavior(t *testing.T) {
	cfg := Config{
		Namespace:      "ns",
		Subsystem:      "sub",
		RequestBuckets: nil,
	}

	iface := NewMetrics(cfg)
	m := iface.(*Metrics) // assert the concrete type

	m.RequestCount.WithLabelValues("handler", "GET", "200").Add(1)
	m.RequestDuration.WithLabelValues("handler").Observe(0.1)
	m.InFlightGauge.Inc()

	mfs, err := m.registry.Gather()
	require.NoError(t, err)

	var names []string
	for _, mf := range mfs {
		names = append(names, mf.GetName())
	}

	assert.Contains(t, names, "ns_sub_request_count")
	assert.Contains(t, names, "ns_sub_request_duration_seconds")
	assert.Contains(t, names, "ns_sub_in_flight_requests")

	// Default collector should be registered
	assert.True(t, containsPrefix(names, "go_"), "expected at least one go_metric")
	assert.True(t, containsPrefix(names, "process_"), "expected at least one process_metric")
}

func TestNewMetrics_DisableDefaultCollectors(t *testing.T) {
	cfg := Config{Namespace: "ns", Subsystem: "sub", RequestBuckets: nil}

	iface := NewMetrics(cfg, WithoutDefaultCollectors())
	m := iface.(*Metrics) // assert the concrete type

	// exercise core metrics
	m.RequestCount.WithLabelValues("h", "POST", "500").Add(1)
	m.RequestDuration.WithLabelValues("h").Observe(0.2)
	m.InFlightGauge.Add(2)

	mfs, err := m.registry.Gather()
	require.NoError(t, err)

	var names []string
	for _, mf := range mfs {
		names = append(names, mf.GetName())
	}

	// core metrics should still be there
	assert.Contains(t, names, "ns_sub_request_count")
	assert.Contains(t, names, "ns_sub_request_duration_seconds")
	assert.Contains(t, names, "ns_sub_in_flight_requests")

	// but absolutely no "go_" or "process_" metrics
	for _, metricName := range names {
		assert.False(
			t,
			strings.HasPrefix(metricName, "go_"),
			fmt.Sprintf("unexpected default-collector metric: %q", metricName),
		)
		assert.False(
			t,
			strings.HasPrefix(metricName, "process_"),
			fmt.Sprintf("unexpected default-collector metric: %q", metricName),
		)
	}
}

func TestNewMetrics_WithRegistryOption(t *testing.T) {
	customReg := prometheus.NewRegistry()
	cfg := Config{
		Namespace:      "ns",
		Subsystem:      "sub",
		RequestBuckets: nil,
	}

	iface := NewMetrics(cfg, WithRegistry(customReg))
	m := iface.(*Metrics) // assert the concrete type

	m.RequestCount.WithLabelValues("handler", "PUT", "201").Add(1)

	mfs, err := m.registry.Gather()
	require.NoError(t, err)

	var names []string
	for _, mf := range mfs {
		names = append(names, mf.GetName())
	}

	assert.Contains(t, names, "ns_sub_request_count")
	assert.False(t, containsPrefix(names, "ns_sub_request_duration_secondss"))
	assert.False(t, containsPrefix(names, "ns_sub_in_flight_requestss"))
}
