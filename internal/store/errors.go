package store

import "errors"

var (
	ErrLoginAlreadyExists = errors.New("login already exists")
	ErrNoUserWasFound     = errors.New("no user was found")
)
