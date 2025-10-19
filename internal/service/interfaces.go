package service

import "go-musthave-diploma-tpl/models"

type AuthService interface {
	RegisterUser(user models.User) error
	Login(user models.User) (models.User, error)
}
