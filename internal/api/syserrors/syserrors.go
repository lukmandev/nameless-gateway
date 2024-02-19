package syserrors

import (
	"github.com/vektah/gqlparser/v2/gqlerror"
)

var (
	ErrUnknown = &gqlerror.Error{
		Message: "Unknown error",
		Extensions: map[string]interface{}{
			"code": "UNKNOWN",
		},
	}
	ErrUserNotFound = &gqlerror.Error{
		Message: "User not found",
		Extensions: map[string]interface{}{
			"code": "USER_NOT_FOUND",
		},
	}
	ErrTokenExpired = &gqlerror.Error{
		Message: "Auth token expired",
		Extensions: map[string]interface{}{
			"code": "TOKEN_EXPIRED",
		},
	}
	ErrWrongCredentials = &gqlerror.Error{
		Message: "Wrong credentials",
		Extensions: map[string]interface{}{
			"code": "WRONG_CREDENTIALS",
		},
	}
	ErrEmailOrUsernameAlreadyInUse = &gqlerror.Error{
		Message: "Email or username already in use",
		Extensions: map[string]interface{}{
			"code": "ALREADY_IN_USE",
		},
	}
)
