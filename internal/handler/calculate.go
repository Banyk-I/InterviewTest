package handler

import (
	"InterviewTest/internal/model"
	"InterviewTest/internal/response"
	"InterviewTest/internal/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Calculator struct {
	Service *service.Calculator
}

func NewCalculator() *Calculator {
	return &Calculator{
		Service: service.NewCalculator(),
	}
}

type CalculateResponse struct {
	FactorialA int `json:"factorial_a"`
	FactorialB int `json:"factorial_b"`
}

func (c *Calculator) Calculate() func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		request, ok := r.Context().Value("calculateRequest").(*model.CalculateRequest)
		if !ok {
			http.Error(w, `{"error": "Failed to get request data"}`, http.StatusInternalServerError)
			return
		}

		a, b := c.Service.CalculateFactorials(request.A, request.B)

		result := CalculateResponse{
			FactorialA: a,
			FactorialB: b,
		}

		response.RespondJSON(w, result)
	}
}
