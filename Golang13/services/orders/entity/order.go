package entity

import "time"

type Order struct {
	ID        int     `gorm:"primaryKey"`
	UserID    int    `gorm:"column:user_id"`
	CreatedAt time.Time 
	UpdatedAt time.Time 
}

type OrderInput struct {
	UserID    int
	ProductID int
	Quantity  int
}
