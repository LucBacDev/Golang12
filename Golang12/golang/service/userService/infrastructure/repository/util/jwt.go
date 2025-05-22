package util

import (
	"source-base-go/golang/service/userService/config"
	"source-base-go/golang/service/userService/entity"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GenerateAccessToken(user *entity.User) (string, time.Time, time.Time, error) {
	random, err := uuid.NewRandom()
	if err != nil {
		return "", time.Time{}, time.Time{}, err
	}

	issuedAt := time.Now()
	expiresAt := issuedAt.Add(time.Duration(config.GetInt("jwt.accessMaxAge")))

	claims := jwt.MapClaims{
		"user_id": user.Id,
		"jti":     random.String(),
		"iat":     issuedAt.Unix(),
		"exp":     expiresAt.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(config.GetString("jwt.secretKey")))
	if err != nil {
		return "", time.Time{}, time.Time{}, err
	}

	return signedToken, issuedAt, expiresAt, nil
}

func GenerateRefreshToken(user *entity.User) (string, time.Time, time.Time, error) {
	issuedAt := time.Now()
	expiresAt := issuedAt.Add(time.Duration(config.GetInt("jwt.refreshMaxAge")))

	claims := jwt.MapClaims{
		"user_id": user.Id,
		"jti":     uuid.New().String(),
		"iat":     issuedAt.Unix(),
		"exp":     expiresAt.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(config.GetString("jwt.secretKey")))
	if err != nil {
		return "", time.Time{}, time.Time{}, err
	}

	return signedToken, issuedAt, expiresAt, nil
}
