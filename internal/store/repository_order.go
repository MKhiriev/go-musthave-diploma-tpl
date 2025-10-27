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
	var order models.Order

	err := o.db.Raw(createNewOrderOrReturnExisting, map[string]interface{}{
		"number":      orderNumber,
		"status_name": models.NEW,
		"user_id":     userId,
		"accrual":     0,
	}).Scan(&order).Error

	if err != nil {
		return models.Order{}, fmt.Errorf("DB error: %w", err)
	}

	return order, nil
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
