package grpc_server

import "google.golang.org/grpc"

func NewGrpcServer(opts ...grpc.ServerOption) *grpc.Server {
	return grpc.NewServer(opts...)
}
