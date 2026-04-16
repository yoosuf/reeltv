package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"reeltv/backend/internal/admin/application"
)

// Handler handles HTTP requests for admin operations
type Handler struct {
	adminService *application.Service
}

// NewHandler creates a new admin HTTP handler
func NewHandler(adminService *application.Service) *Handler {
	return &Handler{
		adminService: adminService,
	}
}

// RequireAdmin middleware to check admin access
func (h *Handler) RequireAdmin(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error": gin.H{
				"code":    "UNAUTHORIZED",
				"message": "User not authenticated",
			},
		})
		c.Abort()
		return
	}

	if !h.adminService.IsAdmin(userID.(uint)) {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"error": gin.H{
				"code":    "FORBIDDEN",
				"message": "Admin access required",
			},
		})
		c.Abort()
		return
	}

	c.Next()
}

// AdminDashboard handles admin dashboard endpoint
func (h *Handler) AdminDashboard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"message": "Admin dashboard",
			"stats": gin.H{
				"total_users":    0,
				"total_series":   0,
				"active_subs":    0,
			},
		},
	})
}
