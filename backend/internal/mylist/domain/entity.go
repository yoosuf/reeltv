package domain

import (
	"time"

	"reeltv/backend/internal/catalog/domain"
	sharedDomain "reeltv/backend/internal/shared/domain"
)

// MyListItem represents a user's favorite series
type MyListItem struct {
	sharedDomain.BaseEntity
	sharedDomain.UUIDEntity
	UserID   uint           `gorm:"not null;index" json:"user_id"`
	SeriesID uint           `gorm:"not null;index" json:"series_id"`
	Series   *domain.Series `gorm:"foreignKey:SeriesID" json:"series,omitempty"`
	AddedAt  time.Time      `json:"added_at"`
}
