# SpiceDB Package

The `spicedb` package provides a lightweight, idiomatic Go wrapper around the [Authzed SpiceDB](https://authzed.com/docs/spicedb) gRPC client. It simplifies client initialization using the functional options pattern, making it easy to configure security and timeout settings across different environments.

## Features

- **Simplified Initialization**: Quick setup with `NewClient`.
- **Functional Options**: Configure TLS, insecure mode, and timeouts easily.
- **Direct Access**: Exposes the underlying `authzed.Client` for full API access.

## Installation

```bash
go get github.com/besanh/go-library/spicedb
```

## Usage

### 1. Basic Initialization (Insecure)

Great for local development with a local SpiceDB instance.

```go
import "github.com/besanh/go-library/spicedb"

// Initialize client
client, err := spicedb.NewClient("localhost:50051", "some-preshared-key", spicedb.WithInsecure())
if err != nil {
    log.Fatalf("failed to create spicedb client: %v", err)
}

// Access the underlying authzed client
// client.Client.CheckPermission(...)
```

### 2. Production Initialization (System TLS)

Uses the host's root certificates for secure communication.

```go
client, err := spicedb.NewClient("spicedb.production:50051", "prod-token", spicedb.WithSystemTLS())
if err != nil {
    log.Fatalf("failed to create spicedb client: %v", err)
}
```

### 3. Checking Permissions

Once initialized, you can use the `Client` field to interact with SpiceDB.

```go
import (
    v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
    "context"
)

resp, err := client.Client.CheckPermission(context.Background(), &v1.CheckPermissionRequest{
    Resource: &v1.ObjectReference{
        ObjectType: "document",
        ObjectId:   "123",
    },
    Permission: "read",
    Subject: &v1.SubjectReference{
        Object: &v1.ObjectReference{
            ObjectType: "user",
            ObjectId:   "me",
        },
    },
})
```

## Available Options

| Option | Description |
| :--- | :--- |
| `WithInsecure()` | Uses insecure bearer token (ideal for local development). |
| `WithSystemTLS()` | Uses host machine root certificates for secure connection. |
| `WithTimeout(d)` | Adds a generic gRPC dial option (e.g., `WithBlock`). |
