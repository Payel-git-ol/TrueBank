package models

type ResultRemittance struct {
	Username         string  `json:"username"`
	SenderСardNumber string  `json:"senderСardNumber"`
	GetterCardNumber string  `json:"getterCardNumber"`
	Sum              float64 `json:"sum"`
}
