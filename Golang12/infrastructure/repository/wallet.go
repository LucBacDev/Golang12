package repository

import (
	"source-base-go/entity"
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

func (r WalletRepository) GetAllWallet() ([]*entity.Wallet, error) {
	var listWallet []*entity.Wallet
	err := r.db.Find(&listWallet).Error
	if err != nil {
		return nil, err
	}

	return listWallet, nil
}

func (r WalletRepository) Create(data *entity.Wallet) error {
	err := r.db.Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func (r WalletRepository) Update(id int, data *entity.Wallet) error {
	err := r.db.Model(&entity.Wallet{}).Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func (r WalletRepository) Delete(id int) error {
	err := r.db.Delete(&entity.Wallet{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
