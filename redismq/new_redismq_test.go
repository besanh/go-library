package redismq

import (
	"context"
	"runtime"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/stretchr/testify/require"
)

func TestNewRMQ_SetsUpConnection(t *testing.T) {
	// Start an in-memory Redis server
	s, err := miniredis.Run()
	require.NoError(t, err)
	defer s.Close()

	cfg := Config{
		Address:         s.Addr(),
		Username:        "",
		Password:        "",
		DB:              0,
		PoolSize:        32,
		DialTimeout:     60 * time.Millisecond,
		PoolTimeout:     60 * time.Millisecond,
		ReadTimeout:     60 * time.Millisecond,
		WriteTimeout:    60 * time.Millisecond,
		ConnMaxIdleTime: 60 * time.Millisecond,
	}

	rmqConn := NewRMQ(cfg)
	require.NotNil(t, rmqConn)

	// Config must be preserved
	require.Equal(t, cfg, rmqConn.Config)

	// Redis client should respond to Ping
	pong, err := rmqConn.RedisClient.Ping(context.Background()).Result()
	require.NoError(t, err)
	require.Equal(t, "PONG", pong)

	// Verify Redis client options
	rOpts := rmqConn.RedisClient.Options()
	require.Equal(t, s.Addr(), rOpts.Addr)
	require.Equal(t, cfg.Username, rOpts.Username)
	require.Equal(t, cfg.Password, rOpts.Password)
	require.Equal(t, cfg.DB, rOpts.DB)
	require.Equal(t, cfg.PoolSize, rOpts.PoolSize)
	require.Equal(t, cfg.DialTimeout, rOpts.DialTimeout)
	require.Equal(t, cfg.PoolTimeout, rOpts.PoolTimeout)
	require.Equal(t, cfg.ReadTimeout, rOpts.ReadTimeout)
	require.Equal(t, cfg.WriteTimeout, rOpts.WriteTimeout)
	require.Equal(t, cfg.ConnMaxIdleTime, rOpts.ConnMaxIdleTime)

	// Default PoolSize should be 4 * NumCPU
	expectedPool := runtime.NumCPU() * 4
	require.Equal(t, expectedPool, rOpts.PoolSize) // RMQ connection should be a valid, non-zero connection object
	require.NotZero(t, rmqConn.Conn)

	// Server and Client share the same underlying connection
	require.Same(t, rmqConn.Conn, rmqConn.Server.conn)
	require.Same(t, rmqConn.Conn, rmqConn.Client.conn)

	// Top-level Queues map should be nil initially
	require.Nil(t, rmqConn.Queues)

	// Server and Client queues maps should be initialized and empty
	require.NotNil(t, rmqConn.Server.Queues)
	require.Empty(t, rmqConn.Server.Queues)
	require.NotNil(t, rmqConn.Client.Queues)
	require.Empty(t, rmqConn.Client.Queues)
}
