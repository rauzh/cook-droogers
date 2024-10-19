package session

import "errors"

var ErrInvalidUsernameOrPassword error = errors.New("invalid username or password")
var ErrCantSetAdminRole error = errors.New("cant set admin role")
