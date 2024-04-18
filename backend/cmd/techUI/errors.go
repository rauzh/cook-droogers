package techUI

import "errors"

var (
	ErrInvalidInput error = errors.New("invalid input")
	ErrEXIT         error = errors.New("exit")
	ErrCase         error = errors.New("invalid menu action")
)
