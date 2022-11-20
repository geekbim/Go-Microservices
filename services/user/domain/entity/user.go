package entity

import (
	"errors"
	"go-grpc-example/pkg/common"
	"go-grpc-example/services/user/domain/valueobject"
	"time"

	"github.com/hashicorp/go-multierror"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        common.ID
	App       valueobject.AppType
	Name      string
	Email     string
	password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserDTO struct {
	Id       *common.ID
	App      valueobject.AppTypeEnum
	Name     string
	Email    string
	Password string
}

func NewUser(userDTO *UserDTO) (*User, *multierror.Error) {
	var multierr *multierror.Error

	if userDTO.Id == nil {
		id := common.NewID()
		userDTO.Id = &id
	}

	app := valueobject.NewAppType(userDTO.App)

	user := &User{
		Id:        *userDTO.Id,
		App:       app,
		Name:      userDTO.Name,
		Email:     userDTO.Email,
		password:  userDTO.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if errValidate := user.Validate(); errValidate != nil {
		multierr = multierror.Append(multierr, errValidate)
	}

	if multierr != nil {
		return nil, multierr
	}

	return user, nil
}

func (u *User) Validate() *multierror.Error {
	var multierr *multierror.Error

	if u.Name == "" {
		multierr = multierror.Append(multierr, errors.New("name cannot be empty"))
	}

	if u.Email == "" {
		multierr = multierror.Append(multierr, errors.New("email cannot be empty"))
	}

	if u.password == "" {
		multierr = multierror.Append(multierr, errors.New("password cannot be empty"))
	}

	return multierr
}

func (u *User) GetPasswordHash() string {
	pwd, _ := bcrypt.GenerateFromPassword([]byte(u.password), bcrypt.DefaultCost)
	return string(pwd)
}
