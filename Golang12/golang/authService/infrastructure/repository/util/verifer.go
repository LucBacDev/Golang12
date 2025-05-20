package util

import (
	"errors"
	"gorm.io/gorm"
	"source-base-go/entity"
)

type Verifier interface {
	Verifier(userId int, token string) (bool, error)
}

type AccessTokenRepository struct {
	db *gorm.DB
}

func NewAccessTokenRepository(db *gorm.DB) *AccessTokenRepository {
	return &AccessTokenRepository{
		db: db,
	}
}

func (r *AccessTokenRepository) Verifier(userId int, token string) (bool, error) {
	var accessToken entity.AccessToken
	err := r.db.Where("user_id = ? AND token = ? AND expires_at < NOW()", userId, token).First(&accessToken).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
