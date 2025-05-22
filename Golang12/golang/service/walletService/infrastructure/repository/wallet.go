package repository

import (
	"context"
	"errors"
	"source-base-go/golang/service/walletService/entity"

	"gorm.io/gorm"
)

type WalletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) *WalletRepository {
	return &WalletRepository{
		db: db,
	}
}

func (w WalletRepository) GetUserByAccountNumber(ctx context.Context, accountNumber string) (*entity.UserResponse, error) {
	var result entity.Result

	if err := w.db.WithContext(ctx).
		Table("wallets").
		Select("wallets.user_id, users.name, wallets.account_number").
		Joins("LEFT JOIN users ON users.id = wallets.user_id").
		Where("wallets.account_number = ?", accountNumber).
		Scan(&result).Error; err != nil {
		return nil, err
	}

	resp := &entity.UserResponse{
		Status:        "success",
		UserID:        result.UserID,
		Name:          result.Name,
		AccountNumber: result.AccountNumber,
	}
	return resp, nil
}
func (r *WalletRepository) DebitBalance(ctx context.Context, data *entity.WalletInput) error {
	var wallet entity.Wallet

	if err := r.db.WithContext(ctx).Where("user_id = ?", data.UserID).First(&wallet).Error; err != nil {
		return err
	}

	if wallet.Balance < data.Amount {
		return errors.New("insufficient balance")
	}

	wallet.Balance -= data.Amount

	if err := r.db.WithContext(ctx).Save(&wallet).Error; err != nil {
		return err
	}

	return nil
}
func (r *WalletRepository) CreditBalance(ctx context.Context, data *entity.WalletInput) error {
	var wallet entity.Wallet

	if err := r.db.WithContext(ctx).Where("user_id = ?", data.UserID).First(&wallet).Error; err != nil {
		return err
	}

	if wallet.Balance < data.Amount {
		return errors.New("insufficient balance")
	}

	wallet.Balance += data.Amount

	if err := r.db.WithContext(ctx).Save(&wallet).Error; err != nil {
		return err
	}

	return nil
}
