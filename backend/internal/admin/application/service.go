package application

import (
	"errors"
)

var (
	ErrUnauthorized = errors.New("unauthorized: admin access required")
)

// Service implements admin use cases
type Service struct {
	// Catalog service reference for admin operations
}

// NewService creates a new admin application service
func NewService() *Service {
	return &Service{}
}

// IsAdmin checks if a user has admin privileges
func (s *Service) IsAdmin(userID uint) bool {
	// For MVP, we'll use a simple check
	// In production, this would check against a roles/permissions table
	// For now, we'll consider user ID 1 as admin
	return userID == 1
}
