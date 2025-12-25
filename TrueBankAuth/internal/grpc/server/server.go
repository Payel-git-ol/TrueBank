package server

import (
	"context"
	"google.golang.org/grpc"
	"log"

	"TrueBankAuth/internal/grpc/authpb"
	"TrueBankAuth/internal/grpc/userservicepb"
	"TrueBankAuth/internal/service"
	"TrueBankAuth/pkg/models"
)

type AuthServer struct {
	authpb.UnimplementedAuthServiceServer
}

func (s *AuthServer) Registration(ctx context.Context, req *authpb.UserRequest) (*authpb.AuthResponse, error) {
	r := models.RequestUser{
		Username: req.GetUsername(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		Balance:  req.GetBalance(),
	}

	service.RegService(r)

	sendToUserService(r)

	return &authpb.AuthResponse{
		Status:  "Success",
		Message: "user saved",
	}, nil
}

func (s *AuthServer) Authentication(ctx context.Context, req *authpb.UserRequest) (*authpb.AuthResponse, error) {
	r := models.RequestUser{
		Username: req.GetUsername(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		Balance:  req.GetBalance(),
	}

	msg, err := service.AuthService(r)
	if err != nil {
		return &authpb.AuthResponse{
			Status:  "Error",
			Message: err.Error(),
		}, nil
	}

	return &authpb.AuthResponse{
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

	client := userservicepb.NewUserServiceClient(conn)

	req := &userservicepb.NewUserRequest{
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
