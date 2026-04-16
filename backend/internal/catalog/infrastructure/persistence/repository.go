package persistence

import (
	"context"

	"gorm.io/gorm"

	"reeltv/backend/internal/catalog/domain"
)

// seriesRepository implements SeriesRepository interface
type seriesRepository struct {
	db *gorm.DB
}

// NewSeriesRepository creates a new series repository
func NewSeriesRepository(db *gorm.DB) domain.SeriesRepository {
	return &seriesRepository{db: db}
}

func (r *seriesRepository) Create(ctx context.Context, series *domain.Series) error {
	return r.db.WithContext(ctx).Create(series).Error
}

func (r *seriesRepository) FindByID(ctx context.Context, id uint) (*domain.Series, error) {
	var series domain.Series
	err := r.db.WithContext(ctx).First(&series, id).Error
	if err != nil {
		return nil, err
	}
	return &series, nil
}

func (r *seriesRepository) FindByUUID(ctx context.Context, uuid string) (*domain.Series, error) {
	var series domain.Series
	err := r.db.WithContext(ctx).Where("uuid = ?", uuid).First(&series).Error
	if err != nil {
		return nil, err
	}
	return &series, nil
}

func (r *seriesRepository) FindBySlug(ctx context.Context, slug string) (*domain.Series, error) {
	var series domain.Series
	err := r.db.WithContext(ctx).Where("slug = ?", slug).First(&series).Error
	if err != nil {
		return nil, err
	}
	return &series, nil
}

func (r *seriesRepository) List(ctx context.Context, offset, limit int) ([]*domain.Series, error) {
	var series []*domain.Series
	err := r.db.WithContext(ctx).Offset(offset).Limit(limit).Find(&series).Error
	if err != nil {
		return nil, err
	}
	return series, nil
}

func (r *seriesRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&domain.Series{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *seriesRepository) Update(ctx context.Context, series *domain.Series) error {
	return r.db.WithContext(ctx).Save(series).Error
}

func (r *seriesRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.Series{}, id).Error
}

// seasonRepository implements SeasonRepository interface
type seasonRepository struct {
	db *gorm.DB
}

// NewSeasonRepository creates a new season repository
func NewSeasonRepository(db *gorm.DB) domain.SeasonRepository {
	return &seasonRepository{db: db}
}

func (r *seasonRepository) Create(ctx context.Context, season *domain.Season) error {
	return r.db.WithContext(ctx).Create(season).Error
}

func (r *seasonRepository) FindByID(ctx context.Context, id uint) (*domain.Season, error) {
	var season domain.Season
	err := r.db.WithContext(ctx).Preload("Series").First(&season, id).Error
	if err != nil {
		return nil, err
	}
	return &season, nil
}

func (r *seasonRepository) FindByUUID(ctx context.Context, uuid string) (*domain.Season, error) {
	var season domain.Season
	err := r.db.WithContext(ctx).Preload("Series").Where("uuid = ?", uuid).First(&season).Error
	if err != nil {
		return nil, err
	}
	return &season, nil
}

func (r *seasonRepository) FindBySeriesID(ctx context.Context, seriesID uint) ([]*domain.Season, error) {
	var seasons []*domain.Season
	err := r.db.WithContext(ctx).Where("series_id = ?", seriesID).Order("season_number ASC").Find(&seasons).Error
	if err != nil {
		return nil, err
	}
	return seasons, nil
}

func (r *seasonRepository) Update(ctx context.Context, season *domain.Season) error {
	return r.db.WithContext(ctx).Save(season).Error
}

func (r *seasonRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.Season{}, id).Error
}

// episodeRepository implements EpisodeRepository interface
type episodeRepository struct {
	db *gorm.DB
}

// NewEpisodeRepository creates a new episode repository
func NewEpisodeRepository(db *gorm.DB) domain.EpisodeRepository {
	return &episodeRepository{db: db}
}

func (r *episodeRepository) Create(ctx context.Context, episode *domain.Episode) error {
	return r.db.WithContext(ctx).Create(episode).Error
}

func (r *episodeRepository) FindByID(ctx context.Context, id uint) (*domain.Episode, error) {
	var episode domain.Episode
	err := r.db.WithContext(ctx).Preload("Season").First(&episode, id).Error
	if err != nil {
		return nil, err
	}
	return &episode, nil
}

func (r *episodeRepository) FindByUUID(ctx context.Context, uuid string) (*domain.Episode, error) {
	var episode domain.Episode
	err := r.db.WithContext(ctx).Preload("Season").Where("uuid = ?", uuid).First(&episode).Error
	if err != nil {
		return nil, err
	}
	return &episode, nil
}

func (r *episodeRepository) FindBySeasonID(ctx context.Context, seasonID uint, offset, limit int) ([]*domain.Episode, error) {
	var episodes []*domain.Episode
	err := r.db.WithContext(ctx).Where("season_id = ?", seasonID).Order("episode_number ASC").Offset(offset).Limit(limit).Find(&episodes).Error
	if err != nil {
		return nil, err
	}
	return episodes, nil
}

func (r *episodeRepository) CountBySeasonID(ctx context.Context, seasonID uint) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&domain.Episode{}).Where("season_id = ?", seasonID).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *episodeRepository) Update(ctx context.Context, episode *domain.Episode) error {
	return r.db.WithContext(ctx).Save(episode).Error
}

func (r *episodeRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.Episode{}, id).Error
}
