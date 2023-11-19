package main

import (
	"fmt"
)

type ReportWriter interface {
	WriteReport(report string) error
}

type ConsoleReportWriter struct{}

func (cw *ConsoleReportWriter) WriteReport(report string) error {
	fmt.Println("Report received:", report)
	return nil
}
