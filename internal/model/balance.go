package model

type Balance struct {
	WalletID       string  `json:"walletID"`
	Balance        float32 `json:"balance"`
	Amount         float32 `json:"amount"`
	AboveThreshold AboveThreshold
	UpdateVersion  int32 `json:"updateVersion"`
}
