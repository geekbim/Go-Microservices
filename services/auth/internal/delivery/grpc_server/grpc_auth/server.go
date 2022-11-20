package grpc_auth

import "go-grpc-example/services/auth/domain/usecase"

type server struct {
	authUC usecase.AuthUseCase
}
