package kafka

import "errors"

var (
	ErrNoBrokers error = errors.New("no Kafka brokers configured")
	ErrNoTopic   error = errors.New("no Kafka topic configured")
	ErrNoGroup   error = errors.New("no Kafka consumer group configured")
)
