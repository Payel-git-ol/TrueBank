package user

type User struct {
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Balance  float64 `json:"balance"`
	Role     string  `json:"role"`
}

type UserResponse struct {
	User User `json:"User"`
}
