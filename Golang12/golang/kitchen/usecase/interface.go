package usecase

import (
	"context"
	"source-base-go/common/genproto/orders"
	"source-base-go/services/kitchen/api/payload"
)

type TransactionRepository interface {
	TransferMoney(ctx context.Context, fromWalletID, toWalletID string, amount float64) error
}

type UseCase interface {
	TransferMoney(ctx context.Context, fromWalletID, toWalletID string, amount float64) error
}
