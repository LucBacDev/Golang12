package api

import (
	"encoding/json"
	"net/http"

	"google.golang.org/grpc"
	"source-base-go/golang/proto/wallet/wallet_grpc.pb.go"
	"context"
	"time"
)

type TransferRequest struct {
	Token     string `json:"token"`
	ToUserID  string `json:"to_user_id"`
	Amount    int64  `json:"amount"`
}

type TransferResponse struct {
	Success       bool   `json:"success"`
	Message       string `json:"message"`
	TransactionID string `json:"transaction_id,omitempty"`
}

func HandleTransferMoney(w http.ResponseWriter, r *http.Request) {
	var req TransferRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Gọi gRPC
	conn, err := grpc.Dial("localhost:50053", grpc.WithInsecure()) // TransactionService gRPC
	if err != nil {
		http.Error(w, "Failed to connect to gRPC", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	client := proto.NewTransactionServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	grpcResp, err := client.TransferMoney(ctx, &proto.TransferRequest{
		Token:     req.Token,
		ToUserId:  req.ToUserID,
		Amount:    req.Amount,
	})
	if err != nil {
		http.Error(w, "gRPC error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resp := TransferResponse{
		Success:       grpcResp.Success,
		Message:       grpcResp.Message,
		TransactionID: grpcResp.TransactionId,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
