package kafka

import (
	"errors"
	"testing"

	"github.com/IBM/sarama"
	"github.com/stretchr/testify/require"
)

// stubSyncProducer implements the full sarama.SyncProducer interface for testing.
type stubSyncProducer struct{}

// AddMessageToTxn implements sarama.SyncProducer.
func (s *stubSyncProducer) AddMessageToTxn(msg *sarama.ConsumerMessage, groupId string, metadata *string) error {
	panic("unimplemented")
}

func (s *stubSyncProducer) SendMessage(msg *sarama.ProducerMessage) (int32, int64, error) {
	return 0, 0, nil
}
func (s *stubSyncProducer) SendMessages(msgs []*sarama.ProducerMessage) error {
	return nil
}
func (s *stubSyncProducer) Close() error {
	return nil
}
func (s *stubSyncProducer) TxnStatus() sarama.ProducerTxnStatusFlag {
	return sarama.ProducerTxnStatusFlag(0)
}
func (s *stubSyncProducer) IsTransactional() bool {
	return false
}
func (s *stubSyncProducer) BeginTxn() error {
	return nil
}
func (s *stubSyncProducer) CommitTxn() error {
	return nil
}
func (s *stubSyncProducer) AbortTxn() error {
	return nil
}
func (s *stubSyncProducer) AddOffsetsToTxn(offsets map[string][]*sarama.PartitionOffsetMetadata, groupId string) error {
	return nil
}

func TestNewProducer_Success(t *testing.T) {
	// Save and override the builder
	orig := syncProducerBuilder
	stub := &stubSyncProducer{}
	syncProducerBuilder = func(cfg *KafkaConfig) (sarama.SyncProducer, error) {
		return stub, nil
	}
	defer func() { syncProducerBuilder = orig }()

	cfg := &KafkaConfig{Brokers: []string{"localhost:9092"}, Topic: "topic1", ConsumerGroup: "group1", Version: "2.6.0"}
	p, err := NewProducer(cfg)

	require.NoError(t, err)
	require.NotNil(t, p)
	require.Equal(t, "topic1", p.Topic)
	require.Same(t, stub, p.Producer)
}

func TestNewProducer_Error(t *testing.T) {
	orig := syncProducerBuilder
	syncProducerBuilder = func(cfg *KafkaConfig) (sarama.SyncProducer, error) {
		return nil, errors.New("boom")
	}
	defer func() { syncProducerBuilder = orig }()

	p, err := NewProducer(&KafkaConfig{Brokers: []string{"localhost:9092"}, Topic: "topic1", ConsumerGroup: "group1", Version: "2.6.0"})

	require.Error(t, err)
	require.Nil(t, p)
}
