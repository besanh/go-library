# MongoDB Package

The `mongodb` package provides a standardized connection wrapper around the official `go.mongodb.org/mongo-driver`, easing boilerplate configuration for timeouts and connection pooling.

## Architecture Flow

1. **Configuration Generation**: Receives a unified `Config` holding the connection URI, database intent, and optional connection/socket pool size parameters. If missing, robust fallback defaults (e.g. 10s connection timeout, 5s socket timeout) are applied automatically.
2. **Client Bootstrapping**: Initializes a base `mongo.Client` with the applied configuration.
3. **Health Check Validation**: Automatically executes a context-bound `Ping` request against the database cluster. Bootstrapping halts organically if the cluster is unreachable, preventing latent failures during request phases.

## Usage

```go
import "github.com/besanh/go-library/mongodb"
```

### 1. Connecting to the Database

Construct a `Config` and initialize the MongoDB client.

```go
cfg := mongodb.Config{
    URI:               "mongodb://username:password@localhost:27017",
    Database:          "my_database",
    ConnectionTimeout: 15 * time.Second,
    MinPoolSize:       10,
    MaxPoolSize:       200,
}

client, err := mongodb.NewClient(cfg)
if err != nil {
    log.Fatalf("Failed to connect to MongoDB: %v", err)
}

// Cleanly disconnect when the application shuts down
defer client.Close()
```

### 2. Performing Database Operations

The `Client` exposes direct access to the `mongo.Client` and `mongo.Database`, as well as a helper function to retrieve collections quickly.

```go
// Access a specific collection
usersCol := client.Collection("users")

// Execute operations using the standard mongo-driver API
result, err := usersCol.InsertOne(context.Background(), myDocument)

// Access base driver instances if needed:
// client.Client
// client.DB
```
