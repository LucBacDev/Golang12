package usecase

import (
	"context"
	"source-base-go/golang/service/authService/interceptor"
	"source-base-go/golang/proto/auth"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	auth.UnimplementedAuthServiceServer
}

func NewService(grpc *grpc.Server) *Service{
	s := &Service{}
	auth.RegisterAuthServiceServer(grpc, s)
	return s
}

func (s *Service) VerifyJWT(ctx context.Context, req *auth.TokenRequest) (*auth.TokenResponse, error) {
    userId, ok := interceptor.UserIDFromContext(ctx)
    if !ok {
        return &auth. TokenResponse{IsValid: false}, status.Error(codes.Unauthenticated, "user_id not found in context")
    }

    return &auth.TokenResponse{
        IsValid: true,
        UserId:  strconv.Itoa(userId), 
    }, nil
}
