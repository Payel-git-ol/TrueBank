package server

import (
	"TrueBankUserService/internal/core/service"
	userservicepb2 "TrueBankUserService/internal/fetcher/grpc/userservicepb"
	"TrueBankUserService/pkg/database"
	"TrueBankUserService/pkg/models"
	"context"
	"log"
)

type UserServer struct {
	userservicepb2.UnimplementedUserServiceServer
}

func (s *UserServer) CreateUser(ctx context.Context, req *userservicepb2.NewUserRequest) (*userservicepb2.NewUserResponse, error) {
	log.Printf("Received new user: %s, %s", req.Username, req.Email)
	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Balance:  req.Balance,
	}

	err := service.SaveUserInCache(user)
	if err != nil {
		log.Printf("Error saving user: %v", err)
	}

	database.Db.Create(&user)

	return &userservicepb2.NewUserResponse{
		Status:  "Success",
		Message: "User created",
		UserId:  123,
	}, nil
}
