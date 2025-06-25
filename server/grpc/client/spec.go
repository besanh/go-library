package client

import "google.golang.org/grpc"

type GrpcConn interface {
	Dial() *grpc.ClientConn
	Close()
}
