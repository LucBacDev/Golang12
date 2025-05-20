// usecase/transfer.go
package usecase

import (
    "context"
    "errors"
    authpb "transaction/proto/auth"
    walletpb "transaction/proto/wallet"
    "transaction/grpcclient"
)

type TransferUsecase struct {
    AuthClient   grpcclient.AuthClient
    WalletClient grpcclient.WalletClient
}

func (uc *TransferUsecase) TransferMoney(ctx context.Context, token, toUserID string, amount int64) (map[string]interface{}, error) {
    // 1. Verify JWT
    verifyResp, err := uc.AuthClient.VerifyJWT(ctx, token)
    if err != nil || !verifyResp.IsValid {
        return nil, errors.New("unauthorized")
    }
    fromUserID := verifyResp.UserId

    // 2. Debit
    _, err = uc.WalletClient.Debit(ctx, fromUserID, amount)
    if err != nil {
        return nil, errors.New("debit failed: " + err.Error())
    }

    // 3. Credit
    _, err = uc.WalletClient.Credit(ctx, toUserID, amount)
    if err != nil {
        return nil, errors.New("credit failed: " + err.Error())
    }

    // 4. Return response
    return map[string]interface{}{
        "status":         "success",
        "transaction_id": "tx_" + generateRandomID(),
    }, nil
}
