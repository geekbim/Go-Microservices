package postgres_repository

import (
	"context"
	"go-grpc-example/services/user/domain/entity"
	"go-grpc-example/services/user/internal/repository/psql/helper"

	"github.com/rocketlaunchr/dbq/v2"
)

func (r *userRepository) StoreUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	var err error

	_ = dbq.Tx(ctx, r.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		err = helper.StoreUser(ctx, E, user)
		if err != nil {
			return
		}
		_ = txCommit()
	})

	return user, err
}
