package user

import (
	"context"
	"source-base-go/entity"
	util "source-base-go/infrastructure/repository/util"
)

type Service struct {
	userRepository UserRepository
}

func NewService(userRepository UserRepository) *Service {
	return &Service{
		userRepository: userRepository,
	}
}

func (s *Service) Login(email string, password string) (*entity.TokenPair, *entity.User, error) {
	//Get user by user name
	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return nil, nil, err
	}

	if user == nil {
		return nil, nil, err
	}

	//Vefiry password
	passwordValidate := user.ValidatePassword(password)
	if !passwordValidate {
		return nil, nil, err
	}

	accessToken, issuedAtAccess, expiresAtAccess, err := util.GenerateAccessToken(user)
	if err != nil {
		return nil, nil, err
	}

	refreshToken, issuedAtRefresh, expiresAtRefresh, err := util.GenerateRefreshToken(user)
	if err != nil {
		return nil, nil, err
	}

	tokenPair := &entity.TokenPair{
		Token:                 accessToken,
		RefreshToken:          refreshToken,
		AccessTokenIssuedAt:   issuedAtAccess,
		AccessTokenExpiresAt:  expiresAtAccess,
		RefreshTokenIssuedAt:  issuedAtRefresh,
		RefreshTokenExpiresAt: expiresAtRefresh,
	}

	return tokenPair, user, nil
}

func (s Service) Register(user *entity.User) error {
	//Create user
	err := s.userRepository.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) SaveToken(ctx context.Context, accessToken *entity.AccessToken, refreshToken *entity.RefreshToken) error {
	err := s.userRepository.SaveToken(ctx, accessToken, refreshToken)
	if err != nil {
		return err
	}

	return nil

}
