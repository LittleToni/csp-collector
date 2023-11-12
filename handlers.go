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
