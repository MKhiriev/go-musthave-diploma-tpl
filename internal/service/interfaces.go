package service

import (
	"context"
	"go-musthave-diploma-tpl/models"
)

type AuthService interface {
	RegisterUser(ctx context.Context, user models.User) error
	Login(ctx context.Context, user models.User) (models.User, error)
}
