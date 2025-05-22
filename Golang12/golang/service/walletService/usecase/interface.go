package usecase

import (
	"context"
	"source-base-go/golang/proto/wallet"
	"source-base-go/golang/service/walletService/entity"
)

type walletRepository interface {
	GetUserByAccountNumber(ctx context.Context, accountNumber string) (*entity.UserResponse,error)
	DebitBalance(ctx context.Context, data *entity.WalletInput) error
	CreditBalance(ctx context.Context, data *entity.WalletInput) error

}


type UseCase interface {
	GetUserByAccountNumber(ctx context.Context, req *wallet.GetUserByAccountNumberRequest) (*wallet.GetUserByAccountNumberResponse, error)
}