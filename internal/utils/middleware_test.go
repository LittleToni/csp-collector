package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMiddlewareNoMiddleware(t *testing.T) {
	originalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	wrappedHandler := Middleware(originalHandler)

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/", nil)

	wrappedHandler.ServeHTTP(recorder, request)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
