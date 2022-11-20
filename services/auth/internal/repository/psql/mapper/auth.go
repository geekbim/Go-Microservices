package mapper

import (
	"go-grpc-example/pkg/common"
	"go-grpc-example/services/auth/domain/entity"
	"go-grpc-example/services/auth/domain/valueobject"
	"go-grpc-example/services/auth/internal/repository/psql/models"

	"github.com/rocketlaunchr/dbq/v2"
)

func ToDomainAuth(auth *models.Auth) *entity.Auth {
	newId, _ := common.StringToID(auth.Id)
	newApp := valueobject.NewAppType(auth.App)
	newUserId, _ := common.StringToID(auth.UserId)

	return &entity.Auth{
		Id:        newId,
		App:       newApp,
		UserId:    newUserId,
		Email:     auth.Email,
		Password:  auth.Password,
		CreatedAt: auth.CreatedAt,
		UpdatedAt: auth.UpdatedAt,
	}
}

func ToModelAuth(auth *entity.Auth) *models.Auth {
	return &models.Auth{
		Id:        auth.Id.String(),
		App:       auth.App.GetValue(),
		UserId:    auth.UserId.String(),
		Email:     auth.Email,
		Password:  auth.Password,
		CreatedAt: auth.CreatedAt,
		UpdatedAt: auth.UpdatedAt,
	}
}

func DataDbqAuth(auth *entity.Auth) []interface{} {
	return dbq.Struct(ToModelAuth(auth))
}

func ToDbqStructAuth(auth *entity.Auth) (dbqStruct []interface{}) {
	dbqStruct = append(dbqStruct, DataDbqAuth(auth))
	return
}
