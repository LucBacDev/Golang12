package handler

import (
	"source-base-go/golang/transactionService/usecase"

	"github.com/gin-gonic/gin"
)


func MakeHandlers(app *gin.Engine, transactionService usecase.UseCase) {
	kitchenGroup := app.Group("/api/")
	{
		kitchenGroup.POST("/transfer", func(ctx *gin.Context) {
			transferMoney(ctx, transactionService)
		})
	}
}
