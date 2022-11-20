package mapper

import (
	"go-grpc-example/pkg/common"
	"go-grpc-example/services/user/domain/entity"
	"go-grpc-example/services/user/domain/valueobject"
	"go-grpc-example/services/user/internal/repository/psql/models"

	"github.com/rocketlaunchr/dbq/v2"
)

func ToDomainUser(u *models.User) *entity.User {
	newId, _ := common.StringToID(u.Id)
	app := valueobject.NewAppType(u.App)
	return &entity.User{
		Id:        newId,
		App:       app,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func ToModelUser(u *entity.User) *models.User {
	return &models.User{
		Id:        u.Id.String(),
		App:       u.App.GetValue(),
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func DataDbqUser(u *entity.User) []interface{} {
	return dbq.Struct(ToModelUser(u))
}

func ToDbqStructUser(u *entity.User) (dbqStruct []interface{}) {
	dbqStruct = append(dbqStruct, DataDbqUser(u))
	return
}
