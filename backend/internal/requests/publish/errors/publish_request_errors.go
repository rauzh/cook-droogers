package errors

import "errors"

var (
	ErrNoReq       error = errors.New("no request provided")
	ErrInvalidDate error = errors.New("invalid date provided. it should be at least week later")
	ErrNoApplierID error = errors.New("no applier id provided")
	ErrNoReleaseID error = errors.New("no release id provided")
	ErrInvalidType error = errors.New("invalid request type")
)
