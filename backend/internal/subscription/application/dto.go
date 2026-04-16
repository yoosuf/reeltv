package application

// CreateSubscriptionRequest represents a request to create a subscription
type CreateSubscriptionRequest struct {
	PlanType  string `json:"plan_type" binding:"required"` // free, premium
	StartDate string `json:"start_date" binding:"required"`
	EndDate   string `json:"end_date" binding:"required"`
	AutoRenew bool   `json:"auto_renew"`
}

// SubscriptionResponse represents subscription data in responses
type SubscriptionResponse struct {
	ID        uint   `json:"id"`
	UUID      string `json:"uuid"`
	UserID    uint   `json:"user_id"`
	PlanType  string `json:"plan_type"`
	Status    string `json:"status"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	AutoRenew bool   `json:"auto_renew"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// CheckAccessRequest represents a request to check episode access
type CheckAccessRequest struct {
	EpisodeID uint `json:"episode_id" binding:"required"`
}

// CheckAccessResponse represents the result of an access check
type CheckAccessResponse struct {
	AccessGranted bool                  `json:"access_granted"`
	Subscription  *SubscriptionResponse `json:"subscription,omitempty"`
	Reason        string                `json:"reason"`
}
