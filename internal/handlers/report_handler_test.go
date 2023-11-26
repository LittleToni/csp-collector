package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/LittleToni/csp-collector/internal/middlewares"
)

func TestPostReportHandler(t *testing.T) {
	body := strings.NewReader("test report")
	req, err := http.NewRequest(http.MethodPost, "/report", body)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(middlewares.Post(ReportHandler))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
	}

	expected := ``
	actual := rr.Body.String()
	if actual != expected {
		t.Errorf("handler returned wrong body: got %v want %v", actual, expected)
	}
}

func TestGetReportHandler(t *testing.T) {
	body := strings.NewReader("test report")
	req, err := http.NewRequest(http.MethodGet, "/report", body)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(middlewares.Post(ReportHandler))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
	}

	expected := `Method not Allowed!`
	actual := strings.TrimSpace(rr.Body.String())
	if actual != expected {
		t.Errorf("handler returned wrong body: got %v want %v", actual, expected)
	}
}

func TestPostReportOnlyHandler(t *testing.T) {
	body := strings.NewReader("test report")
	req, err := http.NewRequest(http.MethodPost, "/report-only", body)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(middlewares.Post(ReportOnlyHandler))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
	}

	expected := ``
	actual := rr.Body.String()
	if actual != expected {
		t.Errorf("handler returned wrong body: got %v want %v", actual, expected)
	}
}

func TestGetReportOnlyHandler(t *testing.T) {
	body := strings.NewReader("test report")
	req, err := http.NewRequest(http.MethodGet, "/report-only", body)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(middlewares.Post(ReportOnlyHandler))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
	}

	expected := `Method not Allowed!`
	actual := strings.TrimSpace(rr.Body.String())
	if actual != expected {
		t.Errorf("handler returned wrong body: got %v want %v", actual, expected)
	}
}
