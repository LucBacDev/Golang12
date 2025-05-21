package entity

import "time"

type OrderItem struct {
	ID        int     `gorm:"primaryKey;autoIncrement"`
	OrderID   int     `gorm:"column:order_id"`
	ProductID int     `gorm:"column:product_id"`
	Quantity  int       `gorm:"column:quantity"`
	CreatedAt time.Time
	UpdatedAt time.Time 

}