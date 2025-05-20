package usecase

import (
	"context"
	"source-base-go/golang/orders/entity"
	"source-base-go/golang/proto/wallet"
)

type orderRepository interface {
	Create(order *entity.OrderInput) error
	GetOrders(ctx context.Context, customerID int32) ([]*wallet.Wallet, error) 

}


type UseCase interface {
	CreateOrder(context.Context, *entity.OrderInput) error
	GetOrders(context.Context) ([]*orders.Order, error)
}