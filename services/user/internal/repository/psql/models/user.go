package models

import (
	"go-grpc-example/services/user/domain/valueobject"
	"time"
)

type User struct {
	Id        string                  `dbq:"id"`
	App       valueobject.AppTypeEnum `dbq:"app"`
	Name      string                  `dbq:"name"`
	Email     string                  `dbq:"email"`
	CreatedAt time.Time               `dbq:"created_at"`
	UpdatedAt time.Time               `dbq:"updated_at"`
}

func (User) TableName() string {
	return `users`
}

func TableUser() []string {
	return []string{
		"id",
		"app",
		"name",
		"email",
		"created_at",
		"updated_at",
	}
}
