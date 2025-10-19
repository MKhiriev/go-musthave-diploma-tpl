package service

import (
	"encoding/hex"
	"fmt"
	"go-musthave-diploma-tpl/internal/config"
	"go-musthave-diploma-tpl/internal/logger"
	"go-musthave-diploma-tpl/internal/store"
	"go-musthave-diploma-tpl/internal/utils"
	"go-musthave-diploma-tpl/models"
)

type authService struct {
	userRepository store.UserRepository
	hashKey        string

	logger *logger.Logger
}

func NewAuthService(userRepository store.UserRepository, cfg *config.Auth, logger *logger.Logger) AuthService {
	return &authService{
		userRepository: userRepository,
		hashKey:        cfg.PasswordHashKey,
		logger:         logger,
	}
}

func (a *authService) RegisterUser(user models.User) error {
	if user.Login == "" || user.Password == "" {
		a.logger.Error().Any("user", user).Msg("invalid user data provided")
		return ErrInvalidDataProvided
	}

	a.hashPassword(&user)
	err := a.userRepository.CreateUser(user)

	if err != nil {
		a.logger.Err(err).Any("user", user).Msg("user creation ended with error")
		return fmt.Errorf("user creation ended with error: %w", err)
	}

	return nil
}

func (a *authService) Login(user models.User) (models.User, error) {
	if user.Login == "" || user.Password == "" {
		a.logger.Error().Any("user", user).Msg("invalid user data provided")
		return models.User{}, ErrInvalidDataProvided
	}

	a.hashPassword(&user)
	foundUser, err := a.userRepository.FindUserByLogin(user)
	if err != nil {
		a.logger.Err(err).Any("user", user).Msg("user search by login failed")
		return models.User{}, fmt.Errorf("user search by login failed: %w", err)
	}

	if foundUser.Password != user.Password {
		a.logger.Err(err).
			Str("login", foundUser.Login).
			Str("input", user.Password).
			Str("actual password", foundUser.Password).
			Msg("wrong password")
		return models.User{}, ErrWrongPassword
	}

	return foundUser, nil
}

func (a *authService) hashPassword(user *models.User) {
	user.Password = hex.EncodeToString(
		utils.Hash([]byte(user.Password), a.hashKey),
	)
}
