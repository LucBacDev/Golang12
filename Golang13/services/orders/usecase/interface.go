package usecase

import (
	"context"
	"source-base-go/services/orders/entity"
	"source-base-go/common/genproto/orders"
)

type orderRepository interface {
	Create(order *entity.OrderInput) error
	GetOrders(ctx context.Context, customerID int32) ([]*orders.Order, error) 

}


type UseCase interface {
	CreateOrder(context.Context, *entity.OrderInput) error
	GetOrders(context.Context) ([]*orders.Order, error)
}