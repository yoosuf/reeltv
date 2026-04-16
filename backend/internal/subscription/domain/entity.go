package domain

import (
	"time"

	"reeltv/backend/internal/shared/domain"
)

// Subscription represents a user's subscription
type Subscription struct {
	domain.BaseEntity
	domain.UUIDEntity
	UserID         uint      `gorm:"not null;index" json:"user_id"`
	PlanType       string    `gorm:"type:varchar(50);not null" json:"plan_type"` // free, premium
	Status         string    `gorm:"type:varchar(50);not null;default:'active'" json:"status"`
	StartDate      time.Time `gorm:"not null" json:"start_date"`
	EndDate        time.Time `gorm:"not null" json:"end_date"`
	AutoRenew      bool      `gorm:"default:false" json:"auto_renew"`
}

// Entitlement represents content access rights
type Entitlement struct {
	domain.BaseEntity
	domain.UUIDEntity
	UserID         uint      `gorm:"not null;index" json:"user_id"`
	EpisodeID      uint      `gorm:"not null;index" json:"episode_id"`
	AccessGranted  bool      `gorm:"default:false" json:"access_granted"`
	GrantedAt      time.Time `json:"granted_at"`
	ExpiresAt      *time.Time `json:"expires_at"`
}
