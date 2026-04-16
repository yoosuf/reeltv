package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"reeltv/backend/internal/app"
	"reeltv/backend/internal/config"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize application
	application, err := app.New(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	// Start application in a goroutine
	go func() {
		if err := application.Run(); err != nil {
			log.Fatalf("Failed to run application: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Graceful shutdown
	if err := application.Shutdown(); err != nil {
		log.Fatalf("Failed to shutdown application: %v", err)
	}

	os.Exit(0)
}
