package usecase

import (
	"context"
	"errors"
	"source-base-go/golang/kitchen/api/payload"
	"source-base-go/golang/kitchen/grpcclient"
	"source-base-go/golang/proto/auth"
	"source-base-go/golang/proto/user"
	"source-base-go/golang/proto/wallet"
	"source-base-go/golang/transactionService/model"
	"source-base-go/services/kitchen/api/payload"
)

type TransactionService struct {
	walletClient grpcclient.WalletClient
	authClient   grpcclient.AuthClient
	userClient   grpcclient.UserClient
	transactionRepo       TransactionRepository
}

func NewOrderService(
	walletClient 	wallet.WalletServiceClient,
	authClient 		grpcclient.AuthClient,
	userClient 		grpcclient.UserClient,
	transactionRepo      	TransactionRepository,
	) *TransactionService {
	return &TransactionService{
		walletClient: walletClient,
		authClient:   authClient,
		userClient: 	 userClient,
		transactionRepo:       transactionRepo,
	}
}

func (s *TransactionService) GetReceiverInfo(ctx context.Context, accountNumber string) (*user.GetUserByAccountNumberResponse, error) {
	resp, err := s.userClient.GetUserByAccountNumber(ctx, accountNumber)
	if err != nil {
		return nil, errors.New("cannot find user: " + err.Error())
	}
	return resp, nil
}

func (s *TransactionService) TransferMoney(ctx context.Context, tranferPayload payload.TransferPayload) (map[string]interface{}, error) {
	// Verify token
	res, err := s.authClient.VerifyJWT(ctx, tranferPayload.Token)
	if err != nil || !res.IsValid {
		return nil, errors.New("unauthorized")
	}
	fromUserID := res.UserId

	// Start transaction
	tx, err := s.transactionRepo.BeginTx(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			s.transactionRepo.RollbackTx(tx)
		}
	}()

	// Call Wallet Service - Debit
	debitInfo := model.DebitInfo{UserID: fromUserID, Amount: tranferPayload.Amount}
	_, err = s.walletClient.DebitBalance(ctx, debitInfo)
	if err != nil {
		return nil, errors.New("debit failed: " + err.Error())
	}

	// Call Wallet Service - Credit
	creditInfo := model.CreditInfo{UserID: tranferPayload.ToUserID, Amount: tranferPayload.Amount}
	_, err = s.walletClient.CreditBalance(ctx, creditInfo)
	if err != nil {
		return nil, errors.New("credit failed: " + err.Error())
	}

	// Save log
	err = s.transactionRepo.SaveLog(ctx, tx, model.TransactionLog{
		SenderID:        fromUserID,
		ReceiverID:      tranferPayload.ToUserID,
		Amount:          tranferPayload.Amount,
		TransactionType: "transfer",
	})
	if err != nil {
		return nil, err
	}

	// Commit transaction
	err = s.transactionRepo.CommitTx(tx)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"status":         "success",
		"transaction_id": "tx_" + generateRandomID(),
	}, nil
}