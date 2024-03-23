package usecase

import (
	rInterfaces "github.com/quocbang/learn/repository/interfaces"
	"github.com/quocbang/learn/usecase/auth"
	uInterfaces "github.com/quocbang/learn/usecase/interfaces"
)

type UseCase struct {
	repository rInterfaces.Repository
}

func (u *UseCase) Auth() uInterfaces.Auth {
	return auth.NewAuthUseCase(u.repository)
}

func NewUseCase(repo rInterfaces.Repository) uInterfaces.UseCase {
	return &UseCase{
		repository: repo,
	}
}
