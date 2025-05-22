package grpcclient

import (
	"context"
	"source-base-go/golang/proto/auth"

)

type AuthClient interface {
	VerifyJWT(ctx context.Context) (*auth.TokenResponse, error)
}
type authClient struct {
	client auth.AuthServiceClient
}

func NewAuthClient(c auth.AuthServiceClient) AuthClient {
	return &authClient{client: c}
}

func (a *authClient) VerifyJWT(ctx context.Context) (*auth.TokenResponse, error) {
    req := &auth.TokenRequest{}
    res, err := a.client.VerifyJWT(ctx, req)
    if err != nil {
        return nil, err
    }
    return res, nil
}