package kafka

import "github.com/IBM/sarama"

var consumerGroupBuilder = func(cfg *KafkaConfig) (sarama.ConsumerGroup, error) {
	return NewConsumerGroupClient(cfg)
}
var syncProducerBuilder = func(cfg *KafkaConfig) (sarama.SyncProducer, error) {
	return NewSyncProducer(cfg)
}
