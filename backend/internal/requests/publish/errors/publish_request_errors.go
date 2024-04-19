package errors

import "errors"

var (
	ErrNoReq                   error = errors.New("no request provided")
	ErrInvalidDate             error = errors.New("invalid date provided. it should be at least week later")
	ErrNoApplierID             error = errors.New("no applier id provided")
	ErrNoReleaseID             error = errors.New("no release id provided")
	ErrReleaseAlreadyPublished error = errors.New("release is already published")
	ErrNotOwner                error = errors.New("not the owner of release")
	ErrEndContract             error = errors.New("contract will have been ended by publication release date")
)
