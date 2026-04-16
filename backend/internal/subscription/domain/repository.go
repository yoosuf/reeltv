package domain

import (
	"context"
)

// SubscriptionRepository defines the interface for subscription data access
type SubscriptionRepository interface {
	Create(ctx context.Context, subscription *Subscription) error
	FindByID(ctx context.Context, id uint) (*Subscription, error)
	FindByUserID(ctx context.Context, userID uint) (*Subscription, error)
	Update(ctx context.Context, subscription *Subscription) error
	Delete(ctx context.Context, id uint) error
}

// EntitlementRepository defines the interface for entitlement data access
type EntitlementRepository interface {
	Create(ctx context.Context, entitlement *Entitlement) error
	FindByID(ctx context.Context, id uint) (*Entitlement, error)
	FindByUserAndEpisode(ctx context.Context, userID, episodeID uint) (*Entitlement, error)
	Update(ctx context.Context, entitlement *Entitlement) error
	Delete(ctx context.Context, id uint) error
}
