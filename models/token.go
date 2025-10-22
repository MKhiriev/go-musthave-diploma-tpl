package models

import (
	"fmt"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
)

type Token struct {
	*jwt.Token `json:"-"`
	jwt.RegisteredClaims
	SignedString string `json:"-"`
	UserId       int64  `json:"-"`
}

func (t *Token) GetUserId() (int64, error) {
	userIdString, err := t.GetSubject()
	if err != nil {
		return 0, fmt.Errorf("error extracting UserId from token: %w", err)
	}

	userId, err := strconv.ParseInt(userIdString, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("error converting UserId from token to int64: %w", err)
	}

	return userId, nil
}

func (t *Token) String() string {
	return t.SignedString
}
