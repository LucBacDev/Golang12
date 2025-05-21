package payload


type TransferPayload struct {
	ToUserID string `json:"to_user_id"`
	Amount   int32 `json:"amount"`
	Token	string `json:"token"`
}
