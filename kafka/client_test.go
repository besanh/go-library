package kafka

import (
	"testing"
	"time"

	"github.com/IBM/sarama"
)

func TestBuildSamaraConfig(t *testing.T) {
	cfg := KafkaConfig{
		Brokers:        []string{"localhost:9092"},
		Topic:          "test",
		ConsumerGroup:  "group",
		Version:        "2.6.0",
		RetryMax:       5,
		RetryBackoff:   50 * time.Millisecond,
		ConsumerOffset: "newest",
	}

	sc, err := buildSaramaConfig(&cfg)
	if err != nil {
		t.Fatalf("failed to build sarama config: %v", err)
	}

	v, err := sarama.ParseKafkaVersion(sc.Version.String())
	if err != nil {
		t.Fatalf("failed to parse kafka version: %v", err)
	}

	if sc.Version != v {
		t.Errorf("expected version %v, got %v", v, sc.Version)
	}

	if sc.Producer.Retry.Max != cfg.RetryMax {
		t.Errorf("expected retry max %v, got %v", cfg.RetryMax, sc.Producer.Retry.Max)
	}

	if sc.Producer.Retry.Backoff != cfg.RetryBackoff {
		t.Errorf("expected retry backoff %v, got %v", cfg.RetryBackoff, sc.Producer.Retry.Backoff)
	}

	if sc.Consumer.Offsets.Initial != sarama.OffsetNewest {
		t.Errorf("expected offset initial %v, got %v", sarama.OffsetNewest, sc.Consumer.Offsets.Initial)
	}
}

func TestNewSyncProducer_Error(t *testing.T) {
	bad := &KafkaConfig{
		Brokers:       []string{"bad:0000"},
		Topic:         "t",
		ConsumerGroup: "g",
	}
	if _, err := NewSyncProducer(bad); err == nil {
		t.Error("expected error creating SyncProducer with bad broker, got nil")
	}
}

func TestNewConsumerGroupClient_Error(t *testing.T) {
	bad := &KafkaConfig{
		Brokers:       []string{"bad:0000"},
		Topic:         "t",
		ConsumerGroup: "g",
	}
	if _, err := NewConsumerGroupClient(bad); err == nil {
		t.Error("expected error creating ConsumerGroup with bad broker, got nil")
	}
}
