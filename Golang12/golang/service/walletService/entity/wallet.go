package entity

import "time"

type Wallet struct {
	ID            int `gorm:"primaryKey"`
	UserID        string
	AccountNumber string
	Balance       int32
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type WalletInput struct {
	UserID string
	Amount int32
}
