package main

import (
	"log"
	"net/http"
	"os"

	"github.com/securerent/risk-assessment-service/internal/handler"
	"github.com/securerent/risk-assessment-service/internal/service"
)

func main() {
	// Initialize services
	riskService := service.NewRiskService()
	
	// Initialize handlers
	riskHandler := handler.NewRiskHandler(riskService)

	// Setup routes
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/assess", riskHandler.Assess)
	mux.HandleFunc("/health", riskHandler.HealthCheck)

	// Configure server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Risk Assessment Service starting on port %s...", port)
	
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
