package model

type Login struct {
	Username string
	Password string
}

type LoginReply struct {
	AccessToken  string
	RefreshToken string
}

type CreateUser struct {
	Username string
	Password string
}
