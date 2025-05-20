package grpcclient

import (
	"context"
	"source-base-go/golang/proto/auth"

	"google.golang.org/grpc/metadata"
)

type AuthClient interface {
	VerifyJWT(ctx context.Context, token string) (*auth.TokenResponse, error)
}
type authClient struct {
	client auth.AuthServiceClient
}

func NewAuthClient(c auth.AuthServiceClient) AuthClient {
	return &authClient{client: c}
}

func (a *authClient) VerifyJWT(ctx context.Context, token string) (*auth.TokenResponse, error) {
    md := metadata.New(map[string]string{"authorization": "Bearer " + token})
    ctxWithMeta := metadata.NewOutgoingContext(ctx, md)

    req := &auth.TokenRequest{Token: token}
    res, err := a.client.VerifyJWT(ctxWithMeta, req)
    if err != nil {
        return nil, err
    }
    return res, nil
}