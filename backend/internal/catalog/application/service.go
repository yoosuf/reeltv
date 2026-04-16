package application

import (
	"context"
	"fmt"
	"time"

	"reeltv/backend/internal/catalog/domain"
	sharedDomain "reeltv/backend/internal/shared/domain"
	"reeltv/backend/pkg/utils"
)

var (
	ErrSeriesNotFound      = sharedDomain.ErrNotFound
	ErrSeasonNotFound      = sharedDomain.ErrNotFound
	ErrEpisodeNotFound     = sharedDomain.ErrNotFound
	ErrSeriesAlreadyExists = sharedDomain.ErrAlreadyExists
)

// Service implements catalog use cases
type Service struct {
	seriesRepo  domain.SeriesRepository
	seasonRepo  domain.SeasonRepository
	episodeRepo domain.EpisodeRepository
}

// NewService creates a new catalog application service
func NewService(
	seriesRepo domain.SeriesRepository,
	seasonRepo domain.SeasonRepository,
	episodeRepo domain.EpisodeRepository,
) *Service {
	return &Service{
		seriesRepo:  seriesRepo,
		seasonRepo:  seasonRepo,
		episodeRepo: episodeRepo,
	}
}

// CreateSeries creates a new series
func (s *Service) CreateSeries(ctx context.Context, req *CreateSeriesRequest) (*SeriesResponse, error) {
	// Check if series already exists by slug
	existingSeries, err := s.seriesRepo.FindBySlug(ctx, req.Slug)
	if err == nil && existingSeries != nil {
		return nil, ErrSeriesAlreadyExists
	}

	// Parse release date
	var releaseDate time.Time
	if req.ReleaseDate != "" {
		releaseDate, err = time.Parse(time.RFC3339, req.ReleaseDate)
		if err != nil {
			return nil, fmt.Errorf("invalid release date format")
		}
	}

	// Create series entity
	series := &domain.Series{
		UUIDEntity:  sharedDomain.UUIDEntity{UUID: utils.GenerateUUID()},
		Title:       req.Title,
		Slug:        req.Slug,
		Description: req.Description,
		PosterURL:   req.PosterURL,
		BackdropURL: req.BackdropURL,
		Year:        req.Year,
		Rating:      req.Rating,
		Genre:       req.Genre,
		Language:    req.Language,
		Status:      "active",
		ReleaseDate: releaseDate,
	}

	if err := s.seriesRepo.Create(ctx, series); err != nil {
		return nil, fmt.Errorf("failed to create series: %w", err)
	}

	return s.toSeriesResponse(series), nil
}

// GetSeriesByID retrieves a series by ID
func (s *Service) GetSeriesByID(ctx context.Context, id uint) (*SeriesResponse, error) {
	series, err := s.seriesRepo.FindByID(ctx, id)
	if err != nil {
		return nil, ErrSeriesNotFound
	}

	return s.toSeriesResponse(series), nil
}

// GetSeriesBySlug retrieves a series by slug
func (s *Service) GetSeriesBySlug(ctx context.Context, slug string) (*SeriesResponse, error) {
	series, err := s.seriesRepo.FindBySlug(ctx, slug)
	if err != nil {
		return nil, ErrSeriesNotFound
	}

	return s.toSeriesResponse(series), nil
}

// ListSeries lists series with pagination
func (s *Service) ListSeries(ctx context.Context, offset, limit int) ([]*SeriesResponse, error) {
	seriesList, err := s.seriesRepo.List(ctx, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to list series: %w", err)
	}

	responses := make([]*SeriesResponse, len(seriesList))
	for i, series := range seriesList {
		responses[i] = s.toSeriesResponse(series)
	}

	return responses, nil
}

// UpdateSeries updates a series
func (s *Service) UpdateSeries(ctx context.Context, id uint, req *UpdateSeriesRequest) (*SeriesResponse, error) {
	series, err := s.seriesRepo.FindByID(ctx, id)
	if err != nil {
		return nil, ErrSeriesNotFound
	}

	// Update fields
	if req.Title != "" {
		series.Title = req.Title
	}
	if req.Slug != "" {
		series.Slug = req.Slug
	}
	if req.Description != "" {
		series.Description = req.Description
	}
	if req.PosterURL != "" {
		series.PosterURL = req.PosterURL
	}
	if req.BackdropURL != "" {
		series.BackdropURL = req.BackdropURL
	}
	if req.Year > 0 {
		series.Year = req.Year
	}
	if req.Rating > 0 {
		series.Rating = req.Rating
	}
	if req.Genre != "" {
		series.Genre = req.Genre
	}
	if req.Language != "" {
		series.Language = req.Language
	}
	if req.Status != "" {
		series.Status = req.Status
	}

	if err := s.seriesRepo.Update(ctx, series); err != nil {
		return nil, fmt.Errorf("failed to update series: %w", err)
	}

	return s.toSeriesResponse(series), nil
}

// DeleteSeries deletes a series
func (s *Service) DeleteSeries(ctx context.Context, id uint) error {
	return s.seriesRepo.Delete(ctx, id)
}

// CreateSeason creates a new season
func (s *Service) CreateSeason(ctx context.Context, req *CreateSeasonRequest) (*SeasonResponse, error) {
	// Parse release date
	var releaseDate time.Time
	var err error
	if req.ReleaseDate != "" {
		releaseDate, err = time.Parse(time.RFC3339, req.ReleaseDate)
		if err != nil {
			return nil, fmt.Errorf("invalid release date format")
		}
	}

	// Create season entity
	season := &domain.Season{
		UUIDEntity:   sharedDomain.UUIDEntity{UUID: utils.GenerateUUID()},
		SeriesID:     req.SeriesID,
		SeasonNumber: req.SeasonNumber,
		Title:        req.Title,
		Description:  req.Description,
		PosterURL:    req.PosterURL,
		ReleaseDate:  releaseDate,
		EpisodeCount: 0,
	}

	if err := s.seasonRepo.Create(ctx, season); err != nil {
		return nil, fmt.Errorf("failed to create season: %w", err)
	}

	return s.toSeasonResponse(season), nil
}

// GetSeasonByID retrieves a season by ID
func (s *Service) GetSeasonByID(ctx context.Context, id uint) (*SeasonResponse, error) {
	season, err := s.seasonRepo.FindByID(ctx, id)
	if err != nil {
		return nil, ErrSeasonNotFound
	}

	return s.toSeasonResponse(season), nil
}

// GetSeasonsBySeriesID retrieves all seasons for a series
func (s *Service) GetSeasonsBySeriesID(ctx context.Context, seriesID uint) ([]*SeasonResponse, error) {
	seasons, err := s.seasonRepo.FindBySeriesID(ctx, seriesID)
	if err != nil {
		return nil, fmt.Errorf("failed to get seasons: %w", err)
	}

	responses := make([]*SeasonResponse, len(seasons))
	for i, season := range seasons {
		responses[i] = s.toSeasonResponse(season)
	}

	return responses, nil
}

// CreateEpisode creates a new episode
func (s *Service) CreateEpisode(ctx context.Context, req *CreateEpisodeRequest) (*EpisodeResponse, error) {
	// Parse release date
	var releaseDate time.Time
	var err error
	if req.ReleaseDate != "" {
		releaseDate, err = time.Parse(time.RFC3339, req.ReleaseDate)
		if err != nil {
			return nil, fmt.Errorf("invalid release date format")
		}
	}

	// Create episode entity
	episode := &domain.Episode{
		UUIDEntity:    sharedDomain.UUIDEntity{UUID: utils.GenerateUUID()},
		SeasonID:      req.SeasonID,
		EpisodeNumber: req.EpisodeNumber,
		Title:         req.Title,
		Description:   req.Description,
		ThumbnailURL:  req.ThumbnailURL,
		Duration:      req.Duration,
		VideoURL:      req.VideoURL,
		ReleaseDate:   releaseDate,
		IsPremium:     req.IsPremium,
	}

	if err := s.episodeRepo.Create(ctx, episode); err != nil {
		return nil, fmt.Errorf("failed to create episode: %w", err)
	}

	// Update season episode count
	season, err := s.seasonRepo.FindByID(ctx, req.SeasonID)
	if err == nil {
		season.EpisodeCount++
		s.seasonRepo.Update(ctx, season)
	}

	return s.toEpisodeResponse(episode), nil
}

// GetEpisodeByID retrieves an episode by ID
func (s *Service) GetEpisodeByID(ctx context.Context, id uint) (*EpisodeResponse, error) {
	episode, err := s.episodeRepo.FindByID(ctx, id)
	if err != nil {
		return nil, ErrEpisodeNotFound
	}

	return s.toEpisodeResponse(episode), nil
}

// GetEpisodesBySeasonID retrieves all episodes for a season
func (s *Service) GetEpisodesBySeasonID(ctx context.Context, seasonID uint, offset, limit int) ([]*EpisodeResponse, error) {
	episodes, err := s.episodeRepo.FindBySeasonID(ctx, seasonID, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get episodes: %w", err)
	}

	responses := make([]*EpisodeResponse, len(episodes))
	for i, episode := range episodes {
		responses[i] = s.toEpisodeResponse(episode)
	}

	return responses, nil
}

// Helper functions to convert entities to responses
func (s *Service) toSeriesResponse(series *domain.Series) *SeriesResponse {
	return &SeriesResponse{
		ID:          series.ID,
		UUID:        series.UUID,
		Title:       series.Title,
		Slug:        series.Slug,
		Description: series.Description,
		PosterURL:   series.PosterURL,
		BackdropURL: series.BackdropURL,
		Year:        series.Year,
		Rating:      series.Rating,
		Genre:       series.Genre,
		Language:    series.Language,
		Status:      series.Status,
		ReleaseDate: series.ReleaseDate.Format(time.RFC3339),
		CreatedAt:   series.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   series.UpdatedAt.Format(time.RFC3339),
	}
}

func (s *Service) toSeasonResponse(season *domain.Season) *SeasonResponse {
	return &SeasonResponse{
		ID:           season.ID,
		UUID:         season.UUID,
		SeriesID:     season.SeriesID,
		SeasonNumber: season.SeasonNumber,
		Title:        season.Title,
		Description:  season.Description,
		PosterURL:    season.PosterURL,
		ReleaseDate:  season.ReleaseDate.Format(time.RFC3339),
		EpisodeCount: season.EpisodeCount,
		CreatedAt:    season.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    season.UpdatedAt.Format(time.RFC3339),
	}
}

func (s *Service) toEpisodeResponse(episode *domain.Episode) *EpisodeResponse {
	return &EpisodeResponse{
		ID:            episode.ID,
		UUID:          episode.UUID,
		SeasonID:      episode.SeasonID,
		EpisodeNumber: episode.EpisodeNumber,
		Title:         episode.Title,
		Description:   episode.Description,
		ThumbnailURL:  episode.ThumbnailURL,
		Duration:      episode.Duration,
		VideoURL:      episode.VideoURL,
		ReleaseDate:   episode.ReleaseDate.Format(time.RFC3339),
		IsPremium:     episode.IsPremium,
		CreatedAt:     episode.CreatedAt.Format(time.RFC3339),
		UpdatedAt:     episode.UpdatedAt.Format(time.RFC3339),
	}
}
