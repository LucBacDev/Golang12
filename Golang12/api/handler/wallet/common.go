package handler

import (
	walletPayload "source-base-go/api/payload/wallet"
	walletPresenter "source-base-go/api/presenter/wallet"
	"source-base-go/entity"

	"github.com/gin-gonic/gin"
)

func convertListWalletEntityToPresenter(ctx *gin.Context, listData []*entity.Wallet) []*walletPresenter.Wallet {
	listWalletPresenter := []*walletPresenter.Wallet{}
	userId := ctx.GetInt("user_id")
	for _, wallet := range listData {
		walletPresenter := &walletPresenter.Wallet{
			Id: wallet.Id,
			UserId:        userId,
			Balance:        wallet.Balance,
			CreatedAt: wallet.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: wallet.UpdatedAt.Format("2006-01-02 15:04:05"),
		}
		listWalletPresenter = append(listWalletPresenter, walletPresenter)
	}

	return listWalletPresenter
}
func convertWalletPayloadToEntity(ctx *gin.Context, data *walletPayload.WalletPayload) *entity.Wallet {
	userId := ctx.GetInt("user_id")

	return &entity.Wallet{
		UserId: userId,
		Balance: data.Balance,
	}
}

func convertWalletUpdatePayloadToEntity(ctx *gin.Context, id int,data *walletPayload.WalletUpdatePayload) *entity.Wallet {
	userId := ctx.GetInt("user_id")

	return &entity.Wallet{
		Id: id,
		UserId: userId,
		Balance: data.Balance,
	}
}

