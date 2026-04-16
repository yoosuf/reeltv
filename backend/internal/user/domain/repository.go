package domain

import (
	"context"
	"time"
)

// UserRepository defines the interface for user data access
type UserRepository interface {
	Create(ctx context.Context, user *User) error
	FindByID(ctx context.Context, id uint) (*User, error)
	FindByUUID(ctx context.Context, uuid string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByPhone(ctx context.Context, phone string) (*User, error)
	Update(ctx context.Context, user *User) error
	UpdatePassword(ctx context.Context, userID uint, hashedPassword string) error
	UpdateLastLogin(ctx context.Context, userID uint, lastLogin time.Time) error
	List(ctx context.Context, offset, limit int) ([]*User, error)
	Count(ctx context.Context) (int64, error)
}
