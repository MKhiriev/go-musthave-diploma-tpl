package handler

import (
	"context"
	"errors"
	"go-musthave-diploma-tpl/internal/service"
	"go-musthave-diploma-tpl/internal/utils"
	"net/http"
	"strings"
)

func (h *Handler) auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// token is expired case
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			h.logger.Err(ErrEmptyAuthorizationHeader).Send()
			http.Error(w, ErrEmptyAuthorizationHeader.Error(), http.StatusUnauthorized)
			return
		}

		tokenString, err := getTokenFromAuthHeader(authHeader)
		if err != nil {
			h.logger.Err(err).Send()
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := r.Context()
		token, err := h.authService.ParseToken(ctx, tokenString)

		if err != nil {
			switch {
			case errors.Is(err, service.ErrTokenIsExpired):
				h.logger.Err(err).Msg("token expired")
				http.Error(w, service.ErrTokenIsExpired.Error(), http.StatusUnauthorized)
				return
			default:
				h.logger.Err(err).Msg("error occurred during parsing token")
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
		}

		ctx = context.WithValue(ctx, utils.UserIDCtxKey, token.UserID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getTokenFromAuthHeader(authHeader string) (string, error) {
	parts := strings.Split(authHeader, " ")
	if len(parts) < 2 {
		return "", ErrInvalidAuthorizationHeader
	}

	tokenString := parts[1]
	if tokenString == "" {
		return "", ErrEmptyToken
	}

	return tokenString, nil
}
