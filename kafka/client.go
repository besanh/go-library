package kafka

import (
	"fmt"

	"github.com/IBM/sarama"
)

func buildSaramaConfig(c *KafkaConfig) (*sarama.Config, error) {
	if err := c.Validate(); err != nil {
		return nil, err
	}
	ver, err := sarama.ParseKafkaVersion(c.Version)
	if err != nil {
		return nil, err
	}
	sc := sarama.NewConfig()
	sc.Version = ver
	// Producer settings
	sc.Producer.Retry.Max = c.RetryMax
	sc.Producer.Retry.Backoff = c.RetryBackoff
	sc.Producer.Return.Successes = true
	// Consumer settings
	sc.Consumer.Offsets.Initial = sarama.OffsetOldest
	if c.ConsumerOffset == "newest" {
		sc.Consumer.Offsets.Initial = sarama.OffsetNewest
	}
	// TLS/SASL
	if c.TLS != nil {
		sc.Net.TLS.Enable = true
		sc.Net.TLS.Config = c.TLS
	}
	if c.SASLUsername != "" {
		sc.Net.SASL.Enable = true
		sc.Net.SASL.User = c.SASLUsername
		sc.Net.SASL.Password = c.SASLPassword
	}
	return sc, nil
}
func NewSyncProducer(c *KafkaConfig) (sarama.SyncProducer, error) {
	sc, err := buildSaramaConfig(c)
	if err != nil {
		return nil, err
	}
	prod, err := sarama.NewSyncProducer(c.Brokers, sc)
	if err != nil {
		return nil, fmt.Errorf("kafka: NewSyncProducer failed: %w", err)
	}
	return prod, nil
}

// NewConsumerGroupClient returns a sarama.ConsumerGroup
func NewConsumerGroupClient(c *KafkaConfig) (sarama.ConsumerGroup, error) {
	sc, err := buildSaramaConfig(c)
	if err != nil {
		return nil, err
	}
	cg, err := sarama.NewConsumerGroup(c.Brokers, c.ConsumerGroup, sc)
	if err != nil {
		return nil, fmt.Errorf("kafka: NewConsumerGroupClient failed: %w", err)
	}
	return cg, nil
}
