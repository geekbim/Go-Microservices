package repository

import (
	"context"
	"go-grpc-example/services/auth/domain/entity"
	"go-grpc-example/services/auth/domain/valueobject"
)

type AuthRepository interface {
	FindAuthByEmail(ctx context.Context, app valueobject.AppTypeEnum, email string) (*entity.Auth, error)
	CreateAuth(ctx context.Context, auth *entity.Auth) error
}
