package response

import "ApiGateway/pkg/models/user"

type UserResponse struct {
	User user.User `json:"User"`
}
