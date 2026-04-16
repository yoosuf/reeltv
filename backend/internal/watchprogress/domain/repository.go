package domain

import (
	"context"
)

// WatchProgressRepository defines the interface for watch progress data access
type WatchProgressRepository interface {
	Create(ctx context.Context, progress *WatchProgress) error
	FindByID(ctx context.Context, id uint) (*WatchProgress, error)
	FindByUserAndEpisode(ctx context.Context, userID, episodeID uint) (*WatchProgress, error)
	FindByUserID(ctx context.Context, userID uint, offset, limit int) ([]*WatchProgress, error)
	Update(ctx context.Context, progress *WatchProgress) error
	Delete(ctx context.Context, id uint) error
}
