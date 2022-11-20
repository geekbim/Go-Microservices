package grpc_auth

import (
	"context"
	"go-grpc-example/pkg/common"
	authPb "go-grpc-example/proto/_generated/auth"
	"go-grpc-example/services/auth/domain/entity"
	"go-grpc-example/services/auth/domain/valueobject"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) CreateAuth(ctx context.Context, req *authPb.CreateAuthRequest) (*authPb.CreateAuthResponse, error) {
	app, err := valueobject.StringToAppType(req.App)
	if err != nil {
		return nil, err
	}

	uid, err := common.StringToID(req.UserId)
	if err != nil {
		return nil, err
	}
	auth, errValidate := entity.NewAuth(&entity.AuthDTO{
		App:      app,
		UserId:   uid,
		Email:    req.Email,
		Password: req.Password,
	})
	if errValidate != nil {
		return nil, status.Error(codes.InvalidArgument, errValidate.Error())
	}

	errUseCase := s.authUC.CreateAuth(ctx, auth)
	if errUseCase != nil {
		return nil, status.Error(codes.InvalidArgument, errUseCase.Errors.Error())
	}

	return &authPb.CreateAuthResponse{}, nil
}
