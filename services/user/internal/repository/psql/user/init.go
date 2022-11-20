package postgres_repository

import (
	"database/sql"
	"go-grpc-example/services/user/domain/repository"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepository{
		db: db,
	}
}
