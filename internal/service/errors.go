package service

import "errors"

var (
	ErrInvalidDataProvided = errors.New("invalid data provided")
	ErrWrongPassword       = errors.New("wrong password")
)
