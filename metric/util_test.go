package metric

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
)

func TestWithRegistry(t *testing.T) {
	testCases := []struct {
		name     string
		initial  *prometheus.Registry
		input    *prometheus.Registry
		expected *prometheus.Registry
	}{
		{
			name:     "set to new registry",
			initial:  nil,
			input:    prometheus.NewRegistry(),
			expected: prometheus.NewRegistry(),
		},
		{
			name:     "overwrite existing registry",
			initial:  prometheus.NewRegistry(),
			input:    prometheus.NewRegistry(),
			expected: prometheus.NewRegistry(),
		},
		{
			name:     "set to nil",
			initial:  prometheus.NewRegistry(),
			input:    nil,
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			m := &Metrics{
				registry: tc.initial,
			}
			WithRegistry(tc.input)(m)
			assert.Equal(t, tc.input, m.registry)
		})
	}
}

func TestWithoutDefaultCollectors(t *testing.T) {
	testCases := []struct {
		name             string
		initalDisabled   bool
		expectedDisabled bool
	}{
		{
			name:             "disable when initialy false",
			initalDisabled:   false,
			expectedDisabled: true,
		},
		{
			name:             "remain disable when initialy true",
			initalDisabled:   true,
			expectedDisabled: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			m := &Metrics{
				disableDefault: tc.initalDisabled,
			}
			WithoutDefaultCollectors()(m)
			assert.Equal(t, tc.expectedDisabled, m.disableDefault)
		})
	}
}

func TestWithCustomCollectors(t *testing.T) {
	tests := []struct {
		name          string
		collectors    []prometheus.Collector
		expectedNames []string
	}{
		{
			name:          "no collectors",
			collectors:    []prometheus.Collector{},
			expectedNames: []string{},
		},
		{
			name: "single collector",
			collectors: []prometheus.Collector{
				prometheus.NewGauge(prometheus.GaugeOpts{
					Name: "metric_one",
					Help: "test gauge one",
				}),
			},
			expectedNames: []string{"metric_one"},
		},
		{
			name: "multiple collectors",
			collectors: []prometheus.Collector{
				prometheus.NewGauge(prometheus.GaugeOpts{
					Name: "metric_one",
					Help: "test gauge one",
				}),
				prometheus.NewGauge(prometheus.GaugeOpts{
					Name: "metric_two",
					Help: "test gauge two",
				}),
			},
			expectedNames: []string{"metric_one", "metric_two"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// fresh empty registry
			r := prometheus.NewRegistry()
			m := &Metrics{registry: r}

			// apply the option under test
			WithCustomCollectors(tt.collectors...)(m)

			// gather what's registered
			mfs, err := r.Gather()
			assert.NoError(t, err, "Gather should not error")

			// extract the names
			var gotNames []string
			for _, mf := range mfs {
				gotNames = append(gotNames, mf.GetName())
			}
			t.Logf("registered metrics: %v", gotNames)

			assert.Len(t, gotNames, len(tt.expectedNames),
				"expected %d metrics, got %d", len(tt.expectedNames), len(gotNames),
			)

			for _, want := range tt.expectedNames {
				assert.Contains(t, gotNames, want,
					"expected collector %q to be registered", want,
				)
			}
		})
	}
}
