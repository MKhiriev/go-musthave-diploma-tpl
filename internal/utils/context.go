package utils

import (
	"context"
)

type contextKey string

func (c contextKey) String() string {
	return string(c)
}

var UserIDCtxKey = contextKey("userId")

func GetUserIDFromContext(ctx context.Context) (int64, bool) {
	userID, ok := ctx.Value(UserIDCtxKey).(int64)
	return userID, ok
}
