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

	// Register healthcheck handlers.
	http.HandleFunc("/health", HealthcheckHandler)
	http.HandleFunc("/healthcheck", HealthcheckHandler)

	// Register report handlers.
	http.HandleFunc("/report", Post(ReportHandler))
	http.HandleFunc("/report-only", Post(ReportOnlyHandler))

	// Start the server.
	fmt.Printf("Server started at port %s! \n", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}
