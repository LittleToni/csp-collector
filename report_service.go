package main

import (
	"io"
	"net/http"
)

func ProcessReport(r *http.Request, writer ReportWriter) error {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return err
	}

	report := string(body)
	return writer.WriteReport(report)
}
