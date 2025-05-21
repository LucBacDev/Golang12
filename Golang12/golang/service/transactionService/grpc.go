package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

type gRPCServer  struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr}
}

func (s *gRPCServer) Run() error {

	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Printf("failed to listen: %v", err)
	}
	conn := grpc.NewServer()
	return conn.Serve(lis)

}