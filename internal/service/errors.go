package service

import "errors"

var (
	ErrInvalidDataProvided = errors.New("invalid data provided")
	ErrWrongPassword       = errors.New("wrong password")

	ErrInvalidOrderNumber            = errors.New("invalid order number")
	ErrNotCorrectOrderNumber         = errors.New("order number is not luna correct")
	ErrOrderWasUploadedByCurrentUser = errors.New("order number was already uploaded by user")
	ErrOrderWasUploadedByAnotherUser = errors.New("order number was uploaded by another user")
	ErrNoUserOrdersFound             = errors.New("no user orders found")

	ErrTokenIsExpired = errors.New("token is expired")
)
