package domain

import (
	"context"
)

// SeriesRepository defines the interface for series data access
type SeriesRepository interface {
	Create(ctx context.Context, series *Series) error
	FindByID(ctx context.Context, id uint) (*Series, error)
	FindByUUID(ctx context.Context, uuid string) (*Series, error)
	FindBySlug(ctx context.Context, slug string) (*Series, error)
	List(ctx context.Context, offset, limit int) ([]*Series, error)
	Count(ctx context.Context) (int64, error)
	Update(ctx context.Context, series *Series) error
	Delete(ctx context.Context, id uint) error
}

// SeasonRepository defines the interface for season data access
type SeasonRepository interface {
	Create(ctx context.Context, season *Season) error
	FindByID(ctx context.Context, id uint) (*Season, error)
	FindByUUID(ctx context.Context, uuid string) (*Season, error)
	FindBySeriesID(ctx context.Context, seriesID uint) ([]*Season, error)
	Update(ctx context.Context, season *Season) error
	Delete(ctx context.Context, id uint) error
}

// EpisodeRepository defines the interface for episode data access
type EpisodeRepository interface {
	Create(ctx context.Context, episode *Episode) error
	FindByID(ctx context.Context, id uint) (*Episode, error)
	FindByUUID(ctx context.Context, uuid string) (*Episode, error)
	FindBySeasonID(ctx context.Context, seasonID uint, offset, limit int) ([]*Episode, error)
	CountBySeasonID(ctx context.Context, seasonID uint) (int64, error)
	Update(ctx context.Context, episode *Episode) error
	Delete(ctx context.Context, id uint) error
}
