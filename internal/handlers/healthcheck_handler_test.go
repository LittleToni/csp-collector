package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealthcheckHandler(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthcheckHandler)
	handler.ServeHTTP(rr, req)

	t.Run("Healthcheck return expected status code", func(t *testing.T) {
		if rr.Code != http.StatusOK {
			t.Errorf("handler returned wrong status code: %v; got %v", rr.Code, http.StatusOK)
		}
	})

	t.Run("Healthcheck return expected body", func(t *testing.T) {
		actual := strings.TrimSpace(rr.Body.String())
		expect := `{"status":"ok"}`
		if actual != expect {
			t.Errorf("handler returned wrong body: %v; got %v", actual, expect)
		}
	})
}
