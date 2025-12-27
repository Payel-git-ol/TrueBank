package request

type Replenishment struct {
	Username   string  `json:"username"`
	CardNumber string  `json:"card_number"`
	Sum        float64 `json:"sum"`
}
