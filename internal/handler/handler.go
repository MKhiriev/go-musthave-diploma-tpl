package handler

import (
	"net/http"

	"github.com/rs/zerolog"
)

type Handlers struct {
	logger *zerolog.Logger
}

func (h *Handlers) order(writer http.ResponseWriter, request *http.Request) {
	// TODO: implement me!
}

func (h *Handlers) getOrders(writer http.ResponseWriter, request *http.Request) {
	// TODO: implement me!
}

func (h *Handlers) getBalance(writer http.ResponseWriter, request *http.Request) {
	// TODO: implement me!
}

func (h *Handlers) withdraw(writer http.ResponseWriter, request *http.Request) {
	// TODO: implement me!
}

func (h *Handlers) getWithdrawals(writer http.ResponseWriter, request *http.Request) {
	// TODO: implement me!
}

func NewHandler(logger *zerolog.Logger) Handlers {
	return Handlers{logger: logger}
}
