package postgres_repository

import (
	"context"
	"go-grpc-example/services/auth/domain/entity"
	"go-grpc-example/services/auth/internal/repository/psql/helper"

	"github.com/rocketlaunchr/dbq/v2"
)

func (r *authRepository) CreateAuth(ctx context.Context, auth *entity.Auth) error {
	var err error

	_ = dbq.Tx(ctx, r.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		err = helper.StoreAuth(ctx, E, auth)
		if err != nil {
			return
		}
		_ = txCommit()
	})

	return err
}
