package adapter

import (
	"context"
	"go-musthave-diploma-tpl/models"
)

type AccrualAdapter interface {
	GetOrderAccrual(ctx context.Context, orderNumber string) (models.Order, error)
}
