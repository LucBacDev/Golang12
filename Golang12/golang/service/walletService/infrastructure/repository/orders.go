package repository

import (
	"context"
	"source-base-go/common/genproto/orders"
	"source-base-go/services/orders/entity"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r OrderRepository) Create(data *entity.OrderInput) error {
	order := &entity.Order{
		UserID: data.UserID,
	}
	if err := r.db.Create(order).Error; err != nil {
		return err
	}
	orderDetail := &entity.OrderItem{
		OrderID:   order.ID, 
		ProductID: data.ProductID,
		Quantity:  data.Quantity,
	}

	if err := r.db.Create(orderDetail).Error; err != nil {
		return err
	}

	return nil
}

func (r *OrderRepository) GetOrders(ctx context.Context, customerID int32) ([]*orders.Order, error) {
	var entities []*entity.OrderItem

	err := r.db.Model(&entity.Order{}).
		Joins("JOIN order_items ON orders.id = order_items.order_id").
		Select("orders.id AS OrderID, orders.user_id AS CustomerID, SUM(order_items.quantity) AS Quantity").
		Where("orders.user_id = ?", customerID).
		Group("orders.id").
		Scan(&entities).Error

	if err != nil {
		return nil, err
	}
	var result []*orders.Order
	for _, e := range entities {
		result = append(result, &orders.Order{
			OrderID:    int32(e.OrderID),
			CustomerID: customerID,
			Quantity:   int32(e.Quantity),
		})
	}

	return result, nil
}