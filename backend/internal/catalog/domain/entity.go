package domain

import (
	"time"

	"reeltv/backend/internal/shared/domain"
)

// Series represents a TV series
type Series struct {
	domain.BaseEntity
	domain.UUIDEntity
	Title       string    `gorm:"type:varchar(255);not null" json:"title"`
	Slug        string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"slug"`
	Description string    `gorm:"type:text" json:"description"`
	PosterURL   string    `gorm:"type:varchar(500)" json:"poster_url"`
	BackdropURL string    `gorm:"type:varchar(500)" json:"backdrop_url"`
	Year        int       `gorm:"not null" json:"year"`
	Rating      float32   `gorm:"type:decimal(3,1)" json:"rating"`
	Genre       string    `gorm:"type:varchar(100)" json:"genre"`
	Language    string    `gorm:"type:varchar(50);default:'en'" json:"language"`
	Status      string    `gorm:"type:varchar(50);default:'active'" json:"status"`
	ReleaseDate time.Time `json:"release_date"`
}

// Season represents a season in a series
type Season struct {
	domain.BaseEntity
	domain.UUIDEntity
	SeriesID    uint      `gorm:"not null;index" json:"series_id"`
	Series      *Series   `gorm:"foreignKey:SeriesID" json:"series,omitempty"`
	SeasonNumber int      `gorm:"not null" json:"season_number"`
	Title       string    `gorm:"type:varchar(255);not null" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	PosterURL   string    `gorm:"type:varchar(500)" json:"poster_url"`
	ReleaseDate time.Time `json:"release_date"`
	EpisodeCount int      `gorm:"not null;default:0" json:"episode_count"`
}

// Episode represents an episode in a season
type Episode struct {
	domain.BaseEntity
	domain.UUIDEntity
	SeasonID      uint      `gorm:"not null;index" json:"season_id"`
	Season        *Season   `gorm:"foreignKey:SeasonID" json:"season,omitempty"`
	EpisodeNumber int       `gorm:"not null" json:"episode_number"`
	Title         string    `gorm:"type:varchar(255);not null" json:"title"`
	Description   string    `gorm:"type:text" json:"description"`
	ThumbnailURL  string    `gorm:"type:varchar(500)" json:"thumbnail_url"`
	Duration      int       `gorm:"not null" json:"duration"` // in seconds
	VideoURL      string    `gorm:"type:varchar(500)" json:"video_url"`
	ReleaseDate   time.Time `json:"release_date"`
	IsPremium     bool      `gorm:"default:false" json:"is_premium"`
}
