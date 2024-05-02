package middlewhare

import (
	"InterviewTest/internal/model"
	"context"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RequestParamCheckMiddleware(handler httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		decoder := json.NewDecoder(r.Body)
		var request model.CalculateRequest

		if err := decoder.Decode(&request); err != nil {
			http.Error(w, `{"error": "Invalid JSON"}`, http.StatusBadRequest)
			return
		}

		if request.A <= 0 || request.B <= 0 {
			http.Error(w, `{"error": "Incorrect input"}`, http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), "calculateRequest", &request)

		handler(w, r.WithContext(ctx), ps)
	}
}
