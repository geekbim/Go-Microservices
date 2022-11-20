package auth

import (
	"context"
	"go-grpc-example/pkg/exceptions"
	"go-grpc-example/services/auth/domain/entity"

	"github.com/hashicorp/go-multierror"
)

func (i *authInteractor) CreateAuth(ctx context.Context, auth *entity.Auth) *exceptions.CustomerError {
	var multierr *multierror.Error

	err := i.authRepo.CreateAuth(ctx, auth)
	if err != nil {
		multierr = multierror.Append(multierr, err)
		return &exceptions.CustomerError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

	return nil
}
