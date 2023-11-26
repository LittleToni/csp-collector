package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/LittleToni/csp-collector/internal/config"
	"github.com/LittleToni/csp-collector/internal/handlers"
	"github.com/LittleToni/csp-collector/internal/middlewares"
	"github.com/LittleToni/csp-collector/internal/utils"
)

func main() {
	// Load configration settings.
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Could not load configuration:", err)
	}

	// Register healthcheck handlers.
	http.HandleFunc("/health", handlers.HealthcheckHandler)
	http.HandleFunc("/healthcheck", handlers.HealthcheckHandler)

	// Register report handlers.
	http.HandleFunc("/report", utils.Middleware(handlers.ReportHandler, middlewares.Post, middlewares.CSPContentType))
	http.HandleFunc("/report-only", utils.Middleware(handlers.ReportOnlyHandler, middlewares.Post, middlewares.CSPContentType))

	// Start the server.
	fmt.Printf("Server started at port %s! \n", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}
