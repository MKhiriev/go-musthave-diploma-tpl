package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (h *Handler) Init() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer, h.logging)

	// routes without authorization
	router.Group(func(r chi.Router) {
		r.Post("/api/user/register", h.register) // регистрация пользователя
		r.Post("/api/user/login", h.login)       // аутентификация пользователя
	})

	// routes with required authorization
	router.Group(func(r chi.Router) {
		r.Use(h.auth)

		r.Post("/api/user/orders", h.order)              // загрузка пользователем номера заказа для расчёта
		r.Get("/api/user/orders", h.getOrders)           // получение списка загруженных пользователем номеров заказов, статусов их обработки и информации о начислениях
		r.Get("/api/user/balance", h.getBalance)         // получение текущего баланса счёта баллов лояльности пользователя
		r.Post("/api/user/balance/withdraw", h.withdraw) // запрос на списание баллов с накопительного счёта в счёт оплаты нового заказа
		r.Get("/api/user/withdrawals", h.getWithdrawals) // получение информации о выводе средств с накопительного счёта пользователем
	})

	router.MethodNotAllowed(CheckHTTPMethod(router))

	return router
}
