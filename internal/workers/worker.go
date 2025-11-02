package workers

import (
	"go-musthave-diploma-tpl/internal/adapter"
	"go-musthave-diploma-tpl/internal/config"
	"go-musthave-diploma-tpl/internal/logger"
	"go-musthave-diploma-tpl/internal/service"
)

type Workers struct {
	workers []Worker
}

func NewWorkers(services *service.Services, adapters *adapter.Adapters, cfg *config.Workers, logger *logger.Logger) *Workers {
	workers := &Workers{}
	workers.workers = append(workers.workers, NewOrderAccrualWorker(services.OrderService, adapters.AccrualAdapter, cfg.OrderAccrualInterval, logger))

	return workers
}

func (w *Workers) Run() {
	for _, worker := range w.workers {
		worker.Run()
	}
}
