package postgres_repository

import (
	"context"
	"errors"
	"fmt"
	"go-grpc-example/services/auth/domain/entity"
	"go-grpc-example/services/auth/domain/valueobject"
	"go-grpc-example/services/auth/internal/repository/psql/mapper"
	"go-grpc-example/services/auth/internal/repository/psql/models"

	"github.com/rocketlaunchr/dbq/v2"
)

func (r *authRepository) FindAuthByEmail(ctx context.Context, app valueobject.AppTypeEnum, email string) (*entity.Auth, error) {
	opts := &dbq.Options{
		SingleResult:   true,
		ConcreteStruct: models.Auth{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}

	stmt := fmt.Sprintf(`SELECT * FROM %s WHERE email = $1 AND app = $2 LIMIT 1`, models.Auth{}.TableName())

	result := dbq.MustQ(ctx, r.db, stmt, opts, email, app)
	if result != nil {
		return mapper.ToDomainAuth(result.(*models.Auth)), nil
	} else {
		return nil, errors.New("auth not found")
	}
}
