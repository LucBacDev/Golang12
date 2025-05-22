package handler

import (
	"source-base-go/golang/service/userService/usecase/user"

	"github.com/gin-gonic/gin"
)

func MakeHandlers(app *gin.Engine, authService user.UseCase) {
	authGroup := app.Group("/api/auth")
	{
		authGroup.POST("/login", func(ctx *gin.Context) {
			login(ctx, authService)
		})
		authGroup.POST("/register", func(ctx *gin.Context) {
			register(ctx, authService)
		})
	}
}
