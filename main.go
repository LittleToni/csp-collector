package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Load configration settings.
	config, err := LoadConfig()
	if err != nil {
		log.Fatal("Could not load configuration:", err)
	}

	// Register handlers for various endpoints
	http.HandleFunc("/health", HealthcheckHandler)
	http.HandleFunc("/healthcheck", HealthcheckHandler)

	http.HandleFunc("/report", ReportHandler)
	http.HandleFunc("/report-only", ReportOnlyHandler)

	// Start the server.
	fmt.Printf("Server started at port %s!", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}
