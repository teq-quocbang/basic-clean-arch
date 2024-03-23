package auth

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/quocbang/learn/model"
	"github.com/quocbang/learn/usecase/interfaces"
	uModel "github.com/quocbang/learn/usecase/model"
)

type Auth struct {
	useCase interfaces.UseCase
}

func RegisterAuthRoutes(g *echo.Group, useCase interfaces.UseCase) {
	a := Auth{
		useCase: useCase,
	}
	g.POST("/login", a.Login)
	g.POST("", a.CreateUser)
}

func (a Auth) Login(c echo.Context) error {
	var (
		req = model.LoginRequest{}
	)

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  43093402,
			"message": "invalid param",
		})
	}

	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  43093402,
			"message": err.Error(),
		})
	}

	reply, err := a.useCase.Auth().Login(context.Background(), uModel.Login{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  1277,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.LoginResponse{
		AccessToken:  reply.AccessToken,
		RefreshToken: reply.RefreshToken,
	})
}

func (a Auth) CreateUser(c echo.Context) error {
	var (
		req model.CreateUserRequest
	)

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  1277,
			"message": "invalid param",
		})
	}

	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  1277,
			"message": err.Error(),
		})
	}

	// create user
	if err := a.useCase.Auth().CreateUser(context.Background(), uModel.CreateUser{
		Username: req.Username,
		Password: req.Password,
	}); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  1678333,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status":  200,
		"message": "OK",
	})
}
