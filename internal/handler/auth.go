package handler

import (
	"encoding/json"
	"errors"
	"go-musthave-diploma-tpl/internal/service"
	"go-musthave-diploma-tpl/internal/store"
	"go-musthave-diploma-tpl/models"
	"net/http"
)

func (h *Handlers) register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.logger.Err(err).Msg("Invalid JSON was passed")
		http.Error(w, "Invalid JSON was passed", http.StatusBadRequest) // 400
		return
	}

	err := h.authService.RegisterUser(user)

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

	w.WriteHeader(http.StatusOK)
}

func (h *Handlers) login(writer http.ResponseWriter, request *http.Request) {
	// TODO: implement me!
}
