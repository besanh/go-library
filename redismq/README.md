# RedisMQ Package

The `redismq` package implements a reliable, high-performance messaging queue system backed by Redis, wrapping the `github.com/adjust/rmq/v5` library. 

## Architecture Flow

1. **Initialization**: Creates a specialized `RMQConnection` configured to point at a Redis instance, booting up an isolated connection pool explicitly scaled for high concurrency.
2. **Client/Server Separation**:
    * **Server Mode**: Can register queues, attach consumer processes, handle incoming messages, and manage graceful shutdown signals safely.
    * **Client Mode**: Built for pushing messages into defined queues rapidly using synchronized maps ensuring queue initialization is atomic.
3. **Background Monitoring**: Automatically tracks and logs connection heartbeat and delivery errors to ensure observability over message failures.

## Usage

```go
import "github.com/besanh/go-library/redismq"
```

### 1. Connection Initialization

Construct the core connection manager.

```go
cfg := redismq.Config{
    Address: "localhost:6379",
    DB: 1,
}

rmq, err := redismq.NewRMQ(cfg)
```

### 2. Publishing Messages (Client)

Use the client instance to publish strings or raw byte payloads to a named queue.

```go
err = rmq.Client.Publish("my-queue", "task_payload")
err = rmq.Client.PublishBytes("my-queue", []byte(`{"id": 123}`))
```

### 3. Consuming Messages (Server)

Define a consumer function, bind it to a queue, and execute the server block to wait for incoming payloads.

```go
func myHandler(delivery rmq.Delivery) {
    log.Println("Received:", delivery.Payload())
    
    // Process task...

    // Acknowledge or Reject message
    delivery.Ack()
}

// Add the queue and attach 10 concurrent consumers
err := rmq.Server.AddQueue("my-queue", myHandler, 10)

// Block main thread & gracefully handle termination signals
rmq.Server.Run()
```
