package spicedb

import (
	"time"

	"github.com/authzed/grpcutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// WithInsecure configures the client to use an insecure bearer token.
// This is primarily used for local development environments.
func WithInsecure() Option {
	return func(c *clientConfig) {
		c.dialOpts = append(c.dialOpts,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpcutil.WithInsecureBearerToken(c.preSharedKey),
		)
	}
}

// WithSystemTLS configures the client to use the host machine's root certificates for TLS.
// This is recommended for production environments.
func WithSystemTLS() Option {
	return func(c *clientConfig) {
		c.dialOpts = append(c.dialOpts, grpcutil.WithBearerToken(c.preSharedKey))
		otps, err := grpcutil.WithSystemCerts(grpcutil.VerifyCA)
		if err != nil {
			panic(err)
		}
		c.dialOpts = append(c.dialOpts, otps)
	}
}

// WithTimeout allows setting a fallback gRPC dial option (currently uses WithBlock for demonstration).
// Note: In typical usage, timeouts should be managed via context.Context.
func WithTimeout(d time.Duration) Option {
	return func(c *clientConfig) {
		c.dialOpts = append(c.dialOpts, grpc.WithBlock())
	}
}
