package usecase

import (
	"context"
	"log"
	"source-base-go/common/genproto/orders"
	"source-base-go/services/kitchen/api/payload"
)

type OrderService struct {
	orderRepository orders.OrderServiceClient
}

func NewOrderService(orderRepository orders.OrderServiceClient) *OrderService {
	return &OrderService{
		orderRepository: orderRepository,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, orderPayload *payload.OrderPayload) (*orders.CreateOrderResponse, error) {
	req := &orders.CreateOrderRequest{
		CustomerID: orderPayload.CustomerID,
		ProductID:  orderPayload.ProductID,
		Quantity:   orderPayload.Quantity,
	}
	res, err := s.orderRepository.CreateOrder(ctx, req)
	if err != nil {
		log.Printf("Error creating order: %v", err)
		return nil, err
	}
	return res, nil
}

func (s *OrderService) GetOrders(ctx context.Context, customerID int32) (*orders.GetOrderResponse, error) {
	req := &orders.GetOrdersRequest{
		CustomerID: customerID,
	}
	res, err := s.orderRepository.GetOrders(ctx, req)
	if err != nil {
		log.Printf("Error fetching orders: %v", err)
		return nil, err
	}
	return res, nil
}
