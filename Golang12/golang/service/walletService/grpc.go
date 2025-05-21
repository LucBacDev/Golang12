package main

import (
	"log"
	"net"
	"source-base-go/services/walletService/infrastructure/repository"
	"source-base-go/services/walletService/usecase"

	"google.golang.org/grpc"
)

type gRPCServer  struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr}
}

func (s *gRPCServer) Run() error {
	envConfig := getConfig()

	db, err := repository.ConnectDB(envConfig.Sql)
	if err != nil {
		log.Println(err)
		return err
	}

	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	conn := grpc.NewServer()
	orderRepo := repository.NewOrderRepository(db)
	usecase.NewService(orderRepo, conn)

	return conn.Serve(lis)
}