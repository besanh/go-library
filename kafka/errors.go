package kafka

import "errors"

var (
	ErrNoBrokers = errors.New("no Kafka brokers configured")
	ErrNoTopic   = errors.New("no Kafka topic configured")
	ErrNoGroup   = errors.New("no Kafka consumer group configured")
)
