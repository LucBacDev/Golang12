package presenter

type Wallet struct {
	Id int `json:"id"`
	UserId  int `json:"user_id"`
	Balance int `json:"balance"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type ListWalletResp struct {
	Status  string  `json:"status"`
	Message string  `json:"message"`
	Results []*Wallet `json:"results"`
}

type WalletResp struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Results *Wallet  `json:"results"`
}
