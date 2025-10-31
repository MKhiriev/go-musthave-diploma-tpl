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
	FindBalanceByUserID(ctx context.Context, userID int64) (models.Balance, error)
}

type OrderRepository interface {
	CreateOrderOrGetExisting(ctx context.Context, userID int64, orderNumber string) (models.Order, error)
	GetOrdersByUserID(ctx context.Context, userID int64) ([]models.Order, error)
	GetOrderByNumber(ctx context.Context, orderNumber string) (models.Order, error)
	GetOrdersByStatuses(ctx context.Context, statuses ...string) ([]models.Order, error)
	UpdateOrders(ctx context.Context, orders ...models.Order) error
}

type WithdrawalRepository interface {
	CreateWithdrawal(ctx context.Context, withdrawal models.Withdrawal, userID int64) error
	GetWithdrawalsByUserID(ctx context.Context, userID int64) ([]models.Withdrawal, error)
}

type UserBalanceRepository interface {
	CreateUserAndBalance(ctx context.Context, user models.User) (models.User, error)
}
