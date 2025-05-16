package wallet

import (
	"source-base-go/entity"
	"source-base-go/infrastructure/repository/util"
)

type Service struct {
	walletRepo WalletRepository
	verifier       util.Verifier

}

func NewService(walletRepo WalletRepository, verifier util.Verifier) *Service {
	return &Service{
		walletRepo: walletRepo,
		verifier: verifier,
	}
}

func (s Service) GetAllWallet() ([]*entity.Wallet, error) {
	return s.walletRepo.GetAllWallet()
}

func (s Service) CreateWallet(wallet *entity.Wallet) error{
	err := s.walletRepo.Create(wallet)
	if err != nil {
		return err
	}
	return nil
}
func (s Service) UpdateWallet(id int, wallet *entity.Wallet) error{
	err := s.walletRepo.Update(id, wallet)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) DeleteWallet(id int) error {
	err := s.walletRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

