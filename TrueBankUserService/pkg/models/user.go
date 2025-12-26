package models

type User struct {
	Username   string  `json:"username"`
	Email      string  `json:"email"`
	Password   string  `json:"password"`
	CardNumber string  `json:"card_number"`
	Balance    float64 `json:"balance"`
	Role       string  `json:"role"`
}
