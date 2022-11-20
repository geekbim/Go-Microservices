package user

import (
	"context"
	"go-grpc-example/pkg/exceptions"
	"go-grpc-example/services/user/domain/entity"
	"go-grpc-example/services/user/events"

	"github.com/hashicorp/go-multierror"
)

func (u *userServiceInteracts) CreateUser(ctx context.Context, user *entity.User) (*entity.User, *exceptions.CustomerError) {
	var multierr *multierror.Error

	c, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	user, err := u.userRepo.StoreUser(c, user)
	if err != nil {
		multierr = multierror.Append(multierr, err)
		return nil, &exceptions.CustomerError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

	events.UserCreated.Trigger(events.UserCreatedPayload{
		Ctx:  ctx,
		User: user,
	})

	return user, nil
}
