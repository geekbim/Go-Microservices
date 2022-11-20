package auth

import (
	"context"
	"go-grpc-example/pkg/exceptions"
	"go-grpc-example/services/auth/domain/entity"
	"go-grpc-example/services/auth/domain/valueobject"

	"github.com/hashicorp/go-multierror"
	"golang.org/x/crypto/bcrypt"
)

func (i *authInteractor) LoginByEmail(ctx context.Context, app *valueobject.AppType, auth *entity.Auth) (*entity.Auth, *exceptions.CustomerError) {
	var multierr *multierror.Error

	res, err := i.authRepo.FindAuthByEmail(ctx, app.GetValue(), auth.Email)
	if err != nil {
		multierr = multierror.Append(multierr, err)
		return nil, &exceptions.CustomerError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(auth.Password))
	if err != nil {
		multierr = multierror.Append(multierr, err)
		return nil, &exceptions.CustomerError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

	return auth, nil
}
