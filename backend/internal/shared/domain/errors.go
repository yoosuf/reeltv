package domain

import "errors"

// Domain errors
var (
	// Common errors
	ErrNotFound       = errors.New("resource not found")
	ErrAlreadyExists  = errors.New("resource already exists")
	ErrInvalidInput   = errors.New("invalid input")
	ErrUnauthorized   = errors.New("unauthorized access")
	ErrForbidden      = errors.New("forbidden access")
	ErrConflict       = errors.New("resource conflict")
	ErrInternal       = errors.New("internal server error")

	// User errors
	ErrUserNotFound      = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserAlreadyExists = errors.New("user already exists")

	// Auth errors
	ErrInvalidToken      = errors.New("invalid token")
	ErrTokenExpired      = errors.New("token expired")
	ErrRefreshTokenNotFound = errors.New("refresh token not found")

	// Catalog errors
	ErrSeriesNotFound    = errors.New("series not found")
	ErrSeasonNotFound    = errors.New("season not found")
	ErrEpisodeNotFound   = errors.New("episode not found")
	ErrGenreNotFound     = errors.New("genre not found")
	ErrTagNotFound       = errors.New("tag not found")

	// Subscription errors
	ErrSubscriptionNotFound = errors.New("subscription not found")
	ErrPremiumRequired     = errors.New("premium subscription required")

	// My List errors
	ErrAlreadyInMyList = errors.New("already in my list")
	ErrNotInMyList    = errors.New("not in my list")
)
