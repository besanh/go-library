# Memcache Package

The `memcache` package provides a standardized, thread-safe, in-memory caching wrapper built around the `github.com/jellydator/ttlcache/v3` library. 

## Architecture Flow

1. **Initialization**: The package initializes an in-memory TTL cache instance via `NewMemCache`. By default, this cache is configured with a 30-minute Time-To-Live (TTL) for all items.
2. **Data Storage**: Stores mapping of `string` keys to `any` values. The underlying `ttlcache` handles background expiration and thread-safe access.

## Usage

```go
import "github.com/besanh/go-library/memcache"
```

### 1. Initializing the Cache

Create a new instance of the memory cache.

```go
cache := memcache.NewMemCache()
```

### 2. Standard Operations

The `IMemCache` interface supports all standard `ttlcache.Cache` operations dynamically.

```go
// Set a value
cache.Set("user_123", myUserObj, ttlcache.DefaultTTL)

// Get a value
if item := cache.Get("user_123"); item != nil {
    user := item.Value().(UserType)
    // ...
}

// Delete a value
cache.Delete("user_123")
```
