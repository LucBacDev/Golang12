package usecase

import (
	"context"
	"source-base-go/common/genproto/orders"
	"source-base-go/services/orders/grpc/handler"
	"source-base-go/services/orders/grpc/payload"

	"google.golang.org/grpc"
)

type Service struct {
	orderRepo orderRepository
	orders.UnimplementedOrderServiceServer
}

func NewService(orderRepo orderRepository, grpc *grpc.Server) *Service{
	gRPCHandler := &Service{
		orderRepo: orderRepo,
	}
	orders.RegisterOrderServiceServer(grpc, gRPCHandler)
	return &Service{
		orderRepo: orderRepo,
	}
}

func (h *Service) GetOrders(ctx context.Context, req *orders.GetOrdersRequest) (*orders.GetOrderResponse, error) {
	ordersList, err := h.orderRepo.GetOrders(ctx,req.GetCustomerID()) 
	if err != nil {
		return nil, err
	}

	var ordersResponse []*orders.Order
	for _, order := range ordersList {
		ordersResponse = append(ordersResponse, &orders.Order{
			OrderID:   order.OrderID,
			CustomerID: order.CustomerID,
			Quantity:  order.Quantity,
		})
	}

	res := &orders.GetOrderResponse{
		Orders: ordersResponse,
	}

	return res, nil
}

func (h *Service) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) ( *orders.CreateOrderResponse, error) {
	order := &payload.Order{
		UserID: int(req.GetCustomerID()),
        ProductID:  int(req.GetProductID()),
        Quantity:   int(req.GetQuantity()),	
	}
	err := h.orderRepo.Create(handler.ConvertOrderPayloadToEntity(order))
	if err != nil {
		return nil, err
	}

	res := &orders.CreateOrderResponse{
		Status: "success",
	}
	return res, nil
}
