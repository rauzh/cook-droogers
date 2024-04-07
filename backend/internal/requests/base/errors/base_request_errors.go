package errors

import "errors"

var (
	ErrNoApplierID   error = errors.New("no applier id provided")
	ErrInvalidType   error = errors.New("invalid request type")
	ErrAlreadyClosed       = errors.New("request is already closed")
)
