package handler

import (
	"encoding/json"
	"errors"
	"go-musthave-diploma-tpl/internal/service"
	"go-musthave-diploma-tpl/internal/utils"
	"go-musthave-diploma-tpl/models"
	"net/http"
)

func (h *Handler) withdraw(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId, ok := utils.GetUserIdFromContext(ctx)
	if !ok {
		h.logger.Error().Msg("no user id was found in context")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var withdrawal models.Withdrawal
	if err := json.NewDecoder(r.Body).Decode(&withdrawal); err != nil {
		h.logger.Err(err).Msg("Invalid JSON was passed")
		http.Error(w, "Invalid JSON was passed", http.StatusBadRequest)
		return
	}

	err := h.withdrawalService.Withdraw(ctx, withdrawal, userId)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrNotCorrectOrderNumber):
			h.logger.Err(err).Msg("order number is not correct")
			http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
			return
		case errors.Is(err, service.ErrInsufficientFunds):
			h.logger.Err(err).Msg("not enough funds")
			http.Error(w, http.StatusText(http.StatusPaymentRequired), http.StatusPaymentRequired)
			return
		default:
			h.logger.Err(err).Msg("unexpected error occurred")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) getWithdrawals(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId, ok := utils.GetUserIdFromContext(ctx)
	if !ok {
		h.logger.Error().Msg("no user id was found in context")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	withdrawals, err := h.withdrawalService.GetWithdrawals(ctx, userId)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrNoWithdrawalsFound):
			h.logger.Err(err).Msg("no withdrawals found")
			http.Error(w, http.StatusText(http.StatusNoContent), http.StatusNoContent)
			return
		default:
			h.logger.Err(err).Msg("unexpected error occurred")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}

	withdrawalsJSON, err := json.Marshal(&withdrawals)
	if err != nil {
		h.logger.Err(err).Msg("error marshalling JSON")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(withdrawalsJSON)
}
