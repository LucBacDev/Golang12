package usecase

import (
	"context"
	"source-base-go/golang/proto/auth"
)


type UseCase interface {
	VerifyJWT(ctx context.Context, req *auth.TokenRequest) (*auth.TokenResponse, error)
}