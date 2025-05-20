package grpcclient

import (
	"context"
	"source-base-go/golang/proto/wallet"
)

type WalletClient interface {
	DebitBalance(ctx context.Context, req *wallet.DebitRequest) (*wallet.DebitResponse, error)
	CreditBalance(ctx context.Context, userID string, amount float64) (*wallet.CreditResponse, error)
}

type walletClient struct {
	client wallet.WalletServiceClient
}

func NewWalletClient(c wallet.WalletServiceClient) WalletClient {
	return &walletClient{client: c}
}

func (w *walletClient) DebitBalance(ctx context.Context, req *wallet.DebitRequest) (*wallet.DebitResponse, error) {
	return w.client.DebitBalance(ctx, req)
}

func (w *walletClient) CreditBalance(ctx context.Context, userID string, amount float64) (*wallet.CreditResponse, error) {
	return w.client.CreditBalance(ctx, &wallet.CreditRequest{
		UserId: userID,
		Amount: amount,
	})
}