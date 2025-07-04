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

	op := client.Rdb.Options()
	require.Equal(t, "localhost:6379", op.Addr)
	require.Equal(t, "", op.Username)
	require.Equal(t, "", op.Password)
	require.Equal(t, 0, op.DB)
	require.Equal(t, 5*time.Second, op.DialTimeout)
	require.Equal(t, 3*time.Second, op.ReadTimeout)
	require.Equal(t, 3*time.Second, op.WriteTimeout)
	require.Equal(t, 80, op.PoolSize)
	require.Equal(t, 4*time.Second, op.PoolTimeout)
	require.Equal(t, 0, op.MaxActiveConns)
	require.Equal(t, 0, op.MinIdleConns)
	require.Equal(t, 0, op.MaxIdleConns)
	require.NotNil(t, client.Ctx)
}
