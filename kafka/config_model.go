package kafka

import (
	"crypto/tls"
	"fmt"
	"time"
)

type KafkaConfig struct {
	Brokers        []string
	Topic          string
	ConsumerGroup  string
	Version        string
	TLS            *tls.Config
	SASLUsername   string
	SASLPassword   string
	RetryMax       int
	RetryBackoff   time.Duration
	ConsumerOffset string
	BufferSize     int
}

func (c *KafkaConfig) Validate() error {
	if len(c.Brokers) == 0 {
		return fmt.Errorf("brokers is required")
	}

	if len(c.Topic) == 0 {
		return fmt.Errorf("topic is required")
	}

	if len(c.ConsumerGroup) == 0 {
		return fmt.Errorf("consumer group is required")
	}

	if len(c.Version) == 0 {
		return fmt.Errorf("version is required")
	}

	return nil
}
