package model

type DebitInfo struct {
	UserID string
	Amount int32
}

type CreditInfo struct {
	UserID string
	Amount int32
}
type WalletResult struct {
	Success bool
	Message string
}