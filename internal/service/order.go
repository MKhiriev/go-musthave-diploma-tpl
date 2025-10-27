package service

import (
	"context"
	"errors"
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
		return fmt.Errorf("error in order number: %w", ErrInvalidOrderNumber)
	case !isCorrect:
		return ErrNotCorrectOrderNumber
	}

	order, err := o.orderRepository.CreateOrderOrGetExisting(ctx, userId, orderNumber)
	switch {
	case errors.Is(err, store.ErrOrderWasNotCreated) && order.UserId != userId:
		return ErrOrderWasUploadedByAnotherUser
	case errors.Is(err, store.ErrOrderWasNotCreated) && order.UserId == userId:
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

func (o *orderService) GetOrder(ctx context.Context, orderNumber string) (models.Order, error) {
	order, err := o.orderRepository.GetOrderByNumber(ctx, orderNumber)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrOrderNotFound):
			return models.Order{}, ErrOrderNotRegistered
		default:
			return models.Order{}, fmt.Errorf("error occurred during getting order: %w", err)
		}
	}

	if order.StatusText == models.PROCESSED || order.StatusText == models.INVALID {
		return order, nil
	}

	//status, err := o.accrualAdapter.GetOrderAccrual(ctx)

	return order, nil
}
