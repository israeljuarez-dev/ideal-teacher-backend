package myerrors

import "errors"

var (
	UserNotFoundError       = errors.New("user not found")
	EmailAlreadyExistsError = errors.New("email already registered")
)

type UserError struct {
	Msg string
	Err error
}

func (e *UserError) Error() string {
	return e.Msg
}

func (e *UserError) Unwrap() error {
	return e.Err
}
