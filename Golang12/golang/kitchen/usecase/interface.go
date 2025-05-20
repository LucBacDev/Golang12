package usecase

import (
	"context"
	"source-base-go/golang/proto/wallet"
	"source-base-go/golang/kitchen/api/payload"
)

// type TransactionRepository interface {
//     UpdateBalance(ctx context.Context, walletID string, newBalance float64) error
// }
type AuthService interface {
	VerifyJWT(ctx context.Context, token string) (userID string, err error)
}
type WalletClient interface {
	DebitBalance(ctx context.Context, req *wallet.DebitRequest) (*wallet.WalletResponse, error)
	CreditBalance(ctx context.Context, userID string, amount float64) (*wallet.WalletResponse, error)
}

type UseCase interface {
	TransferMoney(ctx context.Context, tranferPayload payload.TransferPayload) error
}
