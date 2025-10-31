package utils

import (
	"errors"
	"fmt"
	"go-musthave-diploma-tpl/models"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWTToken(issuer string, userID int64, tokenDuration time.Duration, signKey string) (models.Token, error) {
	if issuer == "" || tokenDuration == 0 || signKey == "" {
		return models.Token{}, errors.New("invalid params for generating JWT Token")
	}

	now := time.Now()
	claims := &jwt.RegisteredClaims{
		Issuer:    issuer,
		Subject:   strconv.FormatInt(userID, 10),
		ExpiresAt: jwt.NewNumericDate(now.Add(tokenDuration)),
		IssuedAt:  jwt.NewNumericDate(now),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(signKey))
	if err != nil {
		return models.Token{}, fmt.Errorf("error occurred during singing JWT token: %w", err)
	}

	return models.Token{Token: token, SignedString: tokenString}, nil
}

func ValidateAndParseJWTToken(tokenString, tokenSignKey, tokenIssuer string) (models.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.Token{}, func(token *jwt.Token) (any, error) {
		return []byte(tokenSignKey), nil
	}, jwt.WithIssuer(tokenIssuer))
	if err != nil {
		return models.Token{}, fmt.Errorf("error occurred validating and parsing token: %w", err)
	}

	userIDStr, err := token.Claims.GetSubject()
	if err != nil {
		return models.Token{}, fmt.Errorf("error occurred during getting subject from token: %w", err)
	}
	if userIDStr == "" {
		return models.Token{}, errors.New("empty subject error")
	}

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		return models.Token{}, fmt.Errorf("error occurred during converting subject to UserIDCtxKey: %w", err)
	}

	return models.Token{Token: token, UserID: userID}, err
}
