package persistence

import (
	"context"

	"gorm.io/gorm"

	"reeltv/backend/internal/subscription/domain"
)

// subscriptionRepository implements SubscriptionRepository interface
type subscriptionRepository struct {
	db *gorm.DB
}

// NewSubscriptionRepository creates a new subscription repository
func NewSubscriptionRepository(db *gorm.DB) domain.SubscriptionRepository {
	return &subscriptionRepository{db: db}
}

func (r *subscriptionRepository) Create(ctx context.Context, subscription *domain.Subscription) error {
	return r.db.WithContext(ctx).Create(subscription).Error
}

func (r *subscriptionRepository) FindByID(ctx context.Context, id uint) (*domain.Subscription, error) {
	var subscription domain.Subscription
	err := r.db.WithContext(ctx).First(&subscription, id).Error
	if err != nil {
		return nil, err
	}
	return &subscription, nil
}

func (r *subscriptionRepository) FindByUserID(ctx context.Context, userID uint) (*domain.Subscription, error) {
	var subscription domain.Subscription
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).First(&subscription).Error
	if err != nil {
		return nil, err
	}
	return &subscription, nil
}

func (r *subscriptionRepository) Update(ctx context.Context, subscription *domain.Subscription) error {
	return r.db.WithContext(ctx).Save(subscription).Error
}

func (r *subscriptionRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.Subscription{}, id).Error
}

// entitlementRepository implements EntitlementRepository interface
type entitlementRepository struct {
	db *gorm.DB
}

// NewEntitlementRepository creates a new entitlement repository
func NewEntitlementRepository(db *gorm.DB) domain.EntitlementRepository {
	return &entitlementRepository{db: db}
}

func (r *entitlementRepository) Create(ctx context.Context, entitlement *domain.Entitlement) error {
	return r.db.WithContext(ctx).Create(entitlement).Error
}

func (r *entitlementRepository) FindByID(ctx context.Context, id uint) (*domain.Entitlement, error) {
	var entitlement domain.Entitlement
	err := r.db.WithContext(ctx).First(&entitlement, id).Error
	if err != nil {
		return nil, err
	}
	return &entitlement, nil
}

func (r *entitlementRepository) FindByUserAndEpisode(ctx context.Context, userID, episodeID uint) (*domain.Entitlement, error) {
	var entitlement domain.Entitlement
	err := r.db.WithContext(ctx).Where("user_id = ? AND episode_id = ?", userID, episodeID).First(&entitlement).Error
	if err != nil {
		return nil, err
	}
	return &entitlement, nil
}

func (r *entitlementRepository) Update(ctx context.Context, entitlement *domain.Entitlement) error {
	return r.db.WithContext(ctx).Save(entitlement).Error
}

func (r *entitlementRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.Entitlement{}, id).Error
}
