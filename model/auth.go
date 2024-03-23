package model

import "github.com/go-playground/validator/v10"

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"min=8,containsany=!@#?"`
}

func (l *LoginRequest) Validate() error {
	var validate = validator.New()
	return validate.Struct(l)
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type CreateUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"min=8,containsany=!@#?"`
}

func (u *CreateUserRequest) Validate() error {
	var validate = validator.New()
	return validate.Struct(u)
}
