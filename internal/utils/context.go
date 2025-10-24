package utils

import (
	"context"
)

type contextKey string

func (c contextKey) String() string {
	return string(c)
}

var UserIdCtxKey = contextKey("userId")

func GetUserIdFromContext(ctx context.Context) (int64, bool) {
	userId, ok := ctx.Value(UserIdCtxKey).(int64)
	return userId, ok
}
