package memcache

import (
	"testing"
	"time"

	"github.com/jellydator/ttlcache/v3"
	"github.com/stretchr/testify/require"
)

func TestNewMemCache_SetGet(t *testing.T) {
	mc := NewMemCache()
	m, ok := mc.(*MemCache)
	require.True(t, ok, "expected *MemCache, got %T", mc)

	m.Set("foo", "bar", 0)
	val, ok := m.Get("foo").(string)
	require.True(t, ok, "expected key foo is exists")
	require.Equal(t, "bar", val)
}

func TestNewMemCache_TTLExpiration(t *testing.T) {
	mc := &MemCache{}
	cache := ttlcache.New(
		ttlcache.WithTTL[string, any](50 * time.Millisecond),
	)
	mc.Cache = cache

	mc.Set("tmp", 123, 10*time.Millisecond)
	v, ok := mc.Get("tmp").(int)
	require.True(t, ok)
	require.Equal(t, 123, v)

	// Wait for TTL to expire
	time.Sleep(100 * time.Millisecond)
	_, ok = mc.Get("tmp").(int)
	require.False(t, ok, "expected key 'tmp' to expire after TTL")
}

func TestMemCache_CapacityEviction(t *testing.T) {
	// Create a cache with capacity = 1 to force eviction
	service := &MemCache{}
	cache := ttlcache.New(
		ttlcache.WithTTL[string, any](time.Minute),
		ttlcache.WithCapacity[string, any](1),
	)
	service.Cache = cache

	// Insert two items, second insert should evict first
	expiration := 50 * time.Millisecond
	service.Set("a", "one", expiration)
	service.Set("b", "two", expiration)

	_, okA := service.Get("a").(string)
	require.False(t, okA, "expected earliest key 'a' to be evicted when capacity reached")
	vB, okB := service.Get("b").(string)
	require.True(t, okB)
	require.Equal(t, "two", vB)
}
