package service

import (
	"context"
	"go-musthave-diploma-tpl/models"
)

type AuthService interface {
	RegisterUser(ctx context.Context, user models.User) error
	Login(ctx context.Context, user models.User) (models.User, error)
	CreateToken(ctx context.Context, user models.User) (models.Token, error)
	ParseToken(ctx context.Context, tokenString string) (models.Token, error)
}

type BalanceService interface {
	GetBalanceByUserId(ctx context.Context, userId int64) (models.Balance, error)
}

type OrderService interface {
	AddOrder(ctx context.Context, userId int64, orderNumber string) error
	GetUserOrders(ctx context.Context, userId int64) ([]models.Order, error)
}

type WithdrawalService interface {
	Withdraw(ctx context.Context, withdrawal models.Withdrawal, userId int64) error
}
