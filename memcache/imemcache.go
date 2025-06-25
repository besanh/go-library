package memcache

import "time"

type IMemCache interface {
	Set(key string, value any, ttl time.Duration)
	Get(key string) any
	Del(key string)
	Close()
}

var MCache IMemCache

func (s *MemCache) Set(key string, value any, ttl time.Duration) {
	s.Cache.Set(key, value, ttl)
}

func (s *MemCache) Get(key string) any {
	val := s.Cache.Get(key)
	if val == nil {
		return nil
	}
	return val.Value()
}

func (s *MemCache) Del(key string) {
	s.Cache.Delete(key)
}

func (s *MemCache) Close() {
	s.Cache.Stop()
}
