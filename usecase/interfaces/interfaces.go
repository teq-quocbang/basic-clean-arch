package interfaces

import (
	"context"

	"github.com/quocbang/learn/usecase/model"
)

type Auth interface {
	Login(context.Context, model.Login) (*model.LoginReply, error)
	CreateUser(context.Context, model.CreateUser) error
}

type UseCase interface {
	Auth() Auth
}
