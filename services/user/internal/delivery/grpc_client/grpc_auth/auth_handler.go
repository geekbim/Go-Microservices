package grpc_auth

import (
	"context"
	authPb "go-grpc-example/proto/_generated/auth"

	"google.golang.org/grpc"
)

type AuthGrpcClient struct {
	ctx        context.Context
	authClient authPb.AuthServiceClient
}

func NewAuthHandler(ctx context.Context, conn *grpc.ClientConn) (*AuthGrpcClient, error) {
	authClient := authPb.NewAuthServiceClient(conn)
	return &AuthGrpcClient{
		ctx:        ctx,
		authClient: authClient,
	}, nil
}
