package usecase

import (
	"context"
	"errors"
	"fmt"
	"source-base-go/golang/proto/user"
	"source-base-go/golang/service/transactionService/api/payload"
	"source-base-go/golang/service/transactionService/grpcclient"
	"source-base-go/golang/service/transactionService/model"
)

type TransactionService struct {
	walletClient    grpcclient.WalletClient
	authClient      grpcclient.AuthClient
	userClient      grpcclient.UserClient
	transactionRepo TransactionRepository
}

func NewOrderService(
	walletClient grpcclient.WalletClient,
	authClient grpcclient.AuthClient,
	userClient grpcclient.UserClient,
	transactionRepo TransactionRepository,
) *TransactionService {
	return &TransactionService{
		walletClient:    walletClient,
		authClient:      authClient,
		userClient:      userClient,
		transactionRepo: transactionRepo,
	}
}

func (s *TransactionService) GetReceiverInfo(ctx context.Context, accountNumber string) (*user.GetUserByAccountNumberResponse, error) {
	resp, err := s.userClient.GetUserByAccountNumber(ctx, accountNumber)
	if err != nil {
		return nil, errors.New("cannot find user: " + err.Error())
	}
	return resp, nil
}

func (s *TransactionService) TransferMoney(ctx context.Context, transferPayload *payload.TransferPayload) (map[string]interface{}, error) {
	res, err := s.authClient.VerifyJWT(ctx, transferPayload.Token)
	if err != nil || !res.IsValid {
		return nil, errors.New("unauthorized")
	}
	fromUserID := res.UserId

	debitInfo := model.DebitInfo{
		UserID: fromUserID,
		Amount: transferPayload.Amount,
	}
	_, err = s.walletClient.DebitBalance(ctx, debitInfo)
	if err != nil {
		return nil, fmt.Errorf("debit failed: %w", err)
	}

	creditInfo := model.CreditInfo{
		UserID: transferPayload.ToUserID,
		Amount: transferPayload.Amount,
	}
	_, err = s.walletClient.CreditBalance(ctx, creditInfo)
	if err != nil {
		_, _ = s.walletClient.RefundDebit(ctx, debitInfo)
		return nil, fmt.Errorf("credit failed, refunded sender: %w", err)
	}

	info := &model.Transaction{
		SenderID:   fromUserID,
		ReceiverID: transferPayload.ToUserID,
		Amount:     transferPayload.Amount,
		Status:     "success",
		Type:       "transfer",
	}
	err = s.transactionRepo.SaveTransaction(ctx, info)
	if err != nil {
		_, _ = s.walletClient.UndoCredit(ctx, creditInfo)
		_, _ = s.walletClient.RefundDebit(ctx, debitInfo)
		return nil, fmt.Errorf("transaction save failed, rollback triggered: %w", err)
	}

	_ = s.transactionRepo.SaveLog(ctx, model.TransactionLog{
		SenderID:        fromUserID,
		ReceiverID:      transferPayload.ToUserID,
		Amount:          transferPayload.Amount,
		TransactionType: "transfer",
	})

	return map[string]interface{}{
		"status":         "success",
		"transaction_id": info.ID,
	}, nil
}
