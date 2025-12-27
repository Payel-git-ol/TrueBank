package server

import (
	"TrueBankUserService/internal/grpc/userservicepb"
	"TrueBankUserService/internal/service"
	"TrueBankUserService/pkg/database"
	"TrueBankUserService/pkg/models"
	"context"
	"log"
)

type UserServer struct {
	userservicepb.UnimplementedUserServiceServer
}

func (s *UserServer) CreateUser(ctx context.Context, req *userservicepb.NewUserRequest) (*userservicepb.NewUserResponse, error) {
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

	return &userservicepb.NewUserResponse{
		Status:  "Success",
		Message: "User created",
		UserId:  123,
	}, nil
}
