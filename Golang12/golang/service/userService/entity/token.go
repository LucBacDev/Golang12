package entity

import (
	"time"
)

type AccessToken struct {
	ID         int    `gorm:"primaryKey"`
	UserId     int    `json:"user_id" gorm:"column:user_id"`
	Jti        string `json:"jti" gorm:"column:jti"`
	Token      string `gorm:"type:text;not null;column:token"`
	IssuedAt   time.Time `gorm:"column:issued_at"`
	ExpiresAt  time.Time `gorm:"column:expires_at"`
}

type RefreshToken struct {
	ID         int    `gorm:"primaryKey"`
	UserId     int    `json:"user_id" gorm:"column:user_id"`
	Jti        string `json:"jti" gorm:"column:jti"`
	Token      string `gorm:"type:text;not null;column:token"`
	IssuedAt   time.Time `gorm:"column:issued_at"`
	ExpiresAt  time.Time `gorm:"column:expires_at"`
}

type TokenPair struct {
    Token                 string    `json:"token"`
    RefreshToken          string    `json:"refreshToken"`
	AccessTokenJti        string `json:"accessTokenjti"`
	RefreshTokenJti        string `json:"refreshTokenjti"`
    AccessTokenIssuedAt   time.Time `json:"accessTokenIssuedAt"`
    AccessTokenExpiresAt  time.Time `json:"accessTokenExpiresAt"`
    RefreshTokenIssuedAt  time.Time `json:"refreshTokenIssuedAt"`
    RefreshTokenExpiresAt time.Time `json:"refreshTokenExpiresAt"`
}