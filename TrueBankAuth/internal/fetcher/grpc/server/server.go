package server

import (
	service2 "TrueBankAuth/internal/core/service"
	authpb2 "TrueBankAuth/internal/fetcher/grpc/authpb"
	userservicepb2 "TrueBankAuth/internal/fetcher/grpc/userservicepb"
	"context"
	"google.golang.org/grpc"
	"log"

	"TrueBankAuth/pkg/models"
)

type AuthServer struct {
	authpb2.UnimplementedAuthServiceServer
}

func (s *AuthServer) Registration(ctx context.Context, req *authpb2.UserRequest) (*authpb2.AuthResponse, error) {
	r := models.RequestUser{
		Username: req.GetUsername(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		Balance:  req.GetBalance(),
	}

	service2.RegService(r)

	sendToUserService(r)

	return &authpb2.AuthResponse{
		Status:  "Success",
		Message: "user saved",
	}, nil
}

func (s *AuthServer) Authentication(ctx context.Context, req *authpb2.UserRequest) (*authpb2.AuthResponse, error) {
	r := models.RequestUser{
		Username: req.GetUsername(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		Balance:  req.GetBalance(),
	}

	msg, err := service2.AuthService(r)
	if err != nil {
		return &authpb2.AuthResponse{
			Status:  "Error",
			Message: err.Error(),
		}, nil
	}

	return &authpb2.AuthResponse{
		Status:  "Success",
		Message: msg,
	}, nil
}

func sendToUserService(user models.RequestUser) {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Printf("failed to connect to UserService: %v", err)
		return
	}
	defer conn.Close()

	client := userservicepb2.NewUserServiceClient(conn)

	req := &userservicepb2.NewUserRequest{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Balance:  user.Balance,
	}

	resp, err := client.CreateUser(context.Background(), req)
	if err != nil {
		log.Printf("error calling UserService: %v", err)
		return
	}

	log.Printf("UserService response: %s, %s, ID=%d", resp.Status, resp.Message, resp.UserId)
}
