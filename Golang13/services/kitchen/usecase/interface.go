package usecase

import (
	"context"
	"source-base-go/common/genproto/orders"
	"source-base-go/services/kitchen/api/payload"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error)
	GetOrders(ctx context.Context, request *orders.GetOrdersRequest) (*orders.GetOrderResponse, error)
}

type UseCase interface {
	CreateOrder(ctx context.Context, orderPayload *payload.OrderPayload) (*orders.CreateOrderResponse, error)
	GetOrders(ctx context.Context, customerID int32) (*orders.GetOrderResponse, error)
}
