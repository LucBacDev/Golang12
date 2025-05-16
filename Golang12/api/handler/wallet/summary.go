package handler

import (
	"fmt"
	"net/http"
	walletPayload "source-base-go/api/payload/wallet"
	walletPresenter "source-base-go/api/presenter/wallet"
	"strconv"

	"source-base-go/usecase/wallet"

	"github.com/gin-gonic/gin"
)

func getAllWallet(ctx *gin.Context, walletService wallet.UseCase) {
	listRole, err := walletService.GetAllWallet()
	if err != nil {
		return
	}
	//Response in JSON
	response := &walletPresenter.ListWalletResp{
		Status:  fmt.Sprint(http.StatusOK),
		Message: "Success",
		Results: convertListWalletEntityToPresenter(ctx, listRole),
	}

	ctx.JSON(http.StatusOK, response)
}

func createWallet(ctx *gin.Context, walletService wallet.UseCase){
	var payload walletPayload.WalletPayload
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		return
	}
	err = walletService.CreateWallet(convertWalletPayloadToEntity(ctx, &payload))
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
}
func updateWallet(ctx *gin.Context, walletService wallet.UseCase){
	var payload walletPayload.WalletUpdatePayload
	idStr := ctx.DefaultQuery("id", "") 
	if idStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id is not required"})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Id format"})
		return
	}
	err = ctx.ShouldBindJSON(&payload)
	fmt.Println("payload",payload)
	if err != nil {
		return
	}
	err = walletService.UpdateWallet(id,convertWalletUpdatePayloadToEntity(ctx, id,&payload))
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func deleteWallet(ctx *gin.Context, walletService wallet.UseCase) {
	idStr := ctx.DefaultQuery("id", "") 
	if idStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id is not required"})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Id format"})
		return
	}

	err = walletService.DeleteWallet(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Wallet deleted successfully"})
}