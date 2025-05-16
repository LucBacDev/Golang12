package handler

import (
	"fmt"
	"net/http"
	authPayload "source-base-go/api/payload/auth"
	authPresenter "source-base-go/api/presenter/auth"
	"source-base-go/usecase/user"

	"github.com/gin-gonic/gin"
)

func login(ctx *gin.Context, authService user.UseCase) {
	var payload authPayload.Login
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		return
	}

	//Login
	tokenPair, user, err := authService.Login(payload.Email, payload.Password)

	accessTokenEntity, refreshTokenEntity := convertTokenPayloadToEntity(tokenPair, user.Id)
	fmt.Println("abca",accessTokenEntity)
	err = authService.SaveToken(ctx, accessTokenEntity, refreshTokenEntity)
	if err != nil {
		return
	}

	result := &authPresenter.AuthResult{
		AccessToken:  tokenPair.Token,
		RefreshToken: tokenPair.RefreshToken,
		UserId:       user.Id,
		Username:     user.Username,
	}

	//Response in JSON
	response := &authPresenter.AuthResp{
		Status:  fmt.Sprint(http.StatusOK),
		Message: "success",
		Results: result,
	}

	ctx.JSON(http.StatusOK, response)
}

func register(ctx *gin.Context, authService user.UseCase) {
	var payload authPayload.Register
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		return
	}
	err = authService.Register(convertRegisterPayloadToEntity(&payload))
	if err != nil {
		return 
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
}
