package middlewares

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPostMiddleware(t *testing.T) {
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	handler := Post(next)

	t.Run("Accept POST Method", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/", nil)

		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status OK; got %v", rr.Code)
		}
	})

	t.Run("Reject Non-POST Method", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/", nil)

		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusMethodNotAllowed {
			t.Errorf("expected status MethodNotAllowed; got %v", rr.Code)
		}

		actual := strings.TrimSpace(rr.Body.String())
		expect := "Method not Allowed!"
		if actual != expect {
			t.Errorf("expect body %v; got %v", expect, actual)
		}
	})
}
