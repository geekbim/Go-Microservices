package grpc_user

import (
	"context"
	"go-grpc-example/pkg/exceptions"
	userPb "go-grpc-example/proto/_generated/user"
	"go-grpc-example/services/user/domain/entity"
	"go-grpc-example/services/user/domain/valueobject"
	"go-grpc-example/services/user/internal/delivery/grpc_server/response"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) CreateUser(ctx context.Context, req *userPb.CreateUserRequest) (*userPb.CreateUserResponse, error) {
	appTypeEnum, err := valueobject.StringToAppType(req.GetApp().String())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	user, errValidate := entity.NewUser(&entity.UserDTO{
		App:      appTypeEnum,
		Name:     req.GetBody().GetName(),
		Email:    req.GetBody().GetEmail(),
		Password: req.GetBody().GetPassword(),
	})
	if errValidate != nil {
		return nil, status.Error(codes.InvalidArgument, errValidate.Error())
	}

	res, errUseCase := s.userUseCase.CreateUser(ctx, user)
	if errUseCase != nil {
		return nil, status.Error(exceptions.MapToGrpcStatusCode(errUseCase.Status), errUseCase.Errors.Error())
	}

	return response.NewCreateUserResponse(res), nil
}
