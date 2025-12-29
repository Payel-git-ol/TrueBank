package auth

import (
	authpb2 "ApiGateway/internal/fetcher/grpc/authpb"
	"context"
	"google.golang.org/grpc"
	"log"
)

func NewAuthClient() authpb2.AuthServiceClient {
	conn, err := grpc.Dial("auth-service:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return authpb2.NewAuthServiceClient(conn)
}

func CallRegistration(client authpb2.AuthServiceClient, req *authpb2.UserRequest) (*authpb2.AuthResponse, error) {
	return client.Registration(context.Background(), req)
}

func CallAuthentication(client authpb2.AuthServiceClient, req *authpb2.UserRequest) (*authpb2.AuthResponse, error) {
	return client.Authentication(context.Background(), req)
}
