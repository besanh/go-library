package kafka

import (
	"context"
	"errors"
	"testing"

	"github.com/IBM/sarama"
	"github.com/stretchr/testify/require"
)

// dummyHandler satisfies IMessageHandler for tests
type dummyHandler struct{}

func (h *dummyHandler) HandleMessage(msg *sarama.ConsumerMessage) error { return nil }

// stubConsumerGroup is a no-op Sarama ConsumerGroup implementation
type stubConsumerGroup struct{}

func (s *stubConsumerGroup) Consume(ctx context.Context, topics []string, handler sarama.ConsumerGroupHandler) error {
	return nil
}
func (s *stubConsumerGroup) Errors() <-chan error             { return nil }
func (s *stubConsumerGroup) Close() error                     { return nil }
func (s *stubConsumerGroup) Pause(topics map[string][]int32)  {}
func (s *stubConsumerGroup) PauseAll()                        {}
func (s *stubConsumerGroup) Resume(topics map[string][]int32) {}
func (s *stubConsumerGroup) ResumeAll()                       {}

func TestNewConsumerGroup_Success(t *testing.T) {
	// override builder
	orig := consumerGroupBuilder
	consumerGroupBuilder = func(cfg *KafkaConfig) (sarama.ConsumerGroup, error) {
		return &stubConsumerGroup{}, nil
	}
	defer func() { consumerGroupBuilder = orig }()

	cfg := &KafkaConfig{
		Brokers:       []string{"localhost:9092"},
		Topic:         "topic1",
		ConsumerGroup: "group1",
		Version:       "2.6.0",
	}
	h := &dummyHandler{}
	cg, err := NewConsumerGroup(cfg, h)
	require.NoError(t, err)

	impl, ok := cg.(*ConsumerGroup)
	require.True(t, ok, "expected *ConsumerGroup, got %T", cg)
	require.Equal(t, cfg.Topic, impl.topic)
	require.Equal(t, h, impl.handler)
}

func TestNewConsumerGroup_BuilderError(t *testing.T) {
	// override builder to return error
	orig := consumerGroupBuilder
	consumerGroupBuilder = func(cfg *KafkaConfig) (sarama.ConsumerGroup, error) {
		return nil, errors.New("builder failed")
	}
	defer func() { consumerGroupBuilder = orig }()

	cfg := &KafkaConfig{Brokers: []string{"b1:9092"}, Topic: "t", ConsumerGroup: "g"}
	h := &dummyHandler{}
	cg, err := NewConsumerGroup(cfg, h)
	require.Error(t, err)
	require.Nil(t, cg)
}
