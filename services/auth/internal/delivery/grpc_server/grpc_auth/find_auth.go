package grpc_auth

import (
	"context"
	authPb "go-grpc-example/proto/_generated/auth"
	"go-grpc-example/services/auth/domain/entity"
	"go-grpc-example/services/auth/domain/valueobject"
	"go-grpc-example/services/auth/internal/delivery/grpc_server/response"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) FindAuthEmail(ctx context.Context, req *authPb.FindAuthEmailRequest) (*authPb.FindAuthEmailResponse, error) {
	appTypeEnum, err := valueobject.StringToAppType(req.GetApp().String())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	auth, errValidate := entity.NewAuth(&entity.AuthDTO{
		App:      appTypeEnum,
		Email:    req.GetBody().GetEmail(),
		Password: req.GetBody().GetPassword(),
	})
	if errValidate != nil {
		return nil, status.Error(codes.InvalidArgument, errValidate.Error())
	}

	res, errUseCase := s.authUC.LoginByEmail(ctx, &auth.App, auth)
	if errUseCase != nil {
		return nil, status.Error(codes.InvalidArgument, errUseCase.Errors.Error())
	}

	return response.NewFindAuthResponse(res), nil
}
