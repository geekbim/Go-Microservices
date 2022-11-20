package events

import (
	"context"
	"database/sql"
	"go-grpc-example/pkg/logger"
	"go-grpc-example/services/user/events"
	"go-grpc-example/services/user/internal/delivery/grpc_client/grpc_auth"
	psql_user "go-grpc-example/services/user/internal/repository/psql/user"
	"go-grpc-example/services/user/notifier"
)

func EventHandler(authGrpcClient *grpc_auth.AuthGrpcClient, l logger.Logger, db *sql.DB) {
	createNotifier := notifier.UserCreatedNotifier{
		Ctx:      context.TODO(),
		Logger:   l,
		AuthGrpc: authGrpcClient,
		UserRepo: psql_user.NewUserRepository(db),
	}
	events.UserCreated.Register(&createNotifier)
}
