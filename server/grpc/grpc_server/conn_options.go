package grpc_server

import (
	"crypto/tls"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func Secure(cfg *tls.Config) grpc.ServerOption {
	return grpc.Creds(credentials.NewTLS(cfg))
}
