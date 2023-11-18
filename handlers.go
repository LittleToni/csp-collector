package main

import (
	"encoding/json"
	"net/http"
)

// HealthcheckHandler responds to health check requests.
func HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := map[string]string{
		"status": "ok",
	}

	json.NewEncoder(w).Encode(response)
}

// ReportHandler accepts a csp violation report
func ReportHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Todo: stdout the report

	w.WriteHeader(http.StatusNoContent)
}

// ReportOnlyHandler accepts a csp violation report
func ReportOnlyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Todo: stdout the report

	w.WriteHeader(http.StatusNoContent)
}
