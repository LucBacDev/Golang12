package wallet

import (
	"source-base-go/entity"
)

type WalletRepository interface {
	GetAllWallet() ([]*entity.Wallet, error)
	Create(data *entity.Wallet) error
	Update(id int, data *entity.Wallet) error
	Delete(id int) error

}

type UseCase interface {
	GetAllWallet() ([]*entity.Wallet, error)
	CreateWallet(Wallet *entity.Wallet) error
	UpdateWallet(id int, Wallet *entity.Wallet) error
	DeleteWallet(id int) error

}
