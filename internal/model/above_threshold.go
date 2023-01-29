package model

import "time"

type AboveThreshold struct {
	WalletID          string    `json:"walletID"`
	TimeWindowBalance float32   `json:"timeWindowBalance"`
	Amount            float32   `json:"amount"`
	TimeExpired       time.Time `json:"timeExpired"`
	IsAboveThreshold  bool      `json:"isAboveThreshold"`
	UpdateVersion     int32     `json:"updateVersion"`
}
