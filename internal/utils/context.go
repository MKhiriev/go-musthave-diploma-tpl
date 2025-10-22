package utils

import (
	"context"
)

type contextKey string

func (c contextKey) String() string {
	return string(c)
}

var UserId = contextKey("userId")

func GetUserIdFromContext(ctx context.Context) (int64, bool) {
	userId, ok := ctx.Value(UserId).(int64)
	return userId, ok
}
