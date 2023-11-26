package middlewares

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCSPContentTypeMiddleware(t *testing.T) {
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	handler := CSPContentType(next)

	t.Run("Accept Content-Type: application/csp-report", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set("Content-Type", "application/csp-report")

		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status OK; got %v", rr.Code)
		}
	})

	t.Run("Reject Content-Type: application/json", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Content-Type", "application/json")

		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusUnsupportedMediaType {
			t.Errorf("expected status UnsupportedMediaType; got %v", rr.Code)
		}

		actual := strings.TrimSpace(rr.Body.String())
		expect := "Invalid Content-Type!"
		if actual != expect {
			t.Errorf("expect body %v; got %v", expect, actual)
		}
	})
}
