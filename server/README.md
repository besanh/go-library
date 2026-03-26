# Server Package

The `server` package provides standardized bootstrapping mechanisms for spinning up HTTP and gRPC servers or clients within the ecosystem.

## Sub-Packages

### 1. HTTP Server (`server/http`)
Wraps the `github.com/gin-gonic/gin` framework to provide an opinionated HTTP routing foundation.
* **Environment-Aware**: Automatically configures Gin’s operational mode (`DebugMode` vs `ReleaseMode`) based on the environment configuration.
* **Unified Logging**: Redirects Gin's default logging to standard output and secondary streams via `io.MultiWriter`, ensuring standard request logging respects unified log collection policies.

```go
import "github.com/besanh/go-library/server/http"

cfg := http.Config{
    Environment: "release",
}
httpServer := http.New(cfg)
// Use httpServer router to register endpoints...
```

### 2. gRPC Server & Client (`server/grpc`)
Standardizes the setup and teardown for internal gRPC communications. 
* **`grpc_client`**: Interfaces and utilities for dialing and maintaining stubs communicating with other internal microservices.
* **`grpc_server`**: Framework for attaching RPC handlers and serving them over defined ports safely.
