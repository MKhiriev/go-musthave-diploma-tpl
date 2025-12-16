package handler

import (
	"encoding/json"
	"errors"
	"go-musthave-diploma-tpl/internal/store"
	"go-musthave-diploma-tpl/internal/utils"
	"net/http"
)

func (h *Handler) getBalance(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := utils.GetUserIDFromContext(ctx)
	if !ok {
		h.logger.Error().Msg("userID not found in context")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	balance, err := h.balanceService.GetBalanceByUserID(ctx, userID)
	switch {
	case errors.Is(err, store.ErrNoBalanceFound) || errors.Is(err, store.ErrNoUserWasFound):
		h.logger.Err(err).Any("userID", userID).Msg("unexpected error occurred when finding balance for user")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	balanceJSON, err := json.Marshal(&balance)
	if err != nil {
		h.logger.Err(err).Any("balance", balance).Msg("balance marshalling failed")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(balanceJSON)
}
