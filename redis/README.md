# Redis Package

The `redis` package provides a standardized connection wrapper around `github.com/redis/go-redis/v9`.

## Architecture Flow

1. **Configuration Generation**: Receives a `Config` struct mapping directly to `go-redis` native connection settings.
2. **Client Bootstrapping**: Initializes a base `redis.Client` connecting safely to the specified instance or cluster.
3. **Health Check Validation**: Pings the Redis instance upon startup. The initialization immediately fails if the cluster is unreachable, preventing latent connectivity issues.

## Usage

```go
import "github.com/besanh/go-library/redis"
```

### 1. Connecting to Redis

Initialize the Redis client with pool and timeout configurations.

```go
cfg := redis.Config{
    Address:      "localhost:6379",
    Password:     "secret",
    DB:           0,
    PoolSize:     100,
    DialTimeout:  5 * time.Second,
}

client, err := redis.NewRedis(cfg)
if err != nil {
    log.Fatal("Failed to connect to Redis:", err)
}

// Access the underlying go-redis client directly
client.Rdb.Set(client.Ctx, "key", "value", 0)
```
