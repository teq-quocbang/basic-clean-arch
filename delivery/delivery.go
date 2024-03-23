package delivery

import (
	"github.com/labstack/echo/v4"
	"github.com/quocbang/learn/config"
	"github.com/quocbang/learn/delivery/auth"
	"github.com/quocbang/learn/repository"
	"github.com/quocbang/learn/usecase"
)

func NewDelivery(e *echo.Echo, cfg config.Config) error {
	// new repo
	repo, err := repository.NewRepository(repository.Database{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		UserName: cfg.DB.User,
		Password: cfg.DB.Password,
		Name:     cfg.DB.Name,
	})
	if err != nil {
		return err
	}

	// new use case
	useCase := usecase.NewUseCase(repo)

	v1 := e.Group("/v1")

	// register auth
	auth.RegisterAuthRoutes(v1.Group("/users"), useCase)

	return nil
}
