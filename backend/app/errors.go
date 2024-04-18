package app

import "errors"

var (
	ErrInitDB error = errors.New("can't init db")
)
