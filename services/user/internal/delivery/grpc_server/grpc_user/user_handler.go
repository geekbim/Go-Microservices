package grpc_user

import (
	"go-grpc-example/pkg/logger"
	userPb "go-grpc-example/proto/_generated/user"
	user_repository "go-grpc-example/services/user/domain/repository"
	"go-grpc-example/services/user/internal/config"
	user_usecase "go-grpc-example/services/user/internal/usecase/user"

	"google.golang.org/grpc"
)

func HandlerUserService(
	s *grpc.Server,
	logger logger.Logger,
	cfg config.UserConfig,
	userRepo user_repository.UserRepository,
) {
	userPb.RegisterUserServiceServer(s, &server{
		userUseCase: user_usecase.NewUserInteractor(
			cfg.Timeout,
			logger,
			userRepo,
		),
	})
}
