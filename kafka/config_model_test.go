package kafka

import "testing"

func TestKafkaConfig_Validate(t *testing.T) {
	testCases := []struct {
		name        string
		cfg         KafkaConfig
		wantErr     bool
		wantVersion string
		wantOffset  string
	}{
		{
			name:    "missing brokers",
			cfg:     KafkaConfig{},
			wantErr: true,
		},
		{
			name:    "missing topic",
			cfg:     KafkaConfig{Brokers: []string{"b1:9092"}},
			wantErr: true,
		},
		{
			name:    "missing group",
			cfg:     KafkaConfig{Brokers: []string{"b1:9092"}, Topic: "t1"},
			wantErr: true,
		},
		{
			name:        "defaults set",
			cfg:         KafkaConfig{Brokers: []string{"b1:9092"}, Topic: "t1", ConsumerGroup: "g1", Version: "2.6.0", ConsumerOffset: "oldest"},
			wantErr:     false,
			wantVersion: "2.6.0",
			wantOffset:  "oldest",
		},
		{
			name: "custom version and offset",
			cfg: KafkaConfig{
				Brokers: []string{"b1:9092"}, Topic: "t1", ConsumerGroup: "g1",
				Version: "2.3.1", ConsumerOffset: "newest",
			},
			wantErr:     false,
			wantVersion: "2.3.1",
			wantOffset:  "newest",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.cfg.Validate()
			if (err != nil) != tc.wantErr {
				t.Fatalf("got error %v, want error %v", err, tc.wantErr)
			}
			if !tc.wantErr {
				if tc.cfg.Version != tc.wantVersion {
					t.Fatalf("got version %v, want version %v", tc.cfg.Version, tc.wantVersion)
				}
				if tc.cfg.ConsumerOffset != tc.wantOffset {
					t.Fatalf("got offset %v, want offset %v", tc.cfg.ConsumerOffset, tc.wantOffset)
				}
			}
		})
	}
}
