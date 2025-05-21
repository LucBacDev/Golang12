package handler

import (
	"source-base-go/golang/service/transactionService/usecase"

	"github.com/gin-gonic/gin"
)


func MakeHandlers(app *gin.Engine, transactionService usecase.UseCase) {
	transactionGroup := app.Group("/api/")
	{
		transactionGroup.POST("/transfer", func(ctx *gin.Context) {
			transferMoney(ctx, transactionService)
		})
		transactionGroup.POST("/transfer", func(ctx *gin.Context) {
			transferMoney(ctx, transactionService)
		})
	
	}
}
