package usecase

import (
	"context"
	"go-grpc-example/pkg/exceptions"
	"go-grpc-example/services/user/domain/entity"
)

type UserUseCase interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, *exceptions.CustomerError)
}
