package usecase

import (
	"context"
	"source-base-go/golang/proto/wallet"
	"source-base-go/golang/service/transactionService/api/payload"
	"source-base-go/golang/service/transactionService/model"
)

type TransactionRepository interface {
	SaveTransaction(ctx context.Context, info *model.Transaction) error
	SaveLog(ctx context.Context, log model.TransactionLog) error

}
type AuthService interface {
	VerifyJWT(ctx context.Context) (userID string, err error)
}
type WalletClient interface {
	GetUserByAccountNumber(ctx context.Context, accountNumber string) (*wallet.GetUserByAccountNumberResponse, error)
	DebitBalance(ctx context.Context, info model.DebitInfo) (*model.WalletResult, error)
	CreditBalance(ctx context.Context, info model.CreditInfo) (*model.WalletResult, error)
	RefundDebit(ctx context.Context, info model.DebitInfo) (*model.WalletResult, error)
	UndoCredit(ctx context.Context, info model.CreditInfo) (*model.WalletResult, error)
}
type UseCase interface {
	GetReceiverInfo(ctx context.Context, accountNumber string, token string) (*wallet.GetUserByAccountNumberResponse, error)
	TransferMoney(ctx context.Context, tranferPayload *payload.TransferPayload) (map[string]interface{}, error)
}
