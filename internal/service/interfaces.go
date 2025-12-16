package service

import (
	"context"
	"go-musthave-diploma-tpl/models"
)

type AuthService interface {
	RegisterUser(ctx context.Context, user models.User) (models.User, error)
	Login(ctx context.Context, user models.User) (models.User, error)
	CreateToken(ctx context.Context, user models.User) (models.Token, error)
	ParseToken(ctx context.Context, tokenString string) (models.Token, error)
}

type BalanceService interface {
	GetBalanceByUserID(ctx context.Context, userID int64) (models.Balance, error)
}

type OrderService interface {
	AddOrder(ctx context.Context, userID int64, orderNumber string) error
	GetUserOrders(ctx context.Context, userID int64) ([]models.Order, error)
	GetOrder(ctx context.Context, orderNumber string) (models.Order, error)
	GetOrdersForAccrual(ctx context.Context) ([]models.Order, error)
	UpdateOrders(ctx context.Context, orders ...models.Order) error
}

type WithdrawalService interface {
	Withdraw(ctx context.Context, withdrawal models.Withdrawal, userID int64) error
	GetWithdrawals(ctx context.Context, userID int64) ([]models.Withdrawal, error)
}
