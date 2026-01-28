package grpc_client

import (
	"crypto/tls"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

func InsecureConnection() grpc.DialOption {
	return grpc.WithTransportCredentials(insecure.NewCredentials())
}

func SecureConnection() grpc.DialOption {
	return grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{}))
}
