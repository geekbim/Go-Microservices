package notifier

import (
	"context"
	"go-grpc-example/pkg/logger"
	"go-grpc-example/services/user/domain/entity"
	"go-grpc-example/services/user/domain/repository"
	"go-grpc-example/services/user/events"
	"go-grpc-example/services/user/internal/delivery/grpc_client/grpc_auth"
)

type UserCreatedNotifier struct {
	Ctx      context.Context
	Logger   logger.Logger
	AuthGrpc *grpc_auth.AuthGrpcClient
	UserRepo repository.UserRepository
}

func (u *UserCreatedNotifier) createAuth(user *entity.User) {
	err := u.AuthGrpc.CreateAuth(u.Ctx, user)
	if err != nil {
		u.Logger.Error(err)
	}
}

func (u *UserCreatedNotifier) Handle(payload events.UserCreatedPayload) {
	u.createAuth(payload.User)
}
