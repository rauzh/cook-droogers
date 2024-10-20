package errors

import "errors"

var ErrExists error = errors.New("manager already exists")
var ErrInvalidID error = errors.New("invalid id")
var ErrInvalidRole error = errors.New("this user can't be manager")
