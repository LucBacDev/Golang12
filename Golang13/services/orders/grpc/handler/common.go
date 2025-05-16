package handler

import (
	"source-base-go/services/orders/entity"
	orderPayload "source-base-go/services/orders/grpc/payload"
)

func ConvertOrderPayloadToEntity(data *orderPayload.Order) *entity.OrderInput {
	return &entity.OrderInput{
		UserID:    data.UserID,
		ProductID: data.ProductID,
		Quantity:  data.Quantity,
	}
}
