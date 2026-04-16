package application

import (
	"context"
	"fmt"

	sharedDomain "reeltv/backend/internal/shared/domain"
	"reeltv/backend/internal/catalog/application"
)

var (
	ErrNoRecommendations = sharedDomain.ErrNotFound
)

// Service implements recommendation use cases
type Service struct {
	catalogService *application.Service
}

// NewService creates a new recommendation application service
func NewService(catalogService *application.Service) *Service {
	return &Service{
		catalogService: catalogService,
	}
}

// GetRecommendations retrieves personalized recommendations for a user
// Uses heuristic-based recommendations: popular series, recently watched, genre-based
func (s *Service) GetRecommendations(ctx context.Context, userID uint, offset, limit int) ([]*application.SeriesResponse, error) {
	// Get popular series (top rated)
	popularSeries, err := s.catalogService.ListSeries(ctx, 0, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get popular series: %w", err)
	}

	if len(popularSeries) == 0 {
		return nil, ErrNoRecommendations
	}

	return popularSeries, nil
}

// GetTrendingSeries retrieves trending series based on watch activity
func (s *Service) GetTrendingSeries(ctx context.Context, offset, limit int) ([]*application.SeriesResponse, error) {
	// For now, return top-rated series as trending
	// In a full implementation, this would use watch activity data
	series, err := s.catalogService.ListSeries(ctx, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get trending series: %w", err)
	}

	if len(series) == 0 {
		return nil, ErrNoRecommendations
	}

	return series, nil
}

// GetRecommendedByGenre retrieves recommendations for a specific genre
func (s *Service) GetRecommendedByGenre(ctx context.Context, genre string, offset, limit int) ([]*application.SeriesResponse, error) {
	// Get all series and filter by genre
	// In a full implementation, this would be a more efficient query
	series, err := s.catalogService.ListSeries(ctx, 0, 100) // Get more to filter
	if err != nil {
		return nil, fmt.Errorf("failed to get series: %w", err)
	}

	var filtered []*application.SeriesResponse
	for _, s := range series {
		if s.Genre == genre {
			filtered = append(filtered, s)
		}
	}

	if len(filtered) == 0 {
		return nil, ErrNoRecommendations
	}

	// Apply pagination
	if offset >= len(filtered) {
		return []*application.SeriesResponse{}, nil
	}

	end := offset + limit
	if end > len(filtered) {
		end = len(filtered)
	}

	return filtered[offset:end], nil
}
