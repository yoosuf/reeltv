package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/spf13/viper"
)

// Config holds application configuration
type Config struct {
	App        AppConfig
	Database   DatabaseConfig
	Redis      RedisConfig
	JWT        JWTConfig
	S3         S3Config
	CORS       CORSConfig
	RateLimit  RateLimitConfig
	Logging    LoggingConfig
	Features   FeatureFlags
}

// AppConfig holds application configuration
type AppConfig struct {
	Env  string
	Port int
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// RedisConfig holds Redis configuration
type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

// JWTConfig holds JWT configuration
type JWTConfig struct {
	Secret           string
	AccessExpiration time.Duration
	RefreshExpiration time.Duration
}

// S3Config holds S3 configuration
type S3Config struct {
	Endpoint  string
	AccessKey string
	SecretKey string
	Bucket    string
	Region    string
	UseSSL    bool
}

// CORSConfig holds CORS configuration
type CORSConfig struct {
	AllowedOrigins []string
}

// RateLimitConfig holds rate limiting configuration
type RateLimitConfig struct {
	Enabled               bool
	RequestsPerMinute     int
	AuthRequestsPerMinute int
}

// LoggingConfig holds logging configuration
type LoggingConfig struct {
	Level  string
	Format string
}

// FeatureFlags holds feature flag configuration
type FeatureFlags struct {
	EnableAnalytics      bool
	EnableRecommendations bool
}

// Load loads configuration from environment variables and config file
func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath(".")

	// Read config file if exists
	if err := viper.ReadInConfig(); err != nil {
		// Config file not found is OK, we'll use env vars
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("failed to read config file: %w", err)
		}
	}

	// Bind environment variables
	viper.AutomaticEnv()

	cfg := &Config{}

	// App configuration
	cfg.App.Env = getString("APP_ENV", "development")
	cfg.App.Port = getInt("APP_PORT", 8080)

	// Database configuration
	cfg.Database.Host = getString("DB_HOST", "localhost")
	cfg.Database.Port = getString("DB_PORT", "5432")
	cfg.Database.User = getString("DB_USER", "reeltv")
	cfg.Database.Password = getString("DB_PASSWORD", "")
	cfg.Database.DBName = getString("DB_NAME", "reeltv")
	cfg.Database.SSLMode = getString("DB_SSLMODE", "disable")

	// Redis configuration
	cfg.Redis.Host = getString("REDIS_HOST", "localhost")
	cfg.Redis.Port = getString("REDIS_PORT", "6379")
	cfg.Redis.Password = getString("REDIS_PASSWORD", "")
	cfg.Redis.DB = getInt("REDIS_DB", 0)

	// JWT configuration
	cfg.JWT.Secret = getString("JWT_SECRET", "change-this-secret-in-production")
	accessExp := getString("JWT_ACCESS_EXPIRATION", "15m")
	refreshExp := getString("JWT_REFRESH_EXPIRATION", "168h")
	
	var err error
	cfg.JWT.AccessExpiration, err = time.ParseDuration(accessExp)
	if err != nil {
		cfg.JWT.AccessExpiration = 15 * time.Minute
	}
	
	cfg.JWT.RefreshExpiration, err = time.ParseDuration(refreshExp)
	if err != nil {
		cfg.JWT.RefreshExpiration = 168 * time.Hour
	}

	// S3 configuration
	cfg.S3.Endpoint = getString("S3_ENDPOINT", "")
	cfg.S3.AccessKey = getString("S3_ACCESS_KEY", "")
	cfg.S3.SecretKey = getString("S3_SECRET_KEY", "")
	cfg.S3.Bucket = getString("S3_BUCKET", "")
	cfg.S3.Region = getString("S3_REGION", "us-east-1")
	cfg.S3.UseSSL = getBool("S3_USE_SSL", true)

	// CORS configuration
	corsOrigins := getString("CORS_ALLOWED_ORIGINS", "http://localhost:3000")
	cfg.CORS.AllowedOrigins = parseCommaSeparated(corsOrigins)

	// Rate limiting configuration
	cfg.RateLimit.Enabled = getBool("RATE_LIMIT_ENABLED", true)
	cfg.RateLimit.RequestsPerMinute = getInt("RATE_LIMIT_REQUESTS_PER_MINUTE", 100)
	cfg.RateLimit.AuthRequestsPerMinute = getInt("RATE_LIMIT_AUTH_REQUESTS_PER_MINUTE", 5)

	// Logging configuration
	cfg.Logging.Level = getString("LOG_LEVEL", "info")
	cfg.Logging.Format = getString("LOG_FORMAT", "json")

	// Feature flags
	cfg.Features.EnableAnalytics = getBool("ENABLE_ANALYTICS", true)
	cfg.Features.EnableRecommendations = getBool("ENABLE_RECOMMENDATIONS", true)

	return cfg, nil
}

// Helper functions
func getString(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	if value := viper.GetString(key); value != "" {
		return value
	}
	return defaultValue
}

func getInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	if value := viper.GetInt(key); value != 0 {
		return value
	}
	return defaultValue
}

func getBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolVal, err := strconv.ParseBool(value); err == nil {
			return boolVal
		}
	}
	if value := viper.GetBool(key); viper.IsSet(key) {
		return value
	}
	return defaultValue
}

func parseCommaSeparated(value string) []string {
	if value == "" {
		return []string{}
	}
	var result []string
	for _, item := range splitComma(value) {
		if item != "" {
			result = append(result, item)
		}
	}
	return result
}

func splitComma(s string) []string {
	var result []string
	var current string
	for _, r := range s {
		if r == ',' {
			result = append(result, current)
			current = ""
		} else {
			current += string(r)
		}
	}
	result = append(result, current)
	return result
}
