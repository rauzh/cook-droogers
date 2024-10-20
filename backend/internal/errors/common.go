package errors

import "github.com/pkg/errors"

var ErrNoSuchInstance error = errors.New("no such instance")
