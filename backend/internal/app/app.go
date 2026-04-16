package app

import (
	"context"
	"errors"
	"strconv"

	adminapplication "reeltv/backend/internal/admin/application"
	adminhttp "reeltv/backend/internal/admin/interface/http"
	analyticsapplication "reeltv/backend/internal/analytics/application"
	"reeltv/backend/internal/auth/application"
	"reeltv/backend/internal/auth/infrastructure/jwt"
	authpersistence "reeltv/backend/internal/auth/infrastructure/persistence"
	authhttp "reeltv/backend/internal/auth/interface/http"
	catalogapplication "reeltv/backend/internal/catalog/application"
	catalogpersistence "reeltv/backend/internal/catalog/infrastructure/persistence"
	cataloghttp "reeltv/backend/internal/catalog/interface/http"
	"reeltv/backend/internal/config"
	"reeltv/backend/internal/interface/http"
	mylistapplication "reeltv/backend/internal/mylist/application"
	mylistpersistence "reeltv/backend/internal/mylist/infrastructure/persistence"
	mylisthttp "reeltv/backend/internal/mylist/interface/http"
	recommendationapplication "reeltv/backend/internal/recommendation/application"
	recommendationhttp "reeltv/backend/internal/recommendation/interface/http"
	sharedpersistence "reeltv/backend/internal/shared/infrastructure/persistence"
	subscriptionapplication "reeltv/backend/internal/subscription/application"
	subscriptionpersistence "reeltv/backend/internal/subscription/infrastructure/persistence"
	subscriptionhttp "reeltv/backend/internal/subscription/interface/http"
	userapplication "reeltv/backend/internal/user/application"
	userpersistence "reeltv/backend/internal/user/infrastructure/persistence"
	userhttp "reeltv/backend/internal/user/interface/http"
	watchprogressapplication "reeltv/backend/internal/watchprogress/application"
	watchprogresspersistence "reeltv/backend/internal/watchprogress/infrastructure/persistence"
	watchprogresshttp "reeltv/backend/internal/watchprogress/interface/http"
	"reeltv/backend/pkg/logger"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// App represents the application
type App struct {
	config           *config.Config
	db               *gorm.DB
	redis            *redis.Client
	httpServer       *http.Server
	analyticsService *analyticsapplication.Service
}

// GetRouter returns the HTTP router for testing purposes
func (a *App) GetRouter() *http.Router {
	return a.httpServer.GetRouter()
}

// New creates a new application instance
func New(cfg *config.Config) (*App, error) {
	// Initialize logger
	if err := logger.Init(logger.Config{
		Level:  cfg.Logging.Level,
		Format: cfg.Logging.Format,
	}); err != nil {
		return nil, err
	}

	log := logger.GetLogger()
	log.Info().Msg("Initializing application")

	// Initialize database
	db, err := sharedpersistence.NewDB(sharedpersistence.Config{
		Host:     cfg.Database.Host,
		Port:     cfg.Database.Port,
		User:     cfg.Database.User,
		Password: cfg.Database.Password,
		DBName:   cfg.Database.DBName,
		SSLMode:  cfg.Database.SSLMode,
	})
	if err != nil {
		return nil, err
	}
	log.Info().Msg("Database connected")

	// Initialize Redis
	redisClient, err := sharedpersistence.NewRedisClient(sharedpersistence.RedisConfig{
		Host:     cfg.Redis.Host,
		Port:     cfg.Redis.Port,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
	if err != nil {
		return nil, err
	}
	log.Info().Msg("Redis connected")

	// Initialize auth module
	jwtProvider := jwt.NewProvider(
		cfg.JWT.Secret,
		cfg.JWT.AccessExpiration,
		cfg.JWT.RefreshExpiration,
	)

	refreshTokenRepo := authpersistence.NewRefreshTokenRepository(db)

	// TODO: Create user repository implementation
	// userRepo := userpersistence.NewUserRepository(db)

	// For now, we'll use a mock user repository
	userRepo := &mockUserRepository{}

	authService := application.NewService(userRepo, refreshTokenRepo, jwtProvider)
	authHandler := authhttp.NewHandler(authService)

	// Initialize user module
	userRepoImpl := userpersistence.NewUserRepository(db)
	userService := userapplication.NewService(userRepoImpl)
	userHandler := userhttp.NewHandler(userService)

	// Initialize catalog module
	catalogSeriesRepo := catalogpersistence.NewSeriesRepository(db)
	catalogSeasonRepo := catalogpersistence.NewSeasonRepository(db)
	catalogEpisodeRepo := catalogpersistence.NewEpisodeRepository(db)
	catalogService := catalogapplication.NewService(catalogSeriesRepo, catalogSeasonRepo, catalogEpisodeRepo)
	catalogHandler := cataloghttp.NewHandler(catalogService)

	// Initialize watch progress module
	watchProgressRepo := watchprogresspersistence.NewWatchProgressRepository(db)
	watchProgressService := watchprogressapplication.NewService(watchProgressRepo)
	watchProgressHandler := watchprogresshttp.NewHandler(watchProgressService)

	// Initialize my list module
	myListRepo := mylistpersistence.NewMyListRepository(db)
	myListService := mylistapplication.NewService(myListRepo)
	myListHandler := mylisthttp.NewHandler(myListService)

	// Initialize recommendation module
	recommendationService := recommendationapplication.NewService(catalogService)
	recommendationHandler := recommendationhttp.NewHandler(recommendationService)

	// Initialize subscription module
	subscriptionRepo := subscriptionpersistence.NewSubscriptionRepository(db)
	entitlementRepo := subscriptionpersistence.NewEntitlementRepository(db)
	subscriptionService := subscriptionapplication.NewService(subscriptionRepo, entitlementRepo)
	subscriptionHandler := subscriptionhttp.NewHandler(subscriptionService)

	// Initialize admin module
	adminService := adminapplication.NewService()
	adminHandler := adminhttp.NewHandler(adminService)

	// Initialize analytics module
	analyticsService := analyticsapplication.NewService()

	// Initialize router
	router := http.NewRouter(authHandler, userHandler, catalogHandler, watchProgressHandler, myListHandler, recommendationHandler, subscriptionHandler, adminHandler)
	router.SetupRoutes()

	// Initialize HTTP server
	httpServer := http.NewServer(cfg.App.Port, router)

	return &App{
		config:           cfg,
		db:               db,
		redis:            redisClient,
		httpServer:       httpServer,
		analyticsService: analyticsService,
	}, nil
}

// Run starts the application
func (a *App) Run() error {
	log := logger.GetLogger()
	log.Info().Str("port", strconv.Itoa(a.config.App.Port)).Msg("Starting HTTP server")

	if err := a.httpServer.Start(); err != nil {
		return err
	}

	return nil
}

// Shutdown gracefully shuts down the application
func (a *App) Shutdown() error {
	log := logger.GetLogger()
	log.Info().Msg("Shutting down application")

	// Close HTTP server
	if err := a.httpServer.Shutdown(context.Background()); err != nil {
		return err
	}

	// Close Redis
	if err := sharedpersistence.CloseRedis(a.redis); err != nil {
		return err
	}

	// Close database
	if err := sharedpersistence.CloseDB(a.db); err != nil {
		return err
	}

	log.Info().Msg("Application shutdown complete")
	return nil
}

// Mock user repository for now - will be replaced with real implementation
type mockUserRepository struct{}

func (m *mockUserRepository) FindByEmail(ctx context.Context, email string) (interface{}, error) {
	// Mock implementation
	return nil, errors.New("not implemented")
}

func (m *mockUserRepository) FindByPhone(ctx context.Context, phone string) (interface{}, error) {
	// Mock implementation
	return nil, errors.New("not implemented")
}

func (m *mockUserRepository) FindByID(ctx context.Context, id uint) (interface{}, error) {
	// Mock implementation
	return nil, errors.New("not implemented")
}

func (m *mockUserRepository) Create(ctx context.Context, user interface{}) error {
	// Mock implementation
	return errors.New("not implemented")
}

func (m *mockUserRepository) UpdatePassword(ctx context.Context, userID uint, hashedPassword string) error {
	// Mock implementation
	return errors.New("not implemented")
}
