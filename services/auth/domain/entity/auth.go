package entity

import (
	"errors"
	"go-grpc-example/pkg/common"
	"go-grpc-example/services/auth/domain/valueobject"
	"time"

	"github.com/hashicorp/go-multierror"
)

type Auth struct {
	Id        common.ID
	App       valueobject.AppType
	UserId    common.ID
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AuthDTO struct {
	Id       *common.ID
	App      valueobject.AppTypeEnum
	UserId   common.ID
	Email    string
	Password string
}

func NewAuth(authDTO *AuthDTO) (*Auth, *multierror.Error) {
	var multierr *multierror.Error

	if authDTO.Id == nil {
		id := common.NewID()
		authDTO.Id = &id
	}

	app := valueobject.NewAppType(authDTO.App)

	auth := &Auth{
		Id:       *authDTO.Id,
		App:      app,
		UserId:   authDTO.UserId,
		Email:    authDTO.Email,
		Password: authDTO.Password,
	}

	if errValidate := auth.Validate(); errValidate != nil {
		multierr = multierror.Append(multierr, errValidate)
	}

	if multierr != nil {
		return nil, multierr
	}

	return auth, nil
}

func (a *Auth) Validate() *multierror.Error {
	var multierr *multierror.Error

	if a.Email == "" {
		multierr = multierror.Append(multierr, errors.New("email cannot be empty"))
	}

	if a.Password == "" {
		multierr = multierror.Append(multierr, errors.New("password cannot be empty"))
	}

	return multierr
}
