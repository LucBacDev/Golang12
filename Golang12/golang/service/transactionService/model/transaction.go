package model

import "time"

type Transaction struct {
	ID         string `gorm:"primaryKey" json:"id"`
	SenderID   string
	ReceiverID string
	Amount     int32
	Type       string
	Status     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}