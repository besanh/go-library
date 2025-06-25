package kafka

import (
	"context"

	"github.com/IBM/sarama"
)

type IProducer interface {
	Send(ctx context.Context, key, value []byte, headers map[string]string) (partition int32, offset int64, err error)
	Close() error
}

func (p *Producer) Send(ctx context.Context, key, value []byte, headers map[string]string) (partition int32, offset int64, err error) {
	msg := &sarama.ProducerMessage{
		Topic:   p.Topic,
		Key:     sarama.ByteEncoder(key),
		Value:   sarama.ByteEncoder(value),
		Headers: make([]sarama.RecordHeader, 0, len(headers)),
	}

	for k, v := range headers {
		msg.Headers = append(msg.Headers, sarama.RecordHeader{
			Key:   []byte(k),
			Value: []byte(v),
		})
	}

	return p.Producer.SendMessage(msg)
}

func (p *Producer) Close() error {
	return p.Producer.Close()
}
