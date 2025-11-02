package handler

import "errors"

var (
	ErrEmptyAuthorizationHeader   = errors.New("empty `Authorization` header")
	ErrInvalidAuthorizationHeader = errors.New("invalid `Authorization` header")
	ErrEmptyToken                 = errors.New("empty token in `Authorization` header")
)
