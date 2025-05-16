package entity

import (
	"time"
)

type Wallet struct {
	Id        int       `gorm:"column:id;primaryKey" json:"id"`
	UserId   int       `gorm:"column:user_id" json:"user_id"`
	Balance   int       `gorm:"column:balance" json:"balance"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}