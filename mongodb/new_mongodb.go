package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	URI               string
	Database          string
	ConnectionTimeout time.Duration
	SocketTimeout     time.Duration
	MaxPoolSize       uint64
	MinPoolSize       uint64
}

type Client struct {
	Client *mongo.Client
	DB     *mongo.Database
	ctx    context.Context
}

func NewClient(cfg Config) (*Client, error) {
	if cfg.ConnectionTimeout == 0 {
		cfg.ConnectionTimeout = 10 * time.Second
	}
	if cfg.SocketTimeout == 0 {
		cfg.SocketTimeout = 5 * time.Second
	}
	if cfg.MaxPoolSize == 0 {
		cfg.MaxPoolSize = 100
	}

	opts := options.Client().ApplyURI(cfg.URI)
	opts.SetConnectTimeout(cfg.ConnectionTimeout)
	opts.SetSocketTimeout(cfg.SocketTimeout)
	opts.SetMaxPoolSize(cfg.MaxPoolSize)
	if cfg.MinPoolSize > 0 {
		opts.SetMinPoolSize(cfg.MinPoolSize)
	}

	ctx := context.Background()

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("mongo connect error: %w", err)
	}

	pingCtx, cancel := context.WithTimeout(ctx, cfg.ConnectionTimeout)
	defer cancel()
	if err := client.Ping(pingCtx, nil); err != nil {
		return nil, fmt.Errorf("mongo ping error: %w", err)
	}

	return &Client{
		Client: client,
		DB:     client.Database(cfg.Database),
		ctx:    ctx,
	}, nil
}

func (c *Client) Close() error {
	return c.Client.Disconnect(c.ctx)
}

func (c *Client) Collection(name string) *mongo.Collection {
	return c.DB.Collection(name)
}
