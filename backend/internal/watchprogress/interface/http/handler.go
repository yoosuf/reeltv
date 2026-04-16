package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"reeltv/backend/internal/watchprogress/application"
)

// Handler handles HTTP requests for watch progress management
type Handler struct {
	watchProgressService *application.Service
}

// NewHandler creates a new watch progress HTTP handler
func NewHandler(watchProgressService *application.Service) *Handler {
	return &Handler{
		watchProgressService: watchProgressService,
	}
}

// UpdateWatchProgress handles updating watch progress for an episode
func (h *Handler) UpdateWatchProgress(c *gin.Context) {
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

	var req application.UpdateWatchProgressRequest
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

	progress, err := h.watchProgressService.UpdateWatchProgress(c.Request.Context(), userID.(uint), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error": gin.H{
				"code":    "INTERNAL_ERROR",
				"message": "Failed to update watch progress",
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    progress,
		"message": "Watch progress updated successfully",
	})
}

// GetWatchProgress handles getting watch progress for the current user
func (h *Handler) GetWatchProgress(c *gin.Context) {
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

	offsetParam := c.DefaultQuery("offset", "0")
	limitParam := c.DefaultQuery("limit", "20")

	offset, err := strconv.Atoi(offsetParam)
	if err != nil {
		offset = 0
	}

	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		limit = 20
	}

	progress, err := h.watchProgressService.GetWatchProgressByUser(c.Request.Context(), userID.(uint), offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error": gin.H{
				"code":    "INTERNAL_ERROR",
				"message": "Failed to get watch progress",
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"progress": progress,
			"offset":   offset,
			"limit":    limit,
		},
	})
}
