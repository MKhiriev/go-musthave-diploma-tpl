package store

import (
	"context"
	"go-musthave-diploma-tpl/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user models.User) error
	FindUserByLogin(ctx context.Context, user models.User) (models.User, error)
}

type BalanceRepository interface {
	FindBalanceByUserId(ctx context.Context, userId int64) (models.Balance, error)
}

type OrderRepository interface {
	CreateOrderOrGetExisting(ctx context.Context, userId int64, orderNumber string) (models.Order, error)
	GetOrdersByUserId(ctx context.Context, userId int64) ([]models.Order, error)
}
