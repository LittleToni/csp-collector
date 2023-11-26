package report

import "fmt"

// ConsoleWriter is a struct that implements the Writer interface.
// It provides a method to write reports to the console (standard output).
// This type can be used in contexts where reports need to be outputted
// to the console rather than stored or sent elsewhere.
type ConsoleWriter struct{}

// Write implements the Write method of the Writer interface for ConsoleWriter.
// It outputs the provided report to the console along with a prefix message.
func (cw *ConsoleWriter) Write(report string) error {
	_, err := fmt.Println("Report received:", report)
	return err
}
