package usecase

import (
	"context"
	"errors"
	"source-base-go/golang/kitchen/api/payload"
	"source-base-go/golang/kitchen/grpcclient"
	"source-base-go/golang/proto/auth"
	"source-base-go/golang/proto/user"
	"source-base-go/golang/proto/wallet"
	"source-base-go/services/kitchen/api/payload"
)

type TransactionService struct {
	walletClient wallet.WalletServiceClient
	authClient   grpcclient.AuthClient
	userClient       grpcclient.UserClient
}

func NewOrderService(walletClient wallet.WalletServiceClient, authClient grpcclient.AuthClient, userClient grpcclient.UserClient) *TransactionService {
	return &TransactionService{
		walletClient: walletClient,
		authClient:   authClient,
		userClient: 	 userClient,
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
	res, err := s.authClient.VerifyJWT(ctx, tranferPayload.Token)
	if err != nil || !res.IsValid {
		return nil, errors.New("unauthorized")
	}
	
	fromUserID := res.UserId
    // 2. Debit
	debitReq := s.walletClient.DebitBalance(ctx, req)
		&wallet.DebitRequest{
		UserId : fromUserID,
		Amount: tranferPayload.Amount,
	}

    _, err = s.walletRepository.DebitBalance(ctx, debitReq)
    if err != nil {
        return nil, errors.New("debit failed: " + err.Error())
    }

    // 3. Credit
    _, err = s.walletRepository.CreditBalance(ctx, tranferPayload.ToUserID, tranferPayload.Amount)
    if err != nil {
        return nil, errors.New("credit failed: " + err.Error())
    }

    // 4. Return response
    return map[string]interface{}{
        "status":         "success",
        "transaction_id": "tx_" + generateRandomID(),
    }, nil
}