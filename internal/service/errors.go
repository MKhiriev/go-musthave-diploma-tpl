package service

import "errors"

var (
	ErrInvalidDataProvided = errors.New("invalid data provided")
	ErrWrongPassword       = errors.New("wrong password")

	ErrInvalidOrderNumber            = errors.New("invalid order number")
	ErrEmptyOrderNumber              = errors.New("empty order number")
	ErrNotCorrectOrderNumber         = errors.New("order number is not luna correct")
	ErrOrderWasUploadedByCurrentUser = errors.New("order number was already uploaded by user")
	ErrOrderWasUploadedByAnotherUser = errors.New("order number was uploaded by another user")
	ErrNoUserOrdersFound             = errors.New("no user orders found")

	ErrWithdrawalSumIsZero = errors.New("withdrawal sum is zero")
	ErrInsufficientFunds   = errors.New("insufficient funds in the account")
	ErrNoWithdrawalsFound  = errors.New("no withdrawals found")

	ErrTokenIsExpired = errors.New("token is expired")
)
