package postgres_repository

import (
	"database/sql"
	"go-grpc-example/services/auth/domain/repository"
)

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) repository.AuthRepository {
	return &authRepository{
		db: db,
	}
}
