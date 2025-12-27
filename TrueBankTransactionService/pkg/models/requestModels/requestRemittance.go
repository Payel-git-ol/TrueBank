package requestModels

type RemittanceRequest struct {
	Username         string  `json:"username"`
	SenderСardNumber string  `json:"senderСardNumber"`
	GetterCardNumber string  `json:"getterCardNumber"`
	Sum              float64 `json:"sum"`
}
