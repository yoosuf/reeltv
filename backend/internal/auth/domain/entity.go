package domain

import (
	"time"

	"reeltv/backend/internal/shared/domain"
)

// RefreshToken represents a refresh token for JWT authentication
type RefreshToken struct {
	domain.BaseEntity
	Token     string    `gorm:"type:varchar(500);uniqueIndex;not null" json:"-"`
	UserID    uint      `gorm:"not null;index" json:"user_id"`
	ExpiresAt time.Time `gorm:"not null;index" json:"expires_at"`
}
