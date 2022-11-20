package usecase

import (
	"context"
	"go-grpc-example/pkg/exceptions"
	"go-grpc-example/services/auth/domain/entity"
	"go-grpc-example/services/auth/domain/valueobject"
)

type AuthUseCase interface {
	LoginByEmail(ctx context.Context, app *valueobject.AppType, auth *entity.Auth) (*entity.Auth, *exceptions.CustomerError)
	CreateAuth(ctx context.Context, auth *entity.Auth) *exceptions.CustomerError
}
