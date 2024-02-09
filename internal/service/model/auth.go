package model

type LoginInput struct {
	EmailOrUsername string
	Password        string
}

type RegisterInput struct {
	Email    string
	Username string
	Password string
}
