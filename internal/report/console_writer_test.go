package report

import (
	"testing"
)

func TestConsoleWriter_Write(t *testing.T) {
	writer := &ConsoleWriter{}

	err := writer.Write("test report")
	if err != nil {
		t.Errorf("WriteReport returned an error: %v", err)
	}
}
