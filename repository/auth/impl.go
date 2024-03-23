package auth

import (
	"context"

	"gorm.io/gorm"

	"github.com/quocbang/learn/repository/interfaces"
	"github.com/quocbang/learn/repository/orm/model"
)

type Auth struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) interfaces.Auth {
	return Auth{
		db: db,
	}
}

func (a Auth) CreateUser(ctx context.Context, req model.User) error {
	return a.db.Create(&req).Error
}

func (a Auth) GetUserByUsername(ctx context.Context, username string) (model.User, error) {
	var user model.User

	if err := a.db.Where("username = ?", username).Take(&user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}
