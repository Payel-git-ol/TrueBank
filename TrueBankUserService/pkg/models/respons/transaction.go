package respons

type ResultTransaction struct {
	Username   string  `json:"username"`
	Sum        float64 `json:"sum"`
	CardNumber string  `json:"card_number"`
}
