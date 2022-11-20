package helper

import (
	"context"
	"go-grpc-example/services/auth/domain/entity"
	"go-grpc-example/services/auth/internal/repository/psql/mapper"
	"go-grpc-example/services/auth/internal/repository/psql/models"

	"github.com/rocketlaunchr/dbq/v2"
)

func StoreAuth(ctx context.Context, E dbq.EFn, auth *entity.Auth) error {
	authDbq := mapper.ToDbqStructAuth(auth)

	stmt := dbq.INSERTStmt(models.Auth{}.TableName(), models.TableAuth(), len(authDbq), dbq.PostgreSQL)

	_, err := E(ctx, stmt, nil, authDbq)
	if err != nil {
		return err
	}

	return nil
}
