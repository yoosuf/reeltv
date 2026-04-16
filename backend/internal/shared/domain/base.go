package domain

import "time"

// BaseEntity provides common fields for all entities
type BaseEntity struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SoftDeleteEntity provides soft delete capability
type SoftDeleteEntity struct {
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

// UUIDEntity provides UUID field for entities
type UUIDEntity struct {
	UUID string `gorm:"type:varchar(36);uniqueIndex;not null" json:"uuid"`
}
