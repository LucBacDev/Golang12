package grpcclient

import (
    "context"
    walletpb "transaction/proto/wallet"
)

type WalletClient struct {
    Client walletpb.WalletServiceClient
}

func (w *WalletClient) Debit(ctx context.Context, userID string, amount int64) (*walletpb.WalletResponse, error) {
    return w.Client.DebitBalance(ctx, &walletpb.DebitRequest{
        UserId: userID,
        Amount: amount,
    })
}

func (w *WalletClient) Credit(ctx context.Context, userID string, amount int64) (*walletpb.WalletResponse, error) {
    return w.Client.CreditBalance(ctx, &walletpb.CreditRequest{
        UserId: userID,
        Amount: amount,
    })
}


// package client

// import (
// 	"context"
// 	"time"

// 	pb "wallet-app/proto"

// 	"google.golang.org/grpc"
// )

// func Debit(userId string, amount int64) (*pb.WalletResponse, error) {
// 	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer conn.Close()

// 	client := pb.NewWalletServiceClient(conn)
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()

// 	return client.DebitBalance(ctx, &pb.DebitRequest{UserId: userId, Amount: amount})
// }

// func Credit(userId string, amount int64) (*pb.WalletResponse, error) {
// 	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer conn.Close()

// 	client := pb.NewWalletServiceClient(conn)
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()

// 	return client.CreditBalance(ctx, &pb.CreditRequest{UserId: userId, Amount: amount})
// }
