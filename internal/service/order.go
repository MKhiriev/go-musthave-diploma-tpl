package service

import (
	"context"
	"errors"
	"fmt"
	"go-musthave-diploma-tpl/internal/adapter"
	"go-musthave-diploma-tpl/internal/logger"
	"go-musthave-diploma-tpl/internal/store"
	"go-musthave-diploma-tpl/internal/utils"
	"go-musthave-diploma-tpl/models"
)

type orderService struct {
	orderRepository store.OrderRepository
	accrualAdapter  adapter.AccrualAdapter

	logger *logger.Logger
}

func NewOrderService(orderRepository store.OrderRepository, accrualAdapter adapter.AccrualAdapter, logger *logger.Logger) OrderService {
	return &orderService{
		orderRepository: orderRepository,
		accrualAdapter:  accrualAdapter,
		logger:          logger,
	}
}

func (o *orderService) AddOrder(ctx context.Context, userID int64, orderNumber string) error {
	isCorrect, err := utils.LunaCheckString(orderNumber)
	switch {
	case err != nil:
		return fmt.Errorf("error in order number: %w", ErrInvalidOrderNumber)
	case !isCorrect:
		return ErrNotCorrectOrderNumber
	}

	order, err := o.orderRepository.CreateOrderOrGetExisting(ctx, userID, orderNumber)
	switch {
	case errors.Is(err, store.ErrOrderWasNotCreated) && order.UserID != userID:
		return ErrOrderWasUploadedByAnotherUser
	case errors.Is(err, store.ErrOrderWasNotCreated) && order.UserID == userID:
		return ErrOrderWasUploadedByCurrentUser
	case err != nil:
		return fmt.Errorf("error occured during creation of new order: %w", err)
	}

	return nil
}

func (o *orderService) GetUserOrders(ctx context.Context, userID int64) ([]models.Order, error) {
	orders, err := o.orderRepository.GetOrdersByUserID(ctx, userID)
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

	return order, nil
}

func (o *orderService) UpdateOrders(ctx context.Context, order ...models.Order) error {
	if len(order) == 0 {
		return ErrNoOrdersToUpdate
	}

	err := o.orderRepository.UpdateOrders(ctx, order...)
	if err != nil {
		return fmt.Errorf("error during updating orders: %w", err)
	}

	return nil
}

func (o *orderService) GetOrdersForAccrual(ctx context.Context) ([]models.Order, error) {
	unprocessedOrders, err := o.orderRepository.GetOrdersByStatuses(ctx, models.NEW, models.PROCESSING)
	if err != nil {
		err = fmt.Errorf("error while getting orders by status: %w", err)
		o.logger.Err(err).Send()
		return nil, err
	}

	if len(unprocessedOrders) == 0 {
		err = ErrNoOrdersForUpdate
		o.logger.Err(err).Send()
		return nil, err
	}

	ordersForAccrual := make([]models.Order, 0, len(unprocessedOrders))
	for _, order := range unprocessedOrders {
		accrual, accrualErr := o.accrualAdapter.GetOrderAccrual(ctx, order.Number)
		if accrualErr != nil {
			err = fmt.Errorf("error during getting order accrual: %w", accrualErr)
			o.logger.Err(err).Send()
			return nil, err
		}

		ordersForAccrual = append(ordersForAccrual, accrual)

	}

	return ordersForAccrual, nil
}
