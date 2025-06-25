package redis

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewRedis_Defaults(t *testing.T) {
	cfg := Config{}
	client, err := NewRedis(cfg)
	require.NoError(t, err)
	require.NotNil(t, client)

	op := client.rdb.Options()
	require.Equal(t, "localhost:6379", op.Addr)
	require.Equal(t, "", op.Password)
	require.Equal(t, 0, op.DB)
	require.Equal(t, 5*time.Second, op.DialTimeout)
	require.Equal(t, 3*time.Second, op.ReadTimeout)
	require.Equal(t, 3*time.Second, op.WriteTimeout)
	require.Equal(t, 10, op.PoolSize)
	require.Equal(t, 2, op.MinIdleConns)
	require.NotNil(t, client.ctx)
}

func TestNewRedis_CustomConfigs(t *testing.T) {
	custom := Config{
		Addr:         "127.0.0.1:6380",
		Password:     "secret",
		DB:           2,
		DialTimeout:  1 * time.Second,
		ReadTimeout:  500 * time.Millisecond,
		WriteTimeout: 500 * time.Millisecond,
		PoolSize:     20,
		MinIdleConns: 5,
	}
	client, err := NewRedis(custom)
	require.NoError(t, err)

	op := client.rdb.Options()
	require.Equal(t, "127.0.0.1:6380", op.Addr)
	require.Equal(t, "secret", op.Password)
	require.Equal(t, 2, op.DB)
	require.Equal(t, 1*time.Second, op.DialTimeout)
	require.Equal(t, 500*time.Millisecond, op.ReadTimeout)
	require.Equal(t, 500*time.Millisecond, op.WriteTimeout)
	require.Equal(t, 20, op.PoolSize)
	require.Equal(t, 5, op.MinIdleConns)
}
