package errors

import "errors"

var (
	ErrNoReq           error = errors.New("no request provided")
	ErrNickname        error = errors.New("invalid nickname provided")
	ErrNoApplierID     error = errors.New("no applier id provided")
	ErrInvalidType     error = errors.New("invalid request type")
	ErrCantFindManager error = errors.New("cant find manager")
)
