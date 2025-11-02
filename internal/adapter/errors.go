package adapter

import "errors"

var (
	ErrAccrualInternalServerError       = errors.New("accrual internal server error")
	ErrOrderNotRegisteredInAccrual      = errors.New("order is not registered in accrual")
	ErrTooManyAccrualRequestsRetryAfter = errors.New("retry accrual request. try again later")
	ErrUndefinedAccrualStatusReturned   = errors.New("accrual returned not defined status")
)
