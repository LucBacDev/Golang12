package grpcclient

import (
    "context"
    authpb "transaction/proto/auth"
)

type AuthClient struct {
    Client authpb.AuthServiceClient
}

func (a *AuthClient) VerifyJWT(ctx context.Context, token string) (*authpb.TokenData, error) {
    return a.Client.VerifyJWT(ctx, &authpb.TokenRequest{Token: token})
}


// package client

// import (
// 	"context"
// 	"log"
// 	"time"

// 	pb "wallet-app/proto"

// 	"google.golang.org/grpc"
// )

// func VerifyToken(token string) (*pb.TokenData, error) {
// 	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer conn.Close()

// 	client := pb.NewAuthServiceClient(conn)
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()

// 	resp, err := client.VerifyJWT(ctx, &pb.TokenRequest{Token: token})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return resp, nil
// }