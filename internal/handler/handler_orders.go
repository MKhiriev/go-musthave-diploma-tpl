package handler

import (
	"errors"
	"go-musthave-diploma-tpl/internal/service"
	"go-musthave-diploma-tpl/internal/utils"
	"io"
	"net/http"
)

func (h *Handler) order(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId, ok := utils.GetUserIdFromContext(ctx)
	if !ok {
		h.logger.Error().Msg("no user id was found in context")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.logger.Err(err).Msg("invalid data was passed")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	orderNumber := string(body)
	err = h.orderService.AddOrder(ctx, userId, orderNumber)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidOrderNumber):
			h.logger.Err(err).Msg("order number is invalid")
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		case errors.Is(err, service.ErrNotCorrectOrderNumber):
			h.logger.Err(err).Msg("order number is not correct")
			http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
			return
		case errors.Is(err, service.ErrOrderWasUploadedByCurrentUser):
			h.logger.Err(err).Msg("order number was already uploaded by user")
			http.Error(w, http.StatusText(http.StatusOK), http.StatusOK)
			return
		case errors.Is(err, service.ErrOrderWasUploadedByAnotherUser):
			h.logger.Err(err).Msg("order number was uploaded by another user")
			http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
			return
		default:
			h.logger.Err(err).Msg("unexpected error occurred")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusAccepted)
}

func (h *Handler) getOrders(w http.ResponseWriter, r *http.Request) {
	// TODO: implement me!
}
