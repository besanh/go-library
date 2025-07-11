package memcache

import (
	"time"

	"github.com/jellydator/ttlcache/v3"
)

type MemCache struct {
	*ttlcache.Cache[string, any]
}

var memcacheClient = func() IMemCache {
	return NewMemCache()
}

func NewMemCache() IMemCache {
	service := &MemCache{}
	cache := ttlcache.New(
		ttlcache.WithTTL[string, any](30 * time.Minute),
	)

	service.Cache = cache
	return service
}
