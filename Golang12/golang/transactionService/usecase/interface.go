package usecase

import (
	"context"
	"source-base-go/golang/transactionService/api/payload"
	"source-base-go/golang/transactionService/model"
)

type TransactionRepository interface {
	BeginTx(ctx context.Context) (Tx model.Tx, error)
	SaveTransaction(ctx context.Context, tx model.Tx, info model.TransactionInfo) error
	CommitTx(tx model.Tx) error
	RollbackTx(tx model.Tx) error
}

type AuthService interface {
	VerifyJWT(ctx context.Context, token string) (userID string, err error)
}
type WalletClient interface {
	DebitBalance(ctx context.Context, info model.DebitInfo) (*model.WalletResult, error)
	CreditBalance(ctx context.Context, info model.CreditInfo) (*model.WalletResult, error)
}
type UseCase interface {
	TransferMoney(ctx context.Context, tranferPayload payload.TransferPayload) error
}
