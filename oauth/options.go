package oauth

import (
	"log"
	"time"
)

type options struct {
	jwksURL        string
	refreshTimeout time.Duration
	errorHandler   func(err error)
}

// Option is a function that configures the Authenticator
type Option func(*options)

// WithJWKSURL sets the URL to fetch the Public Keys
func WithJWKSURL(url string) Option {
	return func(o *options) {
		o.jwksURL = url
	}
}

// WithRefreshTimeout sets how often to refresh the keys in the background
func WithRefreshTimeout(timeout time.Duration) Option {
	return func(o *options) {
		o.refreshTimeout = timeout
	}
}

// WithErrorHandler allows the consumer to plug in their own logger
func WithErrorHandler(handler func(err error)) Option {
	return func(o *options) {
		o.errorHandler = handler
	}
}

func defaultOptions() *options {
	return &options{
		jwksURL:        "", // must be provided by user
		refreshTimeout: 5 * time.Minute,
		errorHandler: func(err error) {
			log.Printf("authlib jwks error: %v\n", err)
		},
	}
}
