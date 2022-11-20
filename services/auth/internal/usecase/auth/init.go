package auth

import (
	"go-grpc-example/services/auth/domain/repository"
	"go-grpc-example/services/auth/domain/usecase"
	"go-grpc-example/services/auth/internal/config"
)

type authInteractor struct {
	authConfig *config.AuthConfig
	authRepo   repository.AuthRepository
}

func NewAuthInteractor(
	authConfig *config.AuthConfig,
	authRepo repository.AuthRepository,
) usecase.AuthUseCase {
	return &authInteractor{
		authConfig: authConfig,
		authRepo:   authRepo,
	}
}
