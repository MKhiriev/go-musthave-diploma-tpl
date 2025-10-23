package main

import (
	"go-musthave-diploma-tpl/internal/config"
	"go-musthave-diploma-tpl/internal/handler"
	"go-musthave-diploma-tpl/internal/logger"
	"go-musthave-diploma-tpl/internal/server"
	"go-musthave-diploma-tpl/internal/service"
	"go-musthave-diploma-tpl/internal/store"
)

func main() {
	log := logger.NewLogger("gophermart-server")
	cfg, err := config.GetStructuredConfig()
	if err != nil {
		log.Err(err).Any("configs", cfg).Msg("invalid configs provided!")
		return
	}
	log.Info().Any("configs", cfg).Msg("program started")

	conn, err := store.NewPostgresConnection(&cfg.DB, log)
	if err != nil {
		log.Err(err).Msg("error during establishing connection to database")
		return
	}

	db, err := store.NewDatabase(conn, log)
	if err != nil {
		log.Err(err).Msg("error during creating database db")
		return
	}

	services := service.NewServices(db, &cfg.Auth, log)

	myHandler := handler.NewHandler(services, log)
	srv := new(server.Server)

	log.Info().Msg("Server started")
	_ = srv.ServerRun(myHandler.Init(cfg.Server.RequestTimeout), &cfg.Server)
}
