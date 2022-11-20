package repository

import (
	"context"
	"go-grpc-example/services/user/domain/entity"
)

type UserRepository interface {
	StoreUser(ctx context.Context, user *entity.User) (*entity.User, error)
}
