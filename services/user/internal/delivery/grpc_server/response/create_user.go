package response

import (
	userPb "go-grpc-example/proto/_generated/user"
	"go-grpc-example/services/user/domain/entity"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewCreateUserResponse(user *entity.User) *userPb.CreateUserResponse {
	return &userPb.CreateUserResponse{
		Data: &userPb.User{
			Id:        user.Id.String(),
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: timestamppb.New(user.CreatedAt),
			UpdatedAt: timestamppb.New(user.UpdatedAt),
		},
	}
}
