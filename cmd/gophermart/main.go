package main

import (
	"go-musthave-diploma-tpl/internal/config"
	"go-musthave-diploma-tpl/internal/handler"
	"go-musthave-diploma-tpl/internal/logger"
	"go-musthave-diploma-tpl/internal/server"
)

func main() {
	log := logger.NewLogger("gophermart-server")
	cfg := config.GetConfigs()

	handlers := handler.NewHandler(log)
	srv := new(server.Server)

	log.Info().Msg("Server started")
	_ = srv.ServerRun(handlers.Init(), cfg)
}
