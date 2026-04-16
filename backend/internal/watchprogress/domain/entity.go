package domain

import (
	"time"

	"reeltv/backend/internal/shared/domain"
)

// WatchProgress represents a user's watch progress for an episode
type WatchProgress struct {
	domain.BaseEntity
	domain.UUIDEntity
	UserID     uint      `gorm:"not null;index" json:"user_id"`
	EpisodeID  uint      `gorm:"not null;index" json:"episode_id"`
	WatchedAt  time.Time `json:"watched_at"`
	Duration   int       `json:"duration"` // seconds watched
	Completed  bool      `gorm:"default:false" json:"completed"`
	Percentage float32   `json:"percentage"` // 0-100
}
