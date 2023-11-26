package handlers

import (
	"net/http"

	"github.com/LittleToni/csp-collector/internal/report"
)

// ReportHandler accepts a csp violation report
func ReportHandler(w http.ResponseWriter, r *http.Request) {
	writer := &report.ConsoleWriter{}

	if err := report.ProcessReport(r, writer); err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ReportOnlyHandler accepts a csp violation report
func ReportOnlyHandler(w http.ResponseWriter, r *http.Request) {
	writer := &report.ConsoleWriter{}

	if err := report.ProcessReport(r, writer); err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
