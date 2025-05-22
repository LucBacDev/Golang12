package main

import (
	"log"
	"net"
	"source-base-go/golang/service/walletService/infrastructure/repository"
	"source-base-go/golang/service/walletService/infrastructure/repository/util"
	"source-base-go/golang/service/walletService/interceptor"
	"source-base-go/golang/service/walletService/usecase"

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
	var verifier util.Verifier = util.NewAccessTokenRepository(db)

	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	conn := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.AuthUnaryInterceptor(verifier)),

	)
	walletRepo := repository.NewWalletRepository(db)
	usecase.NewService(walletRepo, conn)

	return conn.Serve(lis)
}