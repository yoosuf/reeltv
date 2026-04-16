package domain

import (
	"context"
)

// RefreshTokenRepository defines the interface for refresh token data access
type RefreshTokenRepository interface {
	Create(ctx context.Context, token *RefreshToken) error
	FindByToken(ctx context.Context, token string) (*RefreshToken, error)
	FindByUserID(ctx context.Context, userID uint) ([]*RefreshToken, error)
	Delete(ctx context.Context, token string) error
	DeleteByUserID(ctx context.Context, userID uint) error
	DeleteExpired(ctx context.Context) error
}

// UserRepository defines the interface for user data access (auth-specific operations)
// Note: Full user repository is in user module, this is for auth operations
type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (interface{}, error)
	FindByPhone(ctx context.Context, phone string) (interface{}, error)
	FindByID(ctx context.Context, id uint) (interface{}, error)
	Create(ctx context.Context, user interface{}) error
	UpdatePassword(ctx context.Context, userID uint, hashedPassword string) error
}
