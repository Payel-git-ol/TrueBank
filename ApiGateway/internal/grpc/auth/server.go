package auth

import (
	"ApiGateway/internal/grpc/authpb"
	"context"
	"google.golang.org/grpc"
	"log"
)

func NewAuthClient() authpb.AuthServiceClient {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return authpb.NewAuthServiceClient(conn)
}

func CallRegistration(client authpb.AuthServiceClient, req *authpb.UserRequest) (*authpb.AuthResponse, error) {
	return client.Registration(context.Background(), req)
}

func CallAuthentication(client authpb.AuthServiceClient, req *authpb.UserRequest) (*authpb.AuthResponse, error) {
	return client.Authentication(context.Background(), req)
}
