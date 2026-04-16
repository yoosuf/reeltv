package persistence

import (
	"context"

	"gorm.io/gorm"

	"reeltv/backend/internal/watchprogress/domain"
)

// watchProgressRepository implements WatchProgressRepository interface
type watchProgressRepository struct {
	db *gorm.DB
}

// NewWatchProgressRepository creates a new watch progress repository
func NewWatchProgressRepository(db *gorm.DB) domain.WatchProgressRepository {
	return &watchProgressRepository{db: db}
}

func (r *watchProgressRepository) Create(ctx context.Context, progress *domain.WatchProgress) error {
	return r.db.WithContext(ctx).Create(progress).Error
}

func (r *watchProgressRepository) FindByID(ctx context.Context, id uint) (*domain.WatchProgress, error) {
	var progress domain.WatchProgress
	err := r.db.WithContext(ctx).First(&progress, id).Error
	if err != nil {
		return nil, err
	}
	return &progress, nil
}

func (r *watchProgressRepository) FindByUserAndEpisode(ctx context.Context, userID, episodeID uint) (*domain.WatchProgress, error) {
	var progress domain.WatchProgress
	err := r.db.WithContext(ctx).Where("user_id = ? AND episode_id = ?", userID, episodeID).First(&progress).Error
	if err != nil {
		return nil, err
	}
	return &progress, nil
}

func (r *watchProgressRepository) FindByUserID(ctx context.Context, userID uint, offset, limit int) ([]*domain.WatchProgress, error) {
	var progressList []*domain.WatchProgress
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Order("watched_at DESC").Offset(offset).Limit(limit).Find(&progressList).Error
	if err != nil {
		return nil, err
	}
	return progressList, nil
}

func (r *watchProgressRepository) Update(ctx context.Context, progress *domain.WatchProgress) error {
	return r.db.WithContext(ctx).Save(progress).Error
}

func (r *watchProgressRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.WatchProgress{}, id).Error
}
