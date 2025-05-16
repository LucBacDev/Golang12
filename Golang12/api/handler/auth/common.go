package handler

import (
	authPayload "source-base-go/api/payload/auth"
	"source-base-go/entity"
	"github.com/google/uuid"
)

func convertRegisterPayloadToEntity(data *authPayload.Register) *entity.User {
	return &entity.User{
		Username:    data.Username,
		Password:    data.Password,
		Email:       data.Email,
	}
}

func convertTokenPayloadToEntity(tokenPair *entity.TokenPair, userId int) (*entity.AccessToken, *entity.RefreshToken) {
	accessTokenEntity := &entity.AccessToken{
		UserId: userId,
		Token:  tokenPair.Token,
		Jti:    generateJti(),
		IssuedAt:  tokenPair.AccessTokenIssuedAt,
		ExpiresAt: tokenPair.AccessTokenExpiresAt,
	}

	refreshTokenEntity := &entity.RefreshToken{
		UserId: userId,
		Token:  tokenPair.RefreshToken,
		Jti:    generateJti(),
		IssuedAt:  tokenPair.RefreshTokenIssuedAt,
		ExpiresAt: tokenPair.RefreshTokenExpiresAt,
	}

	return accessTokenEntity, refreshTokenEntity
}

func generateJti() string {
	return uuid.New().String()
}
