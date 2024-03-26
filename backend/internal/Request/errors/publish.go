package errors

import "errors"

var (
	ErrLessThanWeek         error = errors.New("the publication date must be at least one week later")
	ErrInvalidMetaReleaseID       = errors.New("can't get release id from publication request")
)
