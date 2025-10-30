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
	"time"
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

	return order, nil
}

func (o *orderService) GetAccruals(ctx context.Context) error {
	unprocessedOrders, err := o.orderRepository.GetOrdersByStatuses(ctx, models.NEW, models.PROCESSING)
	if err != nil {
		return fmt.Errorf("error while getting orders by status: %w", err)
	}

	if len(unprocessedOrders) == 0 {
		return ErrNoOrdersForUpdate
	}

	ordersToUpdate := make([]models.Order, 0, len(unprocessedOrders))
	for _, order := range unprocessedOrders {
		accrual, accrualErr := o.accrualAdapter.GetOrderAccrual(ctx, order.Number)
		if accrualErr != nil {
			return fmt.Errorf("error during getting order accrual: %w", accrualErr)
		}
		if accrual.StatusText != order.StatusText {
			ordersToUpdate = append(ordersToUpdate, accrual)
		}
	}

	if len(ordersToUpdate) == 0 {
		return ErrNoOrdersToUpdate
	}

	err = o.orderRepository.UpdateOrders(ctx, ordersToUpdate...)
	if err != nil {
		return fmt.Errorf("error during updating orders: %w", err)
	}

	return nil
}

func (o *orderService) UpdateOrder(ctx context.Context, order models.Order) error {
	o.logger.Debug().Str("func", "worker.UpdateOrder").Msg("worker")
	err := o.orderRepository.UpdateOrders(ctx, order)
	if err != nil {
		return fmt.Errorf("error during updating orders: %w", err)
	}

	return nil
}

func (o *orderService) RunAccrualUpdate() error {
	ticker := time.NewTicker(time.Duration(1) * time.Second)

	jobs := o.generateWork(ticker)

	o.logger.Debug().Str("func", "Run").Msg("creating workers")
	o.withWorkers(func() {
		o.worker(jobs)
	}, 3)
	o.logger.Debug().Str("func", "Run").Msg("workers created")

	select {}
}

func (o *orderService) generateWork(ticker *time.Ticker) chan models.Order {
	jobs := make(chan models.Order)

	go func() {
		for {
			select {
			case <-ticker.C:
				ctx := context.Background()
				_ = o.getOrders(ctx, jobs)
			}
		}
	}()
	return jobs
}

func (o *orderService) getOrders(ctx context.Context, jobs chan models.Order) error {
	unprocessedOrders, err := o.orderRepository.GetOrdersByStatuses(ctx, models.NEW, models.PROCESSING)
	if err != nil {
		err = fmt.Errorf("error while getting orders by status: %w", err)
		o.logger.Err(err).Send()
		return err
	}

	if len(unprocessedOrders) == 0 {
		err = ErrNoOrdersForUpdate
		o.logger.Err(err).Send()
		return err
	}

	for i, order := range unprocessedOrders {
		accrual, accrualErr := o.accrualAdapter.GetOrderAccrual(ctx, order.Number)
		if accrualErr != nil {
			err = fmt.Errorf("error during getting order accrual: %w", accrualErr)
			o.logger.Err(err).Send()
			return err
		}
		o.logger.Debug().Int("count", i).Any("accrual to check", accrual).Msg("SEND ACCRUAL JOB")
		jobs <- accrual
	}

	return nil
}

func (o *orderService) worker(jobs <-chan models.Order) {
	for order := range jobs {
		o.logger.Debug().Any("accrual to check", order).Msg("WORKER received")
		ctx := context.Background()
		_ = o.UpdateOrder(ctx, order)
	}
}

func (o *orderService) withWorkers(fn func(), count int64) {
	for i := range count {
		o.logger.Debug().Str("func", "withWorkers").Msgf("creating worker #%d", i)
		go fn()
		o.logger.Debug().Msgf("worker#%d is created", i)
	}
}
