package store

import "errors"

var (
	ErrLoginAlreadyExists = errors.New("login already exists")
	ErrNoUserWasFound     = errors.New("no user was found")
	ErrUserNotFound       = errors.New("user not found")

	ErrBalanceAlreadyExists = errors.New("balance for user already exists")
	ErrNoBalanceFound       = errors.New("user balance not found")

	ErrWithdrawalForOrderAlreadyExists = errors.New("withdrawal was already created for existing order")
	ErrWithdrawalWasNotCreated         = errors.New("withdrawal was not created")

	ErrOrderNotFound = errors.New("order not found")
)
