package main

import (
	"net/http/httptest"
	"strings"
	"testing"
)

type MockReportWriter struct {
	LastReport string
}

func (m *MockReportWriter) WriteReport(report string) error {
	m.LastReport = report
	return nil
}

func TestProcessReport(t *testing.T) {
	mockWriter := &MockReportWriter{}

	body := strings.NewReader("test report")
	request := httptest.NewRequest("POST", "/report", body)
	err := ProcessReport(request, mockWriter)
	if err != nil {
		t.Errorf("ProcessReport returned an error: %v", err)
	}

	if mockWriter.LastReport != "test report" {
		t.Errorf("ProcessReport did not write the correct report: got %v want %v", mockWriter.LastReport, "test report")
	}
}
