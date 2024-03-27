package errors

import "errors"

var (
	ErrNoGenre  error = errors.New("no track genre provided")
	ErrNoArtist error = errors.New("no artists provided")
	ErrNoType   error = errors.New("no track type provided")
	ErrNoTitle  error = errors.New("no track title provided")
)
