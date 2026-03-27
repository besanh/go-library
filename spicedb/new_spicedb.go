// Package spicedb provides a lightweight wrapper around the Authzed SpiceDB client,
// offering a functional options pattern for easy configuration and initialization.
package spicedb

import (
	"fmt"

	"github.com/authzed/authzed-go/v1"
	"google.golang.org/grpc"
)

type clientConfig struct {
	endpoint     string
	preSharedKey string
	dialOpts     []grpc.DialOption
}

// SpiceClient is a wrapper around the authzed.Client.
type SpiceClient struct {
	Client *authzed.Client
}

type Option func(*clientConfig)

// NewClient initializes a new SpiceClient with the given endpoint and pre-shared key (token).
// It accepts optional functional options to configure the client, such as TLS or timeout settings.
func NewClient(endpoint, token string, opts ...Option) (*SpiceClient, error) {
	config := &clientConfig{
		endpoint:     endpoint,
		preSharedKey: token,
		dialOpts:     make([]grpc.DialOption, 0),
	}

	for _, opt := range opts {
		opt(config)
	}

	// Fallback: If they didn't specify TLS or Insecure, default to Insecure for safety in dev
	if len(config.dialOpts) == 0 {
		WithInsecure()(config)
	}

	// Initialize the actual client with the accumulated gRPC dial options
	client, err := authzed.NewClient(
		config.endpoint,
		config.dialOpts...,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to spicedb: %w", err)
	}

	return &SpiceClient{Client: client}, nil
}
