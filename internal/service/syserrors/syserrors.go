package syserrors

import "errors"

var (
	ErrUnknown                     = errors.New("unknown error")
	ErrUserNotFound                = errors.New("user not found")
	ErrWrongCredentials            = errors.New("wrong credentials")
	ErrEmailOrUsernameAlreadyInUse = errors.New("email or username already in use")
)
