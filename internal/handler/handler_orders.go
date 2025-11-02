package handler

import (
	"encoding/json"
	"errors"
	"go-musthave-diploma-tpl/internal/service"
	"go-musthave-diploma-tpl/internal/utils"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) order(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := utils.GetUserIDFromContext(ctx)
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
	err = h.orderService.AddOrder(ctx, userID, orderNumber)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidOrderNumber) || errors.Is(err, service.ErrEmptyOrderNumber):
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
	ctx := r.Context()
	userID, ok := utils.GetUserIDFromContext(ctx)
	if !ok {
		h.logger.Error().Msg("no user id was found in context")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	orders, err := h.orderService.GetUserOrders(ctx, userID)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrNoUserOrdersFound):
			h.logger.Err(err).Msg("no data")
			http.Error(w, http.StatusText(http.StatusNoContent), http.StatusNoContent)
			return
		default:
			h.logger.Err(err).Msg("unexpected error occurred")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}

	ordersJSON, err := json.Marshal(&orders)
	if err != nil {
		h.logger.Err(err).Msg("error marshalling JSON")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(ordersJSON)
}

func (h *Handler) getOrderStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	orderNumber := chi.URLParam(r, "metricName")

	_, err := strconv.ParseInt(orderNumber, 10, 64)
	if err != nil {
		h.logger.Error().Msg("order number is not a number")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	order, err := h.orderService.GetOrder(ctx, orderNumber)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrOrderNotRegistered):
			h.logger.Err(err).Msg("order isn't registered")
			http.Error(w, http.StatusText(http.StatusNoContent), http.StatusNoContent)
			return
		case errors.Is(err, service.ErrTooManyRequests):
			h.logger.Err(err).Msg("too many requests to accrual service")
			http.Error(w, err.Error(), http.StatusTooManyRequests)
			return
		default:
			h.logger.Err(err).Msg("unexpected error occurred")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}

	orderJSON, err := json.Marshal(&order)
	if err != nil {
		h.logger.Err(err).Msg("error marshalling JSON")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(orderJSON)
}
