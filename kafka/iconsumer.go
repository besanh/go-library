package kafka

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
)

// IMessageHandler is a callback for consumed messages
type IMessageHandler interface {
	HandleMessage(msg *sarama.ConsumerMessage) error
}

// IConsumerGroup defines the consumer group interface
type IConsumerGroup interface {
	Start(ctx context.Context) error
	Close() error
}

// Start consuming until context cancellation or SIGINT/SIGTERM
func (c *ConsumerGroup) Start(ctx context.Context) error {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(ctx)
	go func() { <-sigs; cancel() }()

	h := &cgHandler{handler: c.handler}
	for {
		if err := c.group.Consume(ctx, []string{c.topic}, h); err != nil {
			log.Printf("kafka: consume error: %v", err)
		}
		if ctx.Err() != nil {
			return ctx.Err()
		}
	}
}

// Close stops the consumer group
func (c *ConsumerGroup) Close() error {
	return c.group.Close()
}

// adapter for sarama.ConsumerGroupHandler
type cgHandler struct{ handler IMessageHandler }

func (h *cgHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (h *cgHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h *cgHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		if err := h.handler.HandleMessage(msg); err != nil {
			log.Printf("kafka: handler error: %v", err)
		}
		sess.MarkMessage(msg, "")
	}
	return nil
}
