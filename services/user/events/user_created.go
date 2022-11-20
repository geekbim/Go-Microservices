package events

import (
	"context"
	"go-grpc-example/services/user/domain/entity"
)

var UserCreated userCreated

type UserCreatedPayload struct {
	Ctx  context.Context
	User *entity.User
}

type userCreated struct {
	handlers []interface{ Handle(UserCreatedPayload) }
}

func (u *userCreated) Register(handler interface{ Handle(UserCreatedPayload) }) {
	u.handlers = append(u.handlers, handler)
}

func (u *userCreated) Trigger(payload UserCreatedPayload) {
	for _, handler := range u.handlers {
		go handler.Handle(payload)
	}
}
