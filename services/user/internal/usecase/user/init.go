package user

import (
	"go-grpc-example/pkg/logger"
	"go-grpc-example/services/user/domain/repository"
	"go-grpc-example/services/user/domain/usecase"
	"time"
)

type userServiceInteracts struct {
	timeout  time.Duration
	logger   logger.Logger
	userRepo repository.UserRepository
}

func NewUserInteractor(
	timeout time.Duration,
	logger logger.Logger,
	userRepo repository.UserRepository,
) usecase.UserUseCase {
	return &userServiceInteracts{
		timeout:  timeout,
		logger:   logger,
		userRepo: userRepo,
	}
}
