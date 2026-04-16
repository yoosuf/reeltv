package http

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// Server represents the HTTP server
type Server struct {
	port   int
	router *Router
}

// NewServer creates a new HTTP server
func NewServer(port int, router *Router) *Server {
	return &Server{
		port:   port,
		router: router,
	}
}

// Start starts the HTTP server
func (s *Server) Start() error {
	addr := fmt.Sprintf(":%d", s.port)

	server := &http.Server{
		Addr:         addr,
		Handler:      s.router.GetEngine(),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	fmt.Printf("Server starting on %s\n", addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}

// Shutdown gracefully shuts down the server
func (s *Server) Shutdown(ctx context.Context) error {
	fmt.Println("Server shutting down...")
	return nil
}

// GetRouter returns the router for testing purposes
func (s *Server) GetRouter() *Router {
	return s.router
}
