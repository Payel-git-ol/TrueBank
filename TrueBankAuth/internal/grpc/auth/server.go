package auth

import (
	"context"

	"TrueBankAuth/internal/grpc/authpb"
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
