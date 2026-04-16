package http

import (
	"github.com/gin-gonic/gin"

	adminhttp "reeltv/backend/internal/admin/interface/http"
	authhttp "reeltv/backend/internal/auth/interface/http"
	cataloghttp "reeltv/backend/internal/catalog/interface/http"
	mylisthttp "reeltv/backend/internal/mylist/interface/http"
	recommendationhttp "reeltv/backend/internal/recommendation/interface/http"
	subscriptionhttp "reeltv/backend/internal/subscription/interface/http"
	userhttp "reeltv/backend/internal/user/interface/http"
	watchprogresshttp "reeltv/backend/internal/watchprogress/interface/http"
)

// Router sets up the HTTP routes
type Router struct {
	engine                *gin.Engine
	authHandler           *authhttp.Handler
	userHandler           *userhttp.Handler
	catalogHandler        *cataloghttp.Handler
	watchProgressHandler  *watchprogresshttp.Handler
	myListHandler         *mylisthttp.Handler
	recommendationHandler *recommendationhttp.Handler
	subscriptionHandler   *subscriptionhttp.Handler
	adminHandler          *adminhttp.Handler
}

// NewRouter creates a new router
func NewRouter(
	authHandler *authhttp.Handler,
	userHandler *userhttp.Handler,
	catalogHandler *cataloghttp.Handler,
	watchProgressHandler *watchprogresshttp.Handler,
	myListHandler *mylisthttp.Handler,
	recommendationHandler *recommendationhttp.Handler,
	subscriptionHandler *subscriptionhttp.Handler,
	adminHandler *adminhttp.Handler,
) *Router {
	engine := gin.New()

	// Middleware will be added here
	engine.Use(gin.Recovery())

	return &Router{
		engine:                engine,
		authHandler:           authHandler,
		userHandler:           userHandler,
		catalogHandler:        catalogHandler,
		watchProgressHandler:  watchProgressHandler,
		myListHandler:         myListHandler,
		recommendationHandler: recommendationHandler,
		subscriptionHandler:   subscriptionHandler,
		adminHandler:          adminHandler,
	}
}

// SetupRoutes configures all routes
func (r *Router) SetupRoutes() {
	// Health check
	r.engine.GET("/health", r.healthCheck)

	// API v1 routes
	v1 := r.engine.Group("/api/v1")
	{
		// Auth routes
		auth := v1.Group("/auth")
		{
			auth.POST("/register", r.authHandler.Register)
			auth.POST("/login", r.authHandler.Login)
			auth.POST("/refresh", r.authHandler.RefreshToken)
			auth.POST("/logout", r.authHandler.Logout)
		}

		// Catalog routes
		catalog := v1.Group("/catalog")
		{
			catalog.GET("/series", r.catalogHandler.ListSeries)
			catalog.GET("/series/:id", r.catalogHandler.GetSeries)
			catalog.GET("/series/slug/:slug", r.catalogHandler.GetSeriesBySlug)
			catalog.POST("/series", r.catalogHandler.CreateSeries)
			catalog.PUT("/series/:id", r.catalogHandler.UpdateSeries)
			catalog.DELETE("/series/:id", r.catalogHandler.DeleteSeries)
			catalog.GET("/series/:id/seasons", r.catalogHandler.GetSeasonsBySeriesID)
			catalog.GET("/seasons/:season_id/episodes", r.catalogHandler.GetEpisodesBySeasonID)
		}

		// User routes
		users := v1.Group("/users")
		{
			users.GET("/me", r.userHandler.GetProfile)
			users.PUT("/me", r.userHandler.UpdateProfile)
			users.POST("/me/change-password", r.userHandler.ChangePassword)
			users.GET("/:id", r.userHandler.GetUser)
			users.GET("", r.userHandler.ListUsers)
		}

		// Watch progress routes
		watchProgress := v1.Group("/watch-progress")
		{
			watchProgress.POST("", r.watchProgressHandler.UpdateWatchProgress)
			watchProgress.GET("", r.watchProgressHandler.GetWatchProgress)
		}

		// My list routes
		myList := v1.Group("/my-list")
		{
			myList.POST("", r.myListHandler.AddToMyList)
			myList.GET("", r.myListHandler.GetMyList)
			myList.DELETE("/:series_id", r.myListHandler.RemoveFromMyList)
		}

		// Recommendation routes
		recommendations := v1.Group("/recommendations")
		{
			recommendations.GET("", r.recommendationHandler.GetRecommendations)
			recommendations.GET("/trending", r.recommendationHandler.GetTrendingSeries)
			recommendations.GET("/genre/:genre", r.recommendationHandler.GetRecommendedByGenre)
		}

		// Subscription routes
		subscription := v1.Group("/subscription")
		{
			subscription.GET("", r.subscriptionHandler.GetSubscription)
			subscription.POST("/check-access", r.subscriptionHandler.CheckAccess)
		}

		// Admin routes
		admin := v1.Group("/admin")
		admin.Use(r.adminHandler.RequireAdmin)
		{
			admin.GET("/dashboard", r.adminHandler.AdminDashboard)
		}
	}
}

// GetEngine returns the gin engine
func (r *Router) GetEngine() *gin.Engine {
	return r.engine
}

// healthCheck handles health check requests
func (r *Router) healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"success": true,
		"status":  "healthy",
	})
}
