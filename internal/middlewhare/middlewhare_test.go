package middlewhare

import (
	"bytes"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestParamCheckMiddleware_validJSON(t *testing.T) {
	jsonStr := []byte(`{"a": 10, "b": 5}`)
	req, err := http.NewRequest("GET", "/calculate", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	// Підготовка рекордера відповіді
	rr := httptest.NewRecorder()

	// Виклик middleware з тестовим обробником
	middlewhareTest := RequestParamCheckMiddleware(testHandler)
	middlewhareTest(rr, req, nil)

	// Перевірка, чи статус код відповіді є очікуваним
	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
	}
}

func TestRequestParamCheckMiddleware_invalidJSON(t *testing.T) {
	jsonStr := []byte(`{"a": "invalid", "b": 5}`)
	req, err := http.NewRequest("GET", "/calculate", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	//рекордер
	rr := httptest.NewRecorder()

	middleware := RequestParamCheckMiddleware(testHandler)
	middleware(rr, req, nil)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusBadRequest)
	}
}

func TestRequestParamCheckMiddleware_negativeNumbers(t *testing.T) {
	jsonStr := []byte(`{"a": -10, "b": 5}`)
	req, err := http.NewRequest("GET", "/calculate", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	middleware := RequestParamCheckMiddleware(testHandler)
	middleware(rr, req, nil)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusBadRequest)
	}
}

func testHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.WriteHeader(http.StatusOK)
}
