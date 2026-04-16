package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"reeltv/backend/internal/catalog/application"
)

// Handler handles HTTP requests for catalog management
type Handler struct {
	catalogService *application.Service
}

// NewHandler creates a new catalog HTTP handler
func NewHandler(catalogService *application.Service) *Handler {
	return &Handler{
		catalogService: catalogService,
	}
}

// ListSeries handles listing series with pagination
func (h *Handler) ListSeries(c *gin.Context) {
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

	series, err := h.catalogService.ListSeries(c.Request.Context(), offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error": gin.H{
				"code":    "INTERNAL_ERROR",
				"message": "Failed to list series",
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

// GetSeries handles getting a series by ID
func (h *Handler) GetSeries(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
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

	series, err := h.catalogService.GetSeriesByID(c.Request.Context(), uint(id))
	if err != nil {
		if err == application.ErrSeriesNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"error": gin.H{
					"code":    "SERIES_NOT_FOUND",
					"message": "Series not found",
				},
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error": gin.H{
				"code":    "INTERNAL_ERROR",
				"message": "Failed to get series",
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    series,
	})
}

// GetSeriesBySlug handles getting a series by slug
func (h *Handler) GetSeriesBySlug(c *gin.Context) {
	slug := c.Param("slug")

	series, err := h.catalogService.GetSeriesBySlug(c.Request.Context(), slug)
	if err != nil {
		if err == application.ErrSeriesNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"error": gin.H{
					"code":    "SERIES_NOT_FOUND",
					"message": "Series not found",
				},
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error": gin.H{
				"code":    "INTERNAL_ERROR",
				"message": "Failed to get series",
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    series,
	})
}

// CreateSeries handles creating a new series (admin only)
func (h *Handler) CreateSeries(c *gin.Context) {
	var req application.CreateSeriesRequest
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

	series, err := h.catalogService.CreateSeries(c.Request.Context(), &req)
	if err != nil {
		if err == application.ErrSeriesAlreadyExists {
			c.JSON(http.StatusConflict, gin.H{
				"success": false,
				"error": gin.H{
					"code":    "SERIES_ALREADY_EXISTS",
					"message": "Series already exists",
				},
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error": gin.H{
				"code":    "INTERNAL_ERROR",
				"message": "Failed to create series",
			},
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    series,
		"message": "Series created successfully",
	})
}

// UpdateSeries handles updating a series (admin only)
func (h *Handler) UpdateSeries(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
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

	var req application.UpdateSeriesRequest
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

	series, err := h.catalogService.UpdateSeries(c.Request.Context(), uint(id), &req)
	if err != nil {
		if err == application.ErrSeriesNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"error": gin.H{
					"code":    "SERIES_NOT_FOUND",
					"message": "Series not found",
				},
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error": gin.H{
				"code":    "INTERNAL_ERROR",
				"message": "Failed to update series",
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    series,
		"message": "Series updated successfully",
	})
}

// DeleteSeries handles deleting a series (admin only)
func (h *Handler) DeleteSeries(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
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

	if err := h.catalogService.DeleteSeries(c.Request.Context(), uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error": gin.H{
				"code":    "INTERNAL_ERROR",
				"message": "Failed to delete series",
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Series deleted successfully",
	})
}

// GetSeasonsBySeriesID handles getting all seasons for a series
func (h *Handler) GetSeasonsBySeriesID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
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

	seasons, err := h.catalogService.GetSeasonsBySeriesID(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error": gin.H{
				"code":    "INTERNAL_ERROR",
				"message": "Failed to get seasons",
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    seasons,
	})
}

// GetEpisodesBySeasonID handles getting all episodes for a season
func (h *Handler) GetEpisodesBySeasonID(c *gin.Context) {
	seasonIDParam := c.Param("season_id")
	seasonID, err := strconv.ParseUint(seasonIDParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error": gin.H{
				"code":    "INVALID_ID",
				"message": "Invalid season ID",
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

	episodes, err := h.catalogService.GetEpisodesBySeasonID(c.Request.Context(), uint(seasonID), offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error": gin.H{
				"code":    "INTERNAL_ERROR",
				"message": "Failed to get episodes",
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"episodes": episodes,
			"offset":   offset,
			"limit":    limit,
		},
	})
}
