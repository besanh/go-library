# Nats Jetstream Package

The `nats_jetstream` package offers a convenient wrapper for establishing connections to NATS servers and instantiating JetStream clients out of the box.

## Architecture Flow

1. **Configuration**: Accepts a simple configuration including the basic `Host` URL of the NATS cluster.
2. **Client Instantiation**: It initializes two main clients internally: a core `*nats.Conn` and a `*jetstream.JetStream` client, exposing them for immediate use. 
3. **Internal Logging**: Includes an isolated, pre-configured `logger` specifically tailored with the "natsjetstream" service name for accurate tracking of streaming issues.

## Usage

```go
import "github.com/besanh/go-library/nats_jetstream"
```

### 1. Initializing the Client

Create the NATS JetStream client by passing the host configuration.

```go
config := nats_jetstream.Config{
    Host: "nats://127.0.0.1:4222",
}

jsClient := nats_jetstream.NewNatsJetstream(config)
```

### 2. Accessing Raw Clients

The wrapper exposes the raw underlying clients via the struct or interface, permitting access to standard NATS / JetStream functionality.

```go
// Access to Core NATS
// jsClient.NC

// Access to JetStream
// jsClient.Client
```
