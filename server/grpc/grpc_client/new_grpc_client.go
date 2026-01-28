package grpc_client

import (
	"google.golang.org/grpc"
)

type grpcConn struct {
	conn *grpc.ClientConn
}

func NewConn(address string, opts ...grpc.DialOption) IGrpcClient {
	if len(opts) == 0 {
		opts = append(opts, InsecureConnection())
	}

	conn, err := grpc.NewClient(
		address,
		opts...,
	)
	if err != nil {
		panic(err)
	}

	return &grpcConn{
		conn: conn,
	}
}

func (g *grpcConn) Close() {
	defer g.conn.Close()
}

func (g *grpcConn) Dial() *grpc.ClientConn {
	return g.conn
}
