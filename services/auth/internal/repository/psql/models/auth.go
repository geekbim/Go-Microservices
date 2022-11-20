package models

import (
	"go-grpc-example/services/auth/domain/valueobject"
	"time"
)

type Auth struct {
	Id        string                  `dbq:"id"`
	App       valueobject.AppTypeEnum `dbq:"app"`
	UserId    string                  `dbq:"user_id"`
	Email     string                  `dbq:"email"`
	Password  string                  `dbq:"password"`
	CreatedAt time.Time               `dbq:"created_at"`
	UpdatedAt time.Time               `dbq:"updated_at"`
}

func (Auth) TableName() string {
	return `auth`
}

func TableAuth() []string {
	return []string{
		"id",
		"app",
		"user_id",
		"email",
		"password",
		"created_at",
		"updated_at",
	}
}
