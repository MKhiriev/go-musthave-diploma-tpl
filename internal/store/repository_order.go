package store

import (
	"context"
	"errors"
	"fmt"
	"go-musthave-diploma-tpl/internal/logger"
	"go-musthave-diploma-tpl/models"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type orderRepository struct {
	db     *gorm.DB
	logger *logger.Logger
}

func NewOrderRepository(db *gorm.DB, logger *logger.Logger) OrderRepository {
	logger.Debug().Msg("OrderRepository created")
	return &orderRepository{
		db:     db,
		logger: logger,
	}
}

func (o *orderRepository) CreateOrderOrGetExisting(ctx context.Context, userId int64, orderNumber string) (models.Order, error) {
	var order OrderWithFlag
	err := o.db.WithContext(ctx).Raw(createNewOrderOrReturnExisting, map[string]interface{}{
		"number":      orderNumber,
		"status_name": models.NEW,
		"user_id":     userId,
		"accrual":     0,
	}).Scan(&order).Error

	if err != nil {
		return models.Order{}, fmt.Errorf("DB error: %w", err)
	}
	if !order.IsNew {
		return order.Order, ErrOrderWasNotCreated
	}

	return order.Order, nil
}

func (o *orderRepository) GetOrdersByUserId(ctx context.Context, userId int64) ([]models.Order, error) {
	var orders []models.Order

	err := o.db.WithContext(ctx).
		Select("number, s.name as status, accrual, uploaded_at").
		Joins("LEFT JOIN statuses s ON orders.status_id = s.status_id").
		Where("user_id = ?", userId).
		Find(&orders).Error
	if err != nil {
		return nil, fmt.Errorf("DB error: %w", err)
	}

	return orders, nil
}

func (o *orderRepository) GetOrderByNumber(ctx context.Context, orderNumber string) (models.Order, error) {
	var order models.Order

	err := o.db.WithContext(ctx).
		Where("number = ?", orderNumber).
		Find(&order).Error

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case pgerrcode.NoDataFound:
			return models.Order{}, ErrOrderNotFound
		default:
			return models.Order{}, fmt.Errorf("unexpected DB error: %w", err)
		}
	}

	return order, nil
}

func (o *orderRepository) GetOrdersByStatuses(ctx context.Context, statuses ...string) ([]models.Order, error) {
	if len(statuses) == 0 {
		return nil, ErrNoStatusesPassed
	}

	var orders []models.Order
	err := o.db.WithContext(ctx).
		Select("orders.*, s.name AS status").
		Joins("JOIN statuses s ON s.status_id = orders.status_id").
		Where("s.name IN ?", statuses).
		Find(&orders).Error

	if err != nil {
		return nil, fmt.Errorf("failed to get orders by statuses: %w", err)
	}

	return orders, nil
}

func (o *orderRepository) UpdateOrders(ctx context.Context, orders ...models.Order) error {
	type result struct {
		OrdersUpdated  int
		BalanceUpdated int
	}

	err := o.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, order := range orders {
			queryRes := result{}
			res := tx.Raw(updateOrderAccrualAndBalance, map[string]interface{}{
				"status_name":  order.StatusText,
				"accrual":      order.Accrual,
				"order_number": order.Number,
			}).Scan(&queryRes)

			if res.Error != nil {
				return fmt.Errorf("failed to update order %s: %w", order.Number, res.Error)
			}

			if queryRes.OrdersUpdated == 0 || queryRes.BalanceUpdated == 0 {
				return fmt.Errorf("order %s not found", order.Number)
			}
		}

		return nil
	})

	return err
}
