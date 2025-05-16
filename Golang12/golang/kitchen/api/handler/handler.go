package handler

import (
	"source-base-go/services/kitchen/usecase"

	"github.com/gin-gonic/gin"
)


func MakeHandlers(app *gin.Engine, orderService usecase.UseCase) {
	kitchenGroup := app.Group("/api/")
	{
		kitchenGroup.POST("/transfer", func(ctx *gin.Context) {
			transferMoney(ctx, orderService)
		})
	}
}
