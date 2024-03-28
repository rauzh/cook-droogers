package errors

import "errors"

var (
	ErrNoTitle error = errors.New("no release title provided")
	ErrNoDate  error = errors.New("no release date creation provided")
)
