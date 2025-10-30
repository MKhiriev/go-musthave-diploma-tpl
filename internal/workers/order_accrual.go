package workers

import (
	"context"
	"go-musthave-diploma-tpl/internal/adapter"
	"go-musthave-diploma-tpl/internal/logger"
	"go-musthave-diploma-tpl/internal/service"
	"go-musthave-diploma-tpl/models"
	"time"
)

type orderAccrualWorker struct {
	orderService    service.OrderService
	accrualAdapter  adapter.AccrualAdapter
	jobInterval     time.Duration
	goroutinesCount int64

	logger *logger.Logger
}

func NewOrderAccrualWorker(orderService service.OrderService, accrualAdapter adapter.AccrualAdapter, interval time.Duration, logger *logger.Logger) Worker {
	return &orderAccrualWorker{
		orderService:    orderService,
		accrualAdapter:  accrualAdapter,
		jobInterval:     interval,
		goroutinesCount: 3,
		logger:          logger,
	}
}

func (o *orderAccrualWorker) Run() {
	o.logger.Info().Str("worker", "orderAccrualWorker").Msg("worker started")

	go func() {
		jobs := o.generateWork(time.NewTicker(o.jobInterval))

		o.withWorkers(func() {
			o.worker(jobs)
		}, o.goroutinesCount)
	}()
}

func (o *orderAccrualWorker) generateWork(ticker *time.Ticker) chan []models.Order {
	jobs := make(chan []models.Order)

	go func() {
		for {
			select {
			case <-ticker.C:
				orders, err := o.orderService.GetOrdersForAccrual(context.Background())
				if err != nil {
					o.logger.Err(err).Msg("error occurred during generating work")
					return
				}

				jobs <- orders
			}
		}
	}()

	return jobs
}

func (o *orderAccrualWorker) worker(jobs <-chan []models.Order) {
	for ordersBatch := range jobs {
		o.logger.Debug().Any("accrual for batch", ordersBatch).Msg("WORKER received batch")

		_ = o.orderService.UpdateOrders(context.Background(), ordersBatch...)

		o.logger.Debug().Msg("WORKER finished batch")
	}
}

func (o *orderAccrualWorker) withWorkers(fn func(), count int64) {
	for i := range count {
		o.logger.Debug().Str("func", "withWorkers").Msgf("creating worker #%d", i)

		go fn()

		o.logger.Debug().Msgf("worker#%d is created", i)
	}
}
