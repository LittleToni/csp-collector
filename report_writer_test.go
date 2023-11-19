package main

import (
	"testing"
)

func TestConsoleReportWriter_WriteReport(t *testing.T) {
	writer := &ConsoleReportWriter{}

	err := writer.WriteReport("test report")
	if err != nil {
		t.Errorf("WriteReport returned an error: %v", err)
	}
}
