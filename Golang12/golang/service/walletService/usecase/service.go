package usecase

import (
	"context"
	"source-base-go/golang/proto/wallet"
	"source-base-go/golang/service/walletService/grpc/handler"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	walletRepo walletRepository
	wallet.UnimplementedWalletServiceServer
}

func NewService(walletRepo walletRepository, grpc *grpc.Server) *Service {
	gRPCHandler := &Service{
		walletRepo: walletRepo,
	}
	wallet.RegisterWalletServiceServer(grpc, gRPCHandler)
	return &Service{
		walletRepo: walletRepo,
	}
}

func (s *Service) GetUserByAccountNumber(ctx context.Context, req *wallet.GetUserByAccountNumberRequest) (*wallet.GetUserByAccountNumberResponse, error) {
	userID, ok := ctx.Value("user_id").(string)
	if !ok || userID == "" {
		return nil, status.Error(codes.Internal, "user ID missing from context")
	}

	ucResp, err := s.walletRepo.GetUserByAccountNumber(ctx, req.AccountNumber)
	if err != nil {
		return nil, err
	}
	return &wallet.GetUserByAccountNumberResponse{
		Status:        ucResp.Status,
		UserId:        ucResp.UserID,
		Name:          ucResp.Name,
		AccountNumber: ucResp.AccountNumber,
	}, nil

}

func (h *Service) DebitBalance(ctx context.Context, req *wallet.DebitRequest) (*wallet.DebitResponse, error) {
	userID, ok := ctx.Value("user_id").(string)
	if !ok || userID == "" {
		return nil, status.Error(codes.Internal, "user ID missing from context")
	}

	err := h.walletRepo.DebitBalance(ctx, handler.ConvertDebitToEntity(req))
	if err != nil {
		return nil, err
	}

	res := &wallet.DebitResponse{
		Success: true,
		Message: "Debit successful",
	}
	return res, nil
}
func (h *Service) CreditBalance(ctx context.Context, req *wallet.CreditRequest) (*wallet.CreditResponse, error) {
	userID, ok := ctx.Value("user_id").(string)
	if !ok || userID == "" {
		return nil, status.Error(codes.Internal, "user ID missing from context")
	}

	err := h.walletRepo.CreditBalance(ctx, handler.ConvertCreditToEntity(req))
	if err != nil {
		return nil, err
	}

	res := &wallet.CreditResponse{
		Success: true,
		Message: "Debit successful",
	}
	return res, nil
}

func (h *Service) RefundDebit(ctx context.Context, req *wallet.RefundDebitRequest) (*wallet.RefundDebitResponse, error) {
	userID, ok := ctx.Value("user_id").(string)
	if !ok || userID == "" {
		return nil, status.Error(codes.Internal, "user ID missing from context")
	}

	err := h.walletRepo.CreditBalance(ctx, handler.ConvertRefundDebitToEntity(req))
	if err != nil {
		return nil, err
	}

	res := &wallet.RefundDebitResponse{
		Success: true,
		Message: "Debit successful",
	}
	return res, nil
}
func (h *Service) UndoCredit(ctx context.Context, req *wallet.UndoCreditRequest) (*wallet.UndoCreditResponse, error) {
	userID, ok := ctx.Value("user_id").(string)
	if !ok || userID == "" {
		return nil, status.Error(codes.Internal, "user ID missing from context")
	}
	err := h.walletRepo.DebitBalance(ctx, handler.ConvertUndoCreditToEntity(req))
	if err != nil {
		return nil, err
	}

	res := &wallet.UndoCreditResponse{
		Success: true,
		Message: "Debit successful",
	}
	return res, nil
}
