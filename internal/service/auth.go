package service

import (
	"context"
	"fmt"
	"go-musthave-diploma-tpl/internal/config"
	"go-musthave-diploma-tpl/internal/logger"
	"go-musthave-diploma-tpl/internal/store"
	"go-musthave-diploma-tpl/internal/utils"
	"go-musthave-diploma-tpl/models"
	"time"
)

type authService struct {
	userRepository store.UserRepository
	hashKey        string
	tokenSignKey   string
	tokenIssuer    string
	tokenDuration  time.Duration

	logger *logger.Logger
}

func NewAuthService(userRepository store.UserRepository, cfg *config.Auth, logger *logger.Logger) AuthService {
	return &authService{
		userRepository: userRepository,
		hashKey:        cfg.PasswordHashKey,
		tokenSignKey:   cfg.TokenSignKey,
		tokenIssuer:    cfg.TokenIssuer,
		tokenDuration:  cfg.TokenDuration,
		logger:         logger,
	}
}

func (a *authService) RegisterUser(ctx context.Context, user models.User) error {
	if user.Login == "" || user.Password == "" {
		a.logger.Error().Any("user", user).Msg("invalid user data provided")
		return ErrInvalidDataProvided
	}

	a.hashPassword(&user)
	err := a.userRepository.CreateUser(ctx, user)

	if err != nil {
		a.logger.Err(err).Any("user", user).Msg("user creation ended with error")
		return fmt.Errorf("user creation ended with error: %w", err)
	}

	return nil
}

func (a *authService) Login(ctx context.Context, user models.User) (models.User, error) {
	if user.Login == "" || user.Password == "" {
		a.logger.Error().Any("user", user).Msg("invalid user data provided")
		return models.User{}, ErrInvalidDataProvided
	}

	a.hashPassword(&user)
	foundUser, err := a.userRepository.FindUserByLogin(ctx, user)
	if err != nil {
		a.logger.Err(err).Any("user", user).Msg("user search by login failed")
		return models.User{}, fmt.Errorf("user search by login failed: %w", err)
	}

	if foundUser.Password != user.Password {
		a.logger.Err(err).
			Int64("id", foundUser.UserId).
			Str("login", foundUser.Login).
			Str("input", user.Password).
			Str("actual password", foundUser.Password).
			Msg("wrong password")
		return models.User{}, ErrWrongPassword
	}

	return foundUser, nil
}

func (a *authService) CreateToken(ctx context.Context, user models.User) (models.Token, error) {
	token, err := utils.GenerateJWTToken(a.tokenIssuer, user.UserId, a.tokenDuration, a.tokenSignKey)
	if err != nil {
		return models.Token{}, fmt.Errorf("error creating JWT token: %w", err)
	}

	return token, nil
}

func (a *authService) ParseToken(ctx context.Context, tokenString string) (models.Token, error) {
	token, err := utils.ValidateAndParseJWTToken(tokenString, a.tokenSignKey, a.tokenIssuer)
	if err != nil {
		return models.Token{}, fmt.Errorf("error parsing JWT token: %w", err)
	}

	return token, nil
}

func (a *authService) hashPassword(user *models.User) {
	user.Password = utils.HashString(user.Password, a.hashKey)
}
