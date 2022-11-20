package response

import (
	authPb "go-grpc-example/proto/_generated/auth"
	"go-grpc-example/services/auth/domain/entity"
)

func NewFindAuthResponse(auth *entity.Auth) *authPb.FindAuthEmailResponse {
	return &authPb.FindAuthEmailResponse{
		Data: &authPb.AuthEmail{
			Id:    auth.Id.String(),
			Email: auth.Email,
		},
	}
}
