package grpc_auth

import (
	authPb "go-grpc-example/proto/_generated/auth"
	auth_repository "go-grpc-example/services/auth/domain/repository"
	"go-grpc-example/services/auth/internal/config"
	auth_usecase "go-grpc-example/services/auth/internal/usecase/auth"

	"google.golang.org/grpc"
)

func HandlerAuthService(
	s *grpc.Server,
	cfg *config.AuthConfig,
	authRepo auth_repository.AuthRepository,
) {
	authPb.RegisterAuthServiceServer(s, &server{
		authUC: auth_usecase.NewAuthInteractor(cfg, authRepo),
	})
}
