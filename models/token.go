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
	UserID       int64  `json:"-"`
}

func (t *Token) GetUserID() (int64, error) {
	userIDString, err := t.GetSubject()
	if err != nil {
		return 0, fmt.Errorf("error extracting UserID from token: %w", err)
	}

	userID, err := strconv.ParseInt(userIDString, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("error converting UserID from token to int64: %w", err)
	}

	return userID, nil
}

func (t *Token) String() string {
	return t.SignedString
}
