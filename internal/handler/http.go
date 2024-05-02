package handler

import (
	"InterviewTest/internal/middlewhare"
	"InterviewTest/internal/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Handler struct {
	Router           httprouter.Router
	ServiceContainer *service.Service
}

func NewHandler(serviceContainer *service.Service) *Handler {
	r := httprouter.New()
	calculate := NewCalculator()
	r.GET("/calculate", middlewhare.RequestParamCheckMiddleware(calculate.Calculate()))

	return &Handler{
		Router:           *r,
		ServiceContainer: serviceContainer,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Router.ServeHTTP(w, r)
}
