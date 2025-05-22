package payload

type WalletPayload struct {
	UserId  int `json:"user_id"`
	Balance int `json:"balance"`
}

type WalletUpdatePayload struct {
	Id      int `json:"id"`
	UserId  int `json:"user_id"`
	Balance int `json:"balance"`
}
type WalletDeletePayload struct {
	Id      int `json:"id"`
}
