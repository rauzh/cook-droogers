package errors

import "errors"

var (
	ErrInvalidEmail    error = errors.New("the email is invalid")
	ErrInvalidName     error = errors.New("the name is invalid")
	ErrInvalidPassword error = errors.New("the password is invalid")
	ErrAlreadyTaken    error = errors.New("this email is already taken")
)
