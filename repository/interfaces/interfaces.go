package interfaces

import (
	"context"

	"github.com/quocbang/learn/repository/orm/model"
)

type Auth interface {
	CreateUser(ctx context.Context, req model.User) error
	GetUserByUsername(ctx context.Context, username string) (model.User, error)
}

type Todo interface {
	Create(ctx context.Context) error
}

type Repository interface {
	Auth() Auth
	Todo() Todo
}
