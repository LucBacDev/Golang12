package grpcclient

import (
	"context"
	"source-base-go/golang/service/transactionService/model"
	"source-base-go/golang/proto/wallet"
)

type WalletClient interface {
	DebitBalance(ctx context.Context, info model.DebitInfo) (*wallet.DebitResponse, error)
	RefundDebit(ctx context.Context, info model.DebitInfo) (*wallet.RefundDebitResponse, error)
	CreditBalance(ctx context.Context, info model.CreditInfo) (*wallet.CreditResponse, error)
	UndoCredit(ctx context.Context, info model.CreditInfo) (*wallet.UndoCreditResponse, error)

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
func (w *walletClient) RefundDebit(ctx context.Context, info model.DebitInfo) (*wallet.RefundDebitResponse, error) {
	req := &wallet.RefundDebitRequest{
		UserId: info.UserID,
		Amount: info.Amount,
	}
	return w.client.RefundDebit(ctx, req)
}

func (w *walletClient) UndoCredit(ctx context.Context, info model.CreditInfo) (*wallet.UndoCreditResponse, error) {
	req := &wallet.UndoCreditRequest{
		UserId: info.UserID,
		Amount: info.Amount,
	}
	return w.client.UndoCredit(ctx, req)
}