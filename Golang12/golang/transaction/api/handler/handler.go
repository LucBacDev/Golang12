package handler

import (
	"source-base-go/services/kitchen/usecase"

	"github.com/gin-gonic/gin"
)


func MakeHandlers(app *gin.Engine, orderService usecase.UseCase) {
	kitchenGroup := app.Group("/api/")
	{
		kitchenGroup.POST("/orders", func(ctx *gin.Context) {
			orderKitchen(ctx, orderService)
		})
	}
}
