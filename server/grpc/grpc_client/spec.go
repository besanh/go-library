package grpc_client

import "google.golang.org/grpc"

type IGrpcClient interface {
	Dial() *grpc.ClientConn
	Close()
}
