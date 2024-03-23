package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	rInterfaces "github.com/quocbang/learn/repository/interfaces"
	rModel "github.com/quocbang/learn/repository/orm/model"
	uInterfaces "github.com/quocbang/learn/usecase/interfaces"
	"github.com/quocbang/learn/usecase/model"
	"github.com/quocbang/learn/utils/hash"
	"github.com/quocbang/learn/utils/token"
	"gorm.io/gorm"
)

type Auth struct {
	repository    rInterfaces.Repository
	tokenTimeLife time.Duration // TODO: bai tap
	secretKey     string        // TODO: bai tap
}



func NewAuthUseCase(repo rInterfaces.Repository) uInterfaces.Auth {
	return &Auth{
		repository: repo,
		
	}
}

func (a *Auth) Login(ctx context.Context, req model.Login) (*model.LoginReply, error) {
	// get user
	user, err := a.repository.Auth().GetUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	// check password
	if !hash.CheckPasswordHash(req.Password, user.Password) {
		return nil, fmt.Errorf("wrong password")
	}

	// generate access token
	accessToken, err := token.GenerateJWT(ctx, time.Duration(900), a.secretKey, user.Username)
	if err != nil {
		return nil, err
	}

	// generate refresh token
	refreshToken, err := token.GenerateJWT(ctx, time.Duration(64800000), "your-secret-key", user.Username)
	if err != nil {
		return nil, err
	}

	return &model.LoginReply{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (a Auth) CreateUser(ctx context.Context, req model.CreateUser) error {
	// get user
	if _, err := a.repository.Auth().GetUserByUsername(ctx, req.Username); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		} // go ahead
	} else {
		return fmt.Errorf("user existed")
	}

	// to hash password
	hashPassword, err := hash.HashPassword(req.Password)
	if err != nil {
		return err
	}

	// create user
	if err := a.repository.Auth().CreateUser(ctx, rModel.User{
		Username: req.Username,
		Password: hashPassword,
	}); err != nil {
		return err
	}

	return nil
}
