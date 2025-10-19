package store

import "go-musthave-diploma-tpl/models"

type UserRepository interface {
	CreateUser(user models.User) error
}
