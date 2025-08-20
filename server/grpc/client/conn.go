package client

import (
	"google.golang.org/grpc"
)

type grpcConn struct {
	conn *grpc.ClientConn
}

func NewConn(address string, options ...grpc.DialOption) GrpcConn {
	conn, err := grpc.NewClient(
		address,
		options...,
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
