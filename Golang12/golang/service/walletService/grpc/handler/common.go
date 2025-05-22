package handler

import (
	"source-base-go/golang/proto/wallet"
	"source-base-go/golang/service/walletService/entity"
)

func ConvertDebitToEntity(req *wallet.DebitRequest) *entity.WalletInput {
	return &entity.WalletInput{
		UserID:    req.UserId,
		Amount: req.Amount,
	}
}
func ConvertCreditToEntity(req *wallet.CreditRequest) *entity.WalletInput {
	return &entity.WalletInput{
		UserID:    req.UserId,
		Amount: req.Amount,
	}
}
func ConvertRefundDebitToEntity(req *wallet.RefundDebitRequest) *entity.WalletInput {
	return &entity.WalletInput{
		UserID:    req.UserId,
		Amount: req.Amount,
	}
}
func ConvertUndoCreditToEntity(req *wallet.UndoCreditRequest) *entity.WalletInput {
	return &entity.WalletInput{
		UserID:    req.UserId,
		Amount: req.Amount,
	}
}
