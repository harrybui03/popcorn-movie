package utils

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type ErrorCode string

const (
	ErrorCodeNotFound     ErrorCode = "404"
	ErrorCodeInternal     ErrorCode = "500"
	ErrorCodeBadRequest   ErrorCode = "400"
	ErrorCodeUnauthorized ErrorCode = "401"
)

type ErrorMessage string

const (
	ErrorMessageInvalidEmail    ErrorMessage = "Invalid Email"
	ErrorMessageInvalidPassword ErrorMessage = "Password must have at least one upper case , one lower case , one special"
	ErrorMessageNotEqual        ErrorMessage = "%s and %s not equal"
	ErrorMessageInternal        ErrorMessage = "Internal Server Error"
	ErrorInUse                  ErrorMessage = "%s is already in use"
	ErrorIncorrect              ErrorMessage = "%s is incorrect"
	ErrorUnauthorizedRequest    ErrorMessage = "Unauthorized Request"
	ErrorExpired                ErrorMessage = "%s is expired"
)

var (
	ErrorExpiredToken = errors.New("token has expired")
	ErrorInvalidToken = errors.New("invalid token")
)

func WrapGQLError(ctx context.Context, message string, code ErrorCode) *gqlerror.Error {
	e := &gqlerror.Error{
		Message: message,
		Extensions: map[string]interface{}{
			"code": code,
		},
	}

	if ctx != nil {
		e.Path = graphql.GetPath(ctx)
	}

	return e
}
