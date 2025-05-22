package repository

import (
	"context"
	"errors"
	"fmt"
	"source-base-go/golang/service/userService/entity"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
func (r UserRepository) FindByEmail(email string) (*entity.User, error) {
	user := &entity.User{}
	err := r.db.
		Where("email = ?", email).
		First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (r UserRepository) CreateUser(user *entity.User) error {
	//Hash password
	fmt.Println("User:", user)

	err := user.HashPassword()
	if err != nil {
		return err
	}

	//Create user
	err = r.db.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}
func (r *UserRepository) SaveToken(ctx context.Context, accessToken *entity.AccessToken, refreshToken *entity.RefreshToken) error {
	fmt.Println("Access Token:", accessToken)
	fmt.Println("Refresh Token:", refreshToken)
	if err := r.db.Create(&accessToken).Error; err != nil {
		return err
	}

	if err := r.db.Create(&refreshToken).Error; err != nil {
		return err
	}

	return nil
}
