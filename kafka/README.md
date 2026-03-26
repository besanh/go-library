# Kafka Package

The `kafka` package provides standardized wrappers around the `github.com/IBM/sarama` library, offering simplified interfaces for configuring and using Kafka Publishers and Consumer Groups within the system.

## Architecture Flow

1. **Configuration**: A unified `KafkaConfig` object defines the brokers, topic, consumer group, TLS/SASL settings, and retry mechanisms.
2. **Producer**: The `NewProducer` method constructs a synchronous Sarama producer, automatically applying TLS/SASL authentication if provided in the configuration.
3. **Consumer**: The `NewConsumerGroup` method constructs a Sarama consumer group linked to an `IMessageHandler`. When messages are polled from the specified topic, they are piped asynchronously to the provided message handler.

## Usage

```go
import "github.com/besanh/go-library/kafka"
```

### 1. Configuration

Initialize a `KafkaConfig` struct with your broker requirements.

```go
cfg := &kafka.KafkaConfig{
    Brokers:        []string{"localhost:9092"},
    Topic:          "my-events",
    ConsumerGroup:  "my-service-group",
    Version:        "2.8.1",
    BufferSize:     100,
    ConsumerOffset: "newest", 
    // TLS and SASL properties can also be configured here
}
```

### 2. Setting up a Producer

Create a new producer using the configuration to send messages synchronously.

```go
producer, err := kafka.NewProducer(cfg)
if err != nil {
    // handle error
}

// Ensure proper closure when shutting down
defer producer.Producer.Close()

msg := &sarama.ProducerMessage{
    Topic: cfg.Topic,
    Value: sarama.StringEncoder("hello world"),
}

partition, offset, err := producer.Producer.SendMessage(msg)
```

### 3. Setting up a Consumer Group

Implement the `IMessageHandler` interface to process incoming messages, and create a consumer group.

```go
// Consumer logic
consumer, err := kafka.NewConsumerGroup(cfg, myMessageHandler)
if err != nil {
    // handle error
}

// Start consuming from the topic
// Usually run within a dedicated goroutine matching the application's lifecycle context
```
