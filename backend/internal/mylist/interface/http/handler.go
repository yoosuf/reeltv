package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"reeltv/backend/internal/mylist/application"
)

// Handler handles HTTP requests for my list management
type Handler struct {
	myListService *application.Service
}

// NewHandler creates a new my list HTTP handler
func NewHandler(myListService *application.Service) *Handler {
	return &Handler{
		myListService: myListService,
	}
}

// AddToMyList handles adding a series to the user's my list
func (h *Handler) AddToMyList(c *gin.Context) {
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

	var req application.AddToMyListRequest
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

	item, err := h.myListService.AddToMyList(c.Request.Context(), userID.(uint), &req)
	if err != nil {
		if err == application.ErrAlreadyInMyList {
			c.JSON(http.StatusConflict, gin.H{
				"success": false,
				"error": gin.H{
					"code":    "ALREADY_IN_MY_LIST",
					"message": "Series already in my list",
				},
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error": gin.H{
				"code":    "INTERNAL_ERROR",
				"message": "Failed to add to my list",
			},
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    item,
		"message": "Added to my list successfully",
	})
}

// GetMyList handles getting the user's my list
func (h *Handler) GetMyList(c *gin.Context) {
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

	items, err := h.myListService.GetMyList(c.Request.Context(), userID.(uint), offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error": gin.H{
				"code":    "INTERNAL_ERROR",
				"message": "Failed to get my list",
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"items":  items,
			"offset": offset,
			"limit":  limit,
		},
	})
}

// RemoveFromMyList handles removing a series from the user's my list
func (h *Handler) RemoveFromMyList(c *gin.Context) {
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

	seriesIDParam := c.Param("series_id")
	seriesID, err := strconv.ParseUint(seriesIDParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error": gin.H{
				"code":    "INVALID_ID",
				"message": "Invalid series ID",
			},
		})
		return
	}

	if err := h.myListService.RemoveFromMyList(c.Request.Context(), userID.(uint), uint(seriesID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error": gin.H{
				"code":    "INTERNAL_ERROR",
				"message": "Failed to remove from my list",
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Removed from my list successfully",
	})
}
