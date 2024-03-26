package errors

import (
	"errors"
)

var (
	ErrNoNickname error = errors.New("no nickname provided")
)
