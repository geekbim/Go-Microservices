package grpc_user

import "go-grpc-example/services/user/domain/usecase"

type server struct {
	userUseCase usecase.UserUseCase
}
