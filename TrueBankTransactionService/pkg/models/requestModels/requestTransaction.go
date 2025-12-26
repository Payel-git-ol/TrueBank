package requestModels

type TransactionRequest struct {
	Username        string `json:"username"`
	NameTransaction string `json:"name_transaction"`
	Sum             string `json:"sum"`
	NumberCard      string `json:"number_card"`
}
