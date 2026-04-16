package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"reeltv/backend/internal/subscription/application"
)

// Handler handles HTTP requests for subscription management
type Handler struct {
	subscriptionService *application.Service
}

// NewHandler creates a new subscription HTTP handler
func NewHandler(subscriptionService *application.Service) *Handler {
	return &Handler{
		subscriptionService: subscriptionService,
	}
}

// GetSubscription handles getting the user's subscription
func (h *Handler) GetSubscription(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error": gin.H{
				"code":    "UNAUTHORIZED",
				"message": "User not authenticated",
			},
		})
		return
	}

	subscription, err := h.subscriptionService.GetSubscription(c.Request.Context(), userID.(uint))
	if err != nil {
		if err == application.ErrSubscriptionNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"error": gin.H{
					"code":    "SUBSCRIPTION_NOT_FOUND",
					"message": "Subscription not found",
				},
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error": gin.H{
				"code":    "INTERNAL_ERROR",
				"message": "Failed to get subscription",
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    subscription,
	})
}

// CheckAccess handles checking episode access
func (h *Handler) CheckAccess(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error": gin.H{
				"code":    "UNAUTHORIZED",
				"message": "User not authenticated",
			},
		})
		return
	}

	var req application.CheckAccessRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error": gin.H{
				"code":    "INVALID_REQUEST",
				"message": "Invalid request parameters",
			},
		})
		return
	}

	access, err := h.subscriptionService.CheckAccess(c.Request.Context(), userID.(uint), req.EpisodeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error": gin.H{
				"code":    "INTERNAL_ERROR",
				"message": "Failed to check access",
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    access,
	})
}
