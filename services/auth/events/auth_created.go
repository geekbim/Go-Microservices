package events

import (
	"go-grpc-example/services/auth/domain/entity"
	"go-grpc-example/services/auth/domain/valueobject"
)

var AuthCreated authCreated

type AuthCreatedPayload struct {
	Auth *entity.Auth
	App  valueobject.AppType
}

type authCreated struct {
	handlers []interface{ Handle(AuthCreatedPayload) }
}

func (u *authCreated) Register(handler interface{ Handle(AuthCreatedPayload) }) {
	u.handlers = append(u.handlers, handler)
}

func (u *authCreated) Trigger(payload AuthCreatedPayload) {
	for _, handler := range u.handlers {
		go handler.Handle(payload)
	}
}
