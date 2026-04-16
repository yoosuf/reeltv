package application

import (
	"context"
	"errors"
	"fmt"
	"time"

	sharedDomain "reeltv/backend/internal/shared/domain"
	"reeltv/backend/internal/subscription/domain"
	"reeltv/backend/pkg/utils"
)

var (
	ErrSubscriptionNotFound = sharedDomain.ErrNotFound
	ErrAccessDenied         = errors.New("access denied")
)

// Service implements subscription and entitlement use cases
type Service struct {
	subscriptionRepo domain.SubscriptionRepository
	entitlementRepo  domain.EntitlementRepository
}

// NewService creates a new subscription application service
func NewService(subscriptionRepo domain.SubscriptionRepository, entitlementRepo domain.EntitlementRepository) *Service {
	return &Service{
		subscriptionRepo: subscriptionRepo,
		entitlementRepo:  entitlementRepo,
	}
}

// CreateSubscription creates a new subscription for a user
func (s *Service) CreateSubscription(ctx context.Context, userID uint, req *CreateSubscriptionRequest) (*SubscriptionResponse, error) {
	startDate, err := time.Parse(time.RFC3339, req.StartDate)
	if err != nil {
		return nil, fmt.Errorf("invalid start date format: %w", err)
	}

	endDate, err := time.Parse(time.RFC3339, req.EndDate)
	if err != nil {
		return nil, fmt.Errorf("invalid end date format: %w", err)
	}

	subscription := &domain.Subscription{
		UUIDEntity: sharedDomain.UUIDEntity{UUID: utils.GenerateUUID()},
		UserID:     userID,
		PlanType:   req.PlanType,
		Status:     "active",
		StartDate:  startDate,
		EndDate:    endDate,
		AutoRenew:  req.AutoRenew,
	}

	if err := s.subscriptionRepo.Create(ctx, subscription); err != nil {
		return nil, fmt.Errorf("failed to create subscription: %w", err)
	}

	return s.toSubscriptionResponse(subscription), nil
}

// GetSubscription retrieves a user's subscription
func (s *Service) GetSubscription(ctx context.Context, userID uint) (*SubscriptionResponse, error) {
	subscription, err := s.subscriptionRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get subscription: %w", err)
	}

	return s.toSubscriptionResponse(subscription), nil
}

// CheckAccess checks if a user has access to an episode
func (s *Service) CheckAccess(ctx context.Context, userID, episodeID uint) (*CheckAccessResponse, error) {
	// Get user's subscription
	subscription, err := s.subscriptionRepo.FindByUserID(ctx, userID)
	if err != nil {
		// If no subscription, check if it's free content
		return &CheckAccessResponse{
			AccessGranted: false,
			Reason:        "No active subscription",
		}, nil
	}

	// Check if subscription is active
	now := utils.NowUTC()
	if subscription.Status != "active" || now.After(subscription.EndDate) {
		return &CheckAccessResponse{
			AccessGranted: false,
			Reason:        "Subscription expired or inactive",
		}, nil
	}

	// Premium users have access to all content
	if subscription.PlanType == "premium" {
		return &CheckAccessResponse{
			AccessGranted: true,
			Subscription:  s.toSubscriptionResponse(subscription),
			Reason:        "Premium subscription grants access",
		}, nil
	}

	// Free users have limited access (first episode only, for MVP)
	// In a full implementation, this would check episode limits
	return &CheckAccessResponse{
		AccessGranted: false,
		Subscription:  s.toSubscriptionResponse(subscription),
		Reason:        "Free tier has limited access",
	}, nil
}

// Helper function to convert entity to response
func (s *Service) toSubscriptionResponse(subscription *domain.Subscription) *SubscriptionResponse {
	return &SubscriptionResponse{
		ID:        subscription.ID,
		UUID:      subscription.UUID,
		UserID:    subscription.UserID,
		PlanType:  subscription.PlanType,
		Status:    subscription.Status,
		StartDate: subscription.StartDate.Format(time.RFC3339),
		EndDate:   subscription.EndDate.Format(time.RFC3339),
		AutoRenew: subscription.AutoRenew,
		CreatedAt: subscription.CreatedAt.Format(time.RFC3339),
		UpdatedAt: subscription.UpdatedAt.Format(time.RFC3339),
	}
}
