package kafka

import (
	"github.com/IBM/sarama"
)

// ConsumerGroup implements IConsumerGroup
type ConsumerGroup struct {
	topic   string
	group   sarama.ConsumerGroup
	handler IMessageHandler
	msgChan <-chan []byte
}

// NewConsumerGroup constructs a ConsumerGroup
func NewConsumerGroup(cfg *KafkaConfig, handler IMessageHandler) (IConsumerGroup, error) {
	cg, err := NewConsumerGroupClient(cfg)
	if err != nil {
		return nil, err
	}
	return &ConsumerGroup{
		topic:   cfg.Topic,
		group:   cg,
		handler: handler,
		msgChan: make(chan []byte, cfg.BufferSize),
	}, nil
}
