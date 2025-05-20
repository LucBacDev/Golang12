package grpcclient

import (
	"context"
	"source-base-go/golang/proto/user"
)

type UserClient interface {
	GetUserByAccountNumber(ctx context.Context, accountNumber string) (*user.GetUserByAccountNumberResponse, error)
}

type userClient struct {
	client user.UserServiceClient
}

func NewUserClient(c user.UserServiceClient) UserClient {
	return &userClient{client: c}
}

func (u *userClient) GetUserByAccountNumber(ctx context.Context, accountNumber string) (*user.GetUserByAccountNumberResponse, error) {
	req := &user.GetUserByAccountNumberRequest{
		AccountNumber: accountNumber,
	}
	res, err := u.client.GetUserByAccountNumber(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}