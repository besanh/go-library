package oauth

import (
	"fmt"

	"github.com/MicahParks/keyfunc/v2"
)

type Authenticator struct {
	jwks *keyfunc.JWKS
}

func NewAuthenticator(opts ...Option) (IOauth, error) {
	// Apply defaults
	config := defaultOptions()
	for _, opt := range opts {
		opt(config)
	}

	if config.jwksURL == "" {
		return nil, fmt.Errorf("JWKS URL is required")
	}

	// Configure the keyfunc library
	kfOpts := keyfunc.Options{
		RefreshTimeout:      config.refreshTimeout,
		RefreshErrorHandler: config.errorHandler,
	}

	// Fetch keys
	jwks, err := keyfunc.Get(config.jwksURL, kfOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch JWKS: %w", err)
	}

	return &Authenticator{
		jwks: jwks,
	}, nil
}
