package store

import (
	"context"
	"go-musthave-diploma-tpl/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user models.User) error
	FindUserByLogin(ctx context.Context, user models.User) (models.User, error)
}
