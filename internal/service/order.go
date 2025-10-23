package service

import (
	"context"
	"fmt"
	"go-musthave-diploma-tpl/internal/logger"
	"go-musthave-diploma-tpl/internal/store"
	"go-musthave-diploma-tpl/internal/utils"
	"go-musthave-diploma-tpl/models"
)

type orderService struct {
	orderRepository store.OrderRepository
	logger          *logger.Logger
}

func NewOrderService(orderRepository store.OrderRepository, logger *logger.Logger) OrderService {
	return &orderService{
		orderRepository: orderRepository,
		logger:          logger,
	}
}

func (o *orderService) AddOrder(ctx context.Context, userId int64, orderNumber string) error {
	isCorrect, err := utils.LunaCheckString(orderNumber)
	switch {
	case err != nil:
		return fmt.Errorf("%w: %w", ErrInvalidOrderNumber, err)
	case !isCorrect:
		return ErrNotCorrectOrderNumber
	}

	order, err := o.orderRepository.CreateOrderOrGetExisting(ctx, userId, orderNumber)
	switch {
	case order.UserId != userId:
		return ErrOrderWasUploadedByAnotherUser
	case order.UserId == userId:
		return ErrOrderWasUploadedByCurrentUser
	case err != nil:
		return fmt.Errorf("error occured during creation of new order: %w", err)
	}

	return nil
}

func (o *orderService) GetUserOrders(ctx context.Context, userId int64) ([]models.Order, error) {
	orders, err := o.orderRepository.GetOrdersByUserId(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("error occurred during getting order: %w", err)
	}
	if len(orders) == 0 {
		return nil, ErrNoUserOrdersFound
	}

	return orders, nil
}
