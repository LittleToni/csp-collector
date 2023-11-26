package report

import (
	"io"
	"net/http"
)

// ProcessReport reads the body of an HTTP request and writes it to a provided Writer.
// This function is intended to process reports sent via HTTP requests, for instance,
// in a scenario where reports are received as request bodies and need to be written
// to a storage medium or processed further.
func ProcessReport(r *http.Request, w Writer) error {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return err
	}

	report := string(body)
	return w.Write(report)
}
