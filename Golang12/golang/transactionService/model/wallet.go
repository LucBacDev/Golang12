package model

type DebitInfo struct {
	UserID string
	Amount int64
}

type CreditInfo struct {
	UserID string
	Amount int64
}
type WalletResult struct {
	Success bool
	Message string
}