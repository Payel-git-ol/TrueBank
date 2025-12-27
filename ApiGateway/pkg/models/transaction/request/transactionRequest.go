package request

type TransactionRequest struct {
	Username   string `json:"username"`
	Sum        string `json:"sum"`
	NumberCard string `json:"number_card"`
}
