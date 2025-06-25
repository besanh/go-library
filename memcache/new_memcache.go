package memcache

import (
	"context"
	"time"

	"github.com/besanh/go-library/logger"
	"github.com/besanh/go-library/util"
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
	lg := logger.New("memcache")
	utl := util.NewUtil()
	cache.OnInsertion(func(ctx context.Context, item *ttlcache.Item[string, any]) {
		lg.Info("memcache: inserted %s, expires at %s", logger.Field{
			Key:   "key",
			Value: item.Key(),
		}, logger.Field{
			Key:   "expires_at",
			Value: item.ExpiresAt(),
		})
	})
	cache.OnEviction(func(ctx context.Context, reason ttlcache.EvictionReason, item *ttlcache.Item[string, any]) {
		if reason == ttlcache.EvictionReasonCapacityReached {
			val, _ := utl.ParseAnyToString(item.Value())
			lg.Info("memcache: removed %s, value: %v", logger.Field{
				Key:   "key",
				Value: item.Key(),
			}, logger.Field{
				Key:   "value",
				Value: val,
			})
		}
	})
	service.Cache = cache
	return service
}
