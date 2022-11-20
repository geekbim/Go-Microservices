package grpc_auth

import (
	"context"
	"go-grpc-example/pkg/logger"
	authPb "go-grpc-example/proto/_generated/auth"
	"go-grpc-example/services/user/domain/entity"
)

func (p *AuthGrpcClient) CreateAuth(ctx context.Context, user *entity.User) error {

	payloads := &authPb.CreateAuthRequest{
		App:      user.App.String(),
		UserId:   user.Id.String(),
		Email:    user.Email,
		Name:     user.Name,
		Password: user.GetPasswordHash(),
	}

	_, err := p.authClient.CreateAuth(ctx, payloads)
	if err != nil {
		logger.CreateErrorLogEventService(logger.LogEventService{
			Event:   "create_user",
			From:    "user",
			To:      "auth",
			Payload: payloads,
			Message: err.Error(),
		})
		return err
	}

	logger.CreateLogEventService(logger.LogEventService{
		Event:   "create_user",
		From:    "user",
		To:      "auth",
		Payload: payloads,
		Message: "success hit grpc auth",
	})

	return nil
}
