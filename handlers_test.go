package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetHealthcheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthcheckHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"status":"ok"}`
	actual := strings.TrimSpace(rr.Body.String())
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func TestPostReportHandler(t *testing.T) {
	body := strings.NewReader("test report")
	req, err := http.NewRequest("POST", "/report", body)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Post(ReportHandler))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
	}

	expected := ``
	actual := rr.Body.String()
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func TestGetReportHandler(t *testing.T) {
	body := strings.NewReader("test report")
	req, err := http.NewRequest("GET", "/report", body)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Post(ReportHandler))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
	}

	expected := `Method not Allowed!`
	actual := strings.TrimSpace(rr.Body.String())
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func TestPostReportOnlyHandler(t *testing.T) {
	body := strings.NewReader("test report")
	req, err := http.NewRequest("POST", "/report-only", body)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Post(ReportOnlyHandler))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
	}

	expected := ``
	actual := rr.Body.String()
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func TestGetReportOnlyHandler(t *testing.T) {
	body := strings.NewReader("test report")
	req, err := http.NewRequest("GET", "/report-only", body)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Post(ReportOnlyHandler))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
	}

	expected := `Method not Allowed!`
	actual := strings.TrimSpace(rr.Body.String())
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}
