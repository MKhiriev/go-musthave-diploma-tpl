package adapter

import (
	"context"
	"fmt"
	"go-musthave-diploma-tpl/internal/config"
	"go-musthave-diploma-tpl/internal/logger"
	"go-musthave-diploma-tpl/internal/utils"
	"go-musthave-diploma-tpl/models"
	"net/http"
	"net/url"
)

type accrualAdapter struct {
	accrualAddress string
	accrualRoute   string
	*utils.HTTPClient

	logger *logger.Logger
}

func NewAccrualAdapter(cfg *config.Adapter, logger *logger.Logger) AccrualAdapter {
	return &accrualAdapter{
		accrualAddress: cfg.AccrualAddress,
		accrualRoute:   cfg.AccrualRoute,
		HTTPClient:     utils.NewHTTPClient(),
		logger:         logger,
	}
}

func (a *accrualAdapter) GetOrderAccrual(ctx context.Context, orderNumber string) (models.Order, error) {
	var accrual models.Accrual

	route, pathJoinError := url.JoinPath(a.accrualAddress, a.accrualRoute, orderNumber)
	if pathJoinError != nil {
		a.logger.Err(pathJoinError).Msg("url join error")
		return models.Order{}, fmt.Errorf("url join error: %w", pathJoinError)
	}

	response, err := a.HTTPClient.R().
		SetDebug(false).
		SetContext(ctx).
		SetResult(&accrual).
		Get(route)
	if err != nil {
		return models.Order{}, fmt.Errorf("commucation with accrual system error: %w", err)
	}

	a.logger.Debug().
		Str("route", route).
		Int("status", response.StatusCode()).
		Any("accrual", accrual).
		Bytes("body", response.Body()).
		Msg("ADAPTER")

	switch response.StatusCode() {
	case http.StatusInternalServerError:
		return models.Order{}, ErrAccrualInternalServerError
	case http.StatusNoContent:
		return models.Order{}, ErrOrderNotRegisteredInAccrual
	case http.StatusTooManyRequests:
		retryAfter := response.Header().Get("Retry-After")
		return models.Order{}, fmt.Errorf("%w: %s", ErrTooManyAccrualRequestsRetryAfter, retryAfter)
	case http.StatusOK:
		return models.Order{Number: accrual.Order, Accrual: accrual.Accrual, StatusText: accrual.Status}, nil
	default:
		return models.Order{}, ErrUndefinedAccrualStatusReturned
	}
}
