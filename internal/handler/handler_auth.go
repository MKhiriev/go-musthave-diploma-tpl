package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-musthave-diploma-tpl/internal/service"
	"go-musthave-diploma-tpl/internal/store"
	"go-musthave-diploma-tpl/models"
	"net/http"
)

func (h *Handler) register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.logger.Err(err).Msg("Invalid JSON was passed")
		http.Error(w, "Invalid JSON was passed", http.StatusBadRequest)
		return
	}

	registeredUser, err := h.authService.RegisterUser(ctx, user)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidDataProvided):
			h.logger.Err(err).Msg("invalid data provided")
			http.Error(w, "invalid data provided", http.StatusBadRequest)
			return
		case errors.Is(err, store.ErrLoginAlreadyExists):
			h.logger.Err(err).Msg("login already exists")
			http.Error(w, "login already exists", http.StatusConflict)
			return
		default:
			h.logger.Err(err).Msg("unexpected error occurred during user registration")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}

	token, err := h.authService.CreateToken(ctx, registeredUser)
	if err != nil {
		h.logger.Err(err).Msg("creation of token failed")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", token.SignedString))
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.logger.Err(err).Msg("Invalid JSON was passed")
		http.Error(w, "Invalid JSON was passed", http.StatusBadRequest)
		return
	}
	h.logger.Debug().Any("received user info", user).Send()

	foundUser, err := h.authService.Login(ctx, user)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidDataProvided):
			h.logger.Err(err).Msg("invalid data provided")
			http.Error(w, "invalid data provided", http.StatusBadRequest)
			return
		case errors.Is(err, store.ErrNoUserWasFound) || errors.Is(err, service.ErrWrongPassword):
			h.logger.Err(err).Msg("no user was found/wrong password")
			http.Error(w, "invalid login/password", http.StatusUnauthorized)
			return
		default:
			h.logger.Err(err).Msg("unexpected error occurred during user login")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}

	h.logger.Debug().Int64("id", foundUser.UserId).Any("found user", foundUser).Msg("user successfully logged in")

	token, err := h.authService.CreateToken(ctx, foundUser)
	if err != nil {
		h.logger.Err(err).Msg("creation of token failed")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", token.SignedString))
	w.WriteHeader(http.StatusOK)
}
