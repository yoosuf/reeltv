package application

import (
	"context"
	"fmt"
	"time"

	sharedDomain "reeltv/backend/internal/shared/domain"
	"reeltv/backend/internal/watchprogress/domain"
	"reeltv/backend/pkg/utils"
)

var (
	ErrWatchProgressNotFound = sharedDomain.ErrNotFound
)

// Service implements watch progress use cases
type Service struct {
	watchProgressRepo domain.WatchProgressRepository
}

// NewService creates a new watch progress application service
func NewService(watchProgressRepo domain.WatchProgressRepository) *Service {
	return &Service{
		watchProgressRepo: watchProgressRepo,
	}
}

// UpdateWatchProgress updates or creates watch progress for an episode
func (s *Service) UpdateWatchProgress(ctx context.Context, userID uint, req *UpdateWatchProgressRequest) (*WatchProgressResponse, error) {
	// Check if existing progress exists
	existing, err := s.watchProgressRepo.FindByUserAndEpisode(ctx, userID, req.EpisodeID)
	if err == nil && existing != nil {
		// Update existing progress
		existing.Duration = req.Duration
		existing.Percentage = req.Percentage
		existing.Completed = req.Completed
		existing.WatchedAt = utils.NowUTC()

		if err := s.watchProgressRepo.Update(ctx, existing); err != nil {
			return nil, fmt.Errorf("failed to update watch progress: %w", err)
		}

		return s.toWatchProgressResponse(existing), nil
	}

	// Create new progress
	progress := &domain.WatchProgress{
		UUIDEntity: sharedDomain.UUIDEntity{UUID: utils.GenerateUUID()},
		UserID:     userID,
		EpisodeID:  req.EpisodeID,
		WatchedAt:  utils.NowUTC(),
		Duration:   req.Duration,
		Completed:  req.Completed,
		Percentage: req.Percentage,
	}

	if err := s.watchProgressRepo.Create(ctx, progress); err != nil {
		return nil, fmt.Errorf("failed to create watch progress: %w", err)
	}

	return s.toWatchProgressResponse(progress), nil
}

// GetWatchProgressByUser retrieves watch progress for a user
func (s *Service) GetWatchProgressByUser(ctx context.Context, userID uint, offset, limit int) ([]*WatchProgressResponse, error) {
	progressList, err := s.watchProgressRepo.FindByUserID(ctx, userID, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get watch progress: %w", err)
	}

	responses := make([]*WatchProgressResponse, len(progressList))
	for i, progress := range progressList {
		responses[i] = s.toWatchProgressResponse(progress)
	}

	return responses, nil
}

// Helper function to convert entity to response
func (s *Service) toWatchProgressResponse(progress *domain.WatchProgress) *WatchProgressResponse {
	return &WatchProgressResponse{
		ID:         progress.ID,
		UUID:       progress.UUID,
		UserID:     progress.UserID,
		EpisodeID:  progress.EpisodeID,
		WatchedAt:  progress.WatchedAt.Format(time.RFC3339),
		Duration:   progress.Duration,
		Completed:  progress.Completed,
		Percentage: progress.Percentage,
		CreatedAt:  progress.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  progress.UpdatedAt.Format(time.RFC3339),
	}
}
