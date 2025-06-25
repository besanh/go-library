package kafka

import (
	"github.com/IBM/sarama"
)

type Producer struct {
	Topic    string
	Producer sarama.SyncProducer
}

func NewProducer(cfg *KafkaConfig) (*Producer, error) {
	producer, err := syncProducerBuilder(cfg)
	if err != nil {
		return nil, err
	}

	return &Producer{
		Topic:    cfg.Topic,
		Producer: producer,
	}, nil
}
