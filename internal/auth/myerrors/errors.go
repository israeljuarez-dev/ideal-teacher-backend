package myerrors

import "errors"

var (
	InvalidEmailOrPassword = errors.New("invalid email or password")
)

type AuthError struct {
	Msg string
	Err error
}

func (e *AuthError) Error() string {
	return e.Msg
}

func (e *AuthError) Unwrap() error {
	return e.Err
}
