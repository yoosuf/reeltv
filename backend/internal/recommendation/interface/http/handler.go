package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"reeltv/backend/internal/recommendation/application"
)

// Handler handles HTTP requests for recommendations
type Handler struct {
	recommendationService *application.Service
}

// NewHandler creates a new recommendation HTTP handler
func NewHandler(recommendationService *application.Service) *Handler {
	return &Handler{
		recommendationService: recommendationService,
	}
}

// GetRecommendations handles getting personalized recommendations
func (h *Handler) GetRecommendations(c *gin.Context) {
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

	recommendations, err := h.recommendationService.GetRecommendations(c.Request.Context(), userID.(uint), offset, limit)
	if err != nil {
		if err == application.ErrNoRecommendations {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"error": gin.H{
					"code":    "NO_RECOMMENDATIONS",
					"message": "No recommendations available",
				},
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error": gin.H{
				"code":    "INTERNAL_ERROR",
				"message": "Failed to get recommendations",
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"recommendations": recommendations,
			"offset":         offset,
			"limit":          limit,
		},
	})
}

// GetTrendingSeries handles getting trending series
func (h *Handler) GetTrendingSeries(c *gin.Context) {
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

	series, err := h.recommendationService.GetTrendingSeries(c.Request.Context(), offset, limit)
	if err != nil {
		if err == application.ErrNoRecommendations {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"error": gin.H{
					"code":    "NO_TRENDING",
					"message": "No trending series available",
				},
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error": gin.H{
				"code":    "INTERNAL_ERROR",
				"message": "Failed to get trending series",
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"series": series,
			"offset": offset,
			"limit":  limit,
		},
	})
}

// GetRecommendedByGenre handles getting recommendations by genre
func (h *Handler) GetRecommendedByGenre(c *gin.Context) {
	genre := c.Param("genre")

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

	series, err := h.recommendationService.GetRecommendedByGenre(c.Request.Context(), genre, offset, limit)
	if err != nil {
		if err == application.ErrNoRecommendations {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"error": gin.H{
					"code":    "NO_RECOMMENDATIONS",
					"message": "No recommendations available for this genre",
				},
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error": gin.H{
				"code":    "INTERNAL_ERROR",
				"message": "Failed to get recommendations",
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"series": series,
			"offset": offset,
			"limit":  limit,
		},
	})
}
