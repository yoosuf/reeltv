package domain

import (
	"time"

	"reeltv/backend/internal/shared/domain"
)

// User represents a user in the system
type User struct {
	domain.BaseEntity
	domain.UUIDEntity
	Email     string     `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Phone     string     `gorm:"type:varchar(20);uniqueIndex" json:"phone"`
	Password  string     `gorm:"type:varchar(255);not null" json:"-"`
	Name      string     `gorm:"type:varchar(255);not null" json:"name"`
	AvatarURL string     `gorm:"type:varchar(500)" json:"avatar_url"`
	Role      string     `gorm:"type:varchar(50);not null;default:'user'" json:"role"`
	Status    string     `gorm:"type:varchar(50);not null;default:'active'" json:"status"`
	LastLogin *time.Time `json:"last_login_at"`
}
