package handler

import (
	"context"
	"log"
	"net/http"
	"source-base-go/golang/service/transactionService/api/payload"
	"source-base-go/golang/service/transactionService/usecase"
	"time"

	"github.com/gin-gonic/gin"
)
func GetReceiverInfo(ctx *gin.Context, transactionService usecase.UseCase) {
	accountNumber := ctx.Param("accountNumber")
	token := ctx.GetHeader("Authorization")
	ctxWithTimeout, cancel := context.WithTimeout(ctx.Request.Context(), 2*time.Second)
	defer cancel()

	resp, err := transactionService.GetReceiverInfo(ctxWithTimeout, accountNumber, token)
	if err != nil {
		log.Println("GetReceiverInfo error:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get receiver info",})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func transferMoney(ctx *gin.Context, transactionService usecase.UseCase) {
	var payload payload.TransferPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	ctxWithTimeout, cancel := context.WithTimeout(ctx.Request.Context(), 2*time.Second)
	defer cancel()

	
	resp, err := transactionService.TransferMoney(ctxWithTimeout, &payload)

	if err != nil{
		log.Fatalf("client errL %v", err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"TransferMoneyResponse": resp,
	})

	ctx.JSON(http.StatusOK, resp)
}
