package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Address         string
	Username        string
	Password        string
	DB              int
	PoolSize        int
	PoolTimeout     time.Duration
	DialTimeout     time.Duration
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	ConnMaxIdleTime time.Duration
}

type Client struct {
	Rdb      *redis.Client
	Ctx      context.Context
	RedisNil error
}

func NewRedis(cfg Config) (*Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:            cfg.Address,
		Username:        cfg.Username,
		Password:        cfg.Password,
		DB:              cfg.DB,
		PoolSize:        cfg.PoolSize,
		PoolTimeout:     cfg.PoolTimeout,
		DialTimeout:     cfg.DialTimeout,
		ReadTimeout:     cfg.ReadTimeout,
		WriteTimeout:    cfg.WriteTimeout,
		ConnMaxIdleTime: cfg.ConnMaxIdleTime,
	})

	ctx := context.Background()

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &Client{
		Rdb:      rdb,
		Ctx:      ctx,
		RedisNil: redis.Nil,
	}, nil
}
