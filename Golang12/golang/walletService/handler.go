package main

import (
	"context"
	"log"

	pb "wallet-app/proto"
)

type WalletServer struct {
	pb.UnimplementedWalletServiceServer
}

func (s *WalletServer) DebitBalance(ctx context.Context, req *pb.DebitRequest) (*pb.WalletResponse, error) {
	log.Printf("Debiting %d from user %s\n", req.Amount, req.UserId)
	// Giả lập luôn thành công
	return &pb.WalletResponse{Success: true, Message: "Debit successful"}, nil
}

func (s *WalletServer) CreditBalance(ctx context.Context, req *pb.CreditRequest) (*pb.WalletResponse, error) {
	log.Printf("Crediting %d to user %s\n", req.Amount, req.UserId)
	return &pb.WalletResponse{Success: true, Message: "Credit successful"}, nil
}
