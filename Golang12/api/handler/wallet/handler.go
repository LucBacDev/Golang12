package handler

import (
	"source-base-go/api/middleware"
	"source-base-go/infrastructure/repository/util"
	"source-base-go/usecase/wallet"

	"github.com/gin-gonic/gin"
)

func MakeHandlers(app *gin.Engine, walletService wallet.UseCase, verifier util.Verifier) {
	walletGroup := app.Group("/api/wallet")
	{
		walletGroup.Use(middleware.AuthMiddleware(verifier))

		walletGroup.GET("/getAll", func(ctx *gin.Context) {
			getAllWallet(ctx, walletService)
		})
		walletGroup.POST("/create", func(ctx *gin.Context) {
			createWallet(ctx, walletService)
		})
		walletGroup.PUT("/update", func(ctx *gin.Context) {
			updateWallet(ctx, walletService)
		})
		walletGroup.DELETE("/delete", func(ctx *gin.Context) {
			deleteWallet(ctx, walletService)
		})
	}
}
