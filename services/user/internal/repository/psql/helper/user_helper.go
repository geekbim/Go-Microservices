package helper

import (
	"context"
	"go-grpc-example/services/user/domain/entity"
	"go-grpc-example/services/user/internal/repository/psql/mapper"
	"go-grpc-example/services/user/internal/repository/psql/models"

	"github.com/rocketlaunchr/dbq/v2"
)

func StoreUser(ctx context.Context, E dbq.EFn, user *entity.User) error {
	userDbq := mapper.ToDbqStructUser(user)

	stmt := dbq.INSERTStmt(models.User{}.TableName(), models.TableUser(), len(userDbq), dbq.PostgreSQL)

	_, err := E(ctx, stmt, nil, userDbq)
	if err != nil {
		return err
	}

	return nil
}
