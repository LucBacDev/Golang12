package handler

import (
	"context"
	"log"
	"net/http"
	kitchenPayload "source-base-go/services/kitchen/api/payload"
	"source-base-go/services/kitchen/usecase"
	"time"

	"github.com/gin-gonic/gin"
)
func orderKitchen(ctx *gin.Context, orderService usecase.UseCase) {
	var payload kitchenPayload.OrderPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	ctxWithTimeout, cancel := context.WithTimeout(ctx.Request.Context(), 2*time.Second)
	defer cancel()


	resp, err := orderService.CreateOrder(ctxWithTimeout, &payload)

	if err != nil{
		log.Fatalf("client errL %v", err)
	}
	res, err := orderService.GetOrders(ctxWithTimeout, payload.CustomerID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"createOrderResponse": resp,
		"orders":              res.GetOrders(),
	})

	ctx.JSON(http.StatusOK, resp)
}