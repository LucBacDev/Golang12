package grpcclient

import (
	"context"
	"source-base-go/golang/transactionService/model"
	"source-base-go/golang/proto/wallet"
)

type WalletClient interface {
	DebitBalance(ctx context.Context, info model.DebitInfo) (*wallet.DebitResponse, error)
	CreditBalance(ctx context.Context, info model.CreditInfo) (*wallet.CreditResponse, error)
}

type walletClient struct {
	client wallet.WalletServiceClient
}

func NewWalletClient(c wallet.WalletServiceClient) WalletClient {
	return &walletClient{client: c}
}

func (w *walletClient) DebitBalance(ctx context.Context, info model.DebitInfo) (*wallet.DebitResponse, error) {
	req := &wallet.DebitRequest{
		UserId: info.UserID,
		Amount: info.Amount,
	}
	return w.client.DebitBalance(ctx, req)
}

func (w *walletClient) CreditBalance(ctx context.Context, info model.CreditInfo) (*wallet.CreditResponse, error) {
	req := &wallet.CreditRequest{
		UserId: info.UserID,
		Amount: info.Amount,
	}
	return w.client.CreditBalance(ctx, req)
}
