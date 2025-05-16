package user

import (
	"context"
	"source-base-go/entity"
)

type UserRepository interface {
	// create user
	// create infor user for db (token)
	FindByEmail(email string) (*entity.User, error)
	CreateUser(user *entity.User) error
	SaveToken(ctx context.Context, accessToken *entity.AccessToken, refreshToken *entity.RefreshToken) error
}

type UseCase interface {
	Login(email string, password string) (*entity.TokenPair, *entity.User, error)
	Register(user *entity.User) error
	SaveToken(ctx context.Context, accessToken *entity.AccessToken, refreshToken *entity.RefreshToken) error
}
