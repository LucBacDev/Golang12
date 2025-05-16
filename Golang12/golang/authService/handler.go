package main

import (
	"context"
	pb "wallet-app/proto"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
}

func (s *AuthServer) VerifyJWT(ctx context.Context, req *pb.TokenRequest) (*pb.TokenData, error) {
	// ⚠️ Giả lập xác thực token
	if req.Token == "valid-token" {
		return &pb.TokenData{
			UserId:   "user123",
			Role:     "user",
			IsValid:  true,
			Message:  "Token is valid",
		}, nil
	}
	return &pb.TokenData{
		IsValid: false,
		Message: "Invalid token",
	}, nil
}
