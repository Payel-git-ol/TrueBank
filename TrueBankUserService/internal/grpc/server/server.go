package server

import (
	"TrueBankUserService/internal/grpc/userservicepb"
	"context"
	"log"
)

type UserServer struct {
	userservicepb.UnimplementedUserServiceServer
}

func (s *UserServer) CreateUser(ctx context.Context, req *userservicepb.NewUserRequest) (*userservicepb.NewUserResponse, error) {
	log.Printf("Received new user: %s, %s", req.Username, req.Email)

	return &userservicepb.NewUserResponse{
		Status:  "Success",
		Message: "User created",
		UserId:  123,
	}, nil
}
