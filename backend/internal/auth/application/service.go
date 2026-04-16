package application

import (
	"context"
	"errors"
	"fmt"
	"time"

	"reeltv/backend/internal/auth/domain"
	"reeltv/backend/pkg/password"
	"reeltv/backend/pkg/utils"
)

var (
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidToken       = errors.New("invalid token")
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidInput       = errors.New("invalid input")
)

// Service implements authentication use cases
type Service struct {
	userRepo         domain.UserRepository
	refreshTokenRepo domain.RefreshTokenRepository
	jwtService       domain.AuthService
}

// NewService creates a new auth application service
func NewService(
	userRepo domain.UserRepository,
	refreshTokenRepo domain.RefreshTokenRepository,
	jwtService domain.AuthService,
) *Service {
	return &Service{
		userRepo:         userRepo,
		refreshTokenRepo: refreshTokenRepo,
		jwtService:       jwtService,
	}
}

// Register registers a new user
func (s *Service) Register(ctx context.Context, req *RegisterRequest) (*AuthResponse, error) {
	// Check if user already exists by email or phone
	if req.Email != "" {
		existingUser, err := s.userRepo.FindByEmail(ctx, req.Email)
		if err == nil && existingUser != nil {
			return nil, ErrUserAlreadyExists
		}
	}
	if req.Phone != "" {
		existingUser, err := s.userRepo.FindByPhone(ctx, req.Phone)
		if err == nil && existingUser != nil {
			return nil, ErrUserAlreadyExists
		}
	}

	// Hash password
	hashedPassword, err := password.Hash(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Create user (simplified - in real implementation would use user module)
	// TODO: This should delegate to user module's application service
	userData := map[string]interface{}{
		"uuid":     utils.GenerateUUID(),
		"email":    req.Email,
		"phone":    req.Phone,
		"password": hashedPassword,
		"name":     req.Name,
		"role":     "user",
	}

	if err := s.userRepo.Create(ctx, userData); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	// Get the created user
	createdUser, err := s.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve created user: %w", err)
	}

	// Generate tokens
	userID, role := extractUserIDAndRole(createdUser)
	accessToken, refreshToken, err := s.jwtService.GenerateTokens(userID, role)
	if err != nil {
		return nil, fmt.Errorf("failed to generate tokens: %w", err)
	}

	// Store refresh token
	refreshTokenEntity := &domain.RefreshToken{
		Token:     refreshToken,
		UserID:    userID,
		ExpiresAt: calculateRefreshTokenExpiry(),
	}
	if err := s.refreshTokenRepo.Create(ctx, refreshTokenEntity); err != nil {
		return nil, fmt.Errorf("failed to store refresh token: %w", err)
	}

	return &AuthResponse{
		User:         toUserResponse(createdUser),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    900, // 15 minutes in seconds
	}, nil
}

// Login authenticates a user
func (s *Service) Login(ctx context.Context, req *LoginRequest) (*AuthResponse, error) {
	var user interface{}
	var err error

	// Find user by email or phone
	if req.Email != "" {
		user, err = s.userRepo.FindByEmail(ctx, req.Email)
	} else if req.Phone != "" {
		user, err = s.userRepo.FindByPhone(ctx, req.Phone)
	} else {
		return nil, ErrInvalidInput
	}

	if err != nil {
		return nil, ErrInvalidCredentials
	}

	// Verify password
	userData := extractUserData(user)
	storedPassword := userData["password"].(string)
	if !s.jwtService.VerifyPassword(req.Password, storedPassword) {
		return nil, ErrInvalidCredentials
	}

	// Generate tokens
	userID, role := extractUserIDAndRole(user)
	accessToken, refreshToken, err := s.jwtService.GenerateTokens(userID, role)
	if err != nil {
		return nil, fmt.Errorf("failed to generate tokens: %w", err)
	}

	// Store refresh token
	refreshTokenEntity := &domain.RefreshToken{
		Token:     refreshToken,
		UserID:    userID,
		ExpiresAt: calculateRefreshTokenExpiry(),
	}
	if err := s.refreshTokenRepo.Create(ctx, refreshTokenEntity); err != nil {
		return nil, fmt.Errorf("failed to store refresh token: %w", err)
	}

	return &AuthResponse{
		User:         toUserResponse(user),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    900,
	}, nil
}

// RefreshToken refreshes an access token using a refresh token
func (s *Service) RefreshToken(ctx context.Context, req *RefreshTokenRequest) (*AuthResponse, error) {
	// Validate refresh token
	userID, err := s.jwtService.ValidateRefreshToken(req.RefreshToken)
	if err != nil {
		return nil, ErrInvalidToken
	}

	// Check if refresh token exists in database
	_, err = s.refreshTokenRepo.FindByToken(ctx, req.RefreshToken)
	if err != nil {
		return nil, ErrInvalidToken
	}

	// Get user
	user, err := s.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, ErrUserNotFound
	}

	// Generate new tokens
	_, role := extractUserIDAndRole(user)
	accessToken, newRefreshToken, err := s.jwtService.GenerateTokens(userID, role)
	if err != nil {
		return nil, fmt.Errorf("failed to generate tokens: %w", err)
	}

	// Delete old refresh token
	s.refreshTokenRepo.Delete(ctx, req.RefreshToken)

	// Store new refresh token
	newRefreshTokenEntity := &domain.RefreshToken{
		Token:     newRefreshToken,
		UserID:    userID,
		ExpiresAt: calculateRefreshTokenExpiry(),
	}
	if err := s.refreshTokenRepo.Create(ctx, newRefreshTokenEntity); err != nil {
		return nil, fmt.Errorf("failed to store refresh token: %w", err)
	}

	return &AuthResponse{
		User:         toUserResponse(user),
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    900,
	}, nil
}

// Logout invalidates the refresh token
func (s *Service) Logout(ctx context.Context, refreshToken string) error {
	return s.refreshTokenRepo.Delete(ctx, refreshToken)
}

// ValidateAccessToken validates an access token and returns user ID and role
func (s *Service) ValidateAccessToken(token string) (uint, string, error) {
	return s.jwtService.ValidateAccessToken(token)
}

// Helper functions (simplified - in real implementation would use proper user entity)
func extractUserIDAndRole(user interface{}) (uint, string) {
	// This is a simplified version - real implementation would properly extract from user entity
	userData := extractUserData(user)
	userID := uint(userData["id"].(int))
	role := userData["role"].(string)
	return userID, role
}

func extractUserData(user interface{}) map[string]interface{} {
	// Simplified - real implementation would use proper type assertion
	return map[string]interface{}{
		"id":       1,
		"uuid":     "uuid",
		"email":    "email",
		"phone":    "phone",
		"name":     "name",
		"role":     "user",
		"password": "hash",
	}
}

func toUserResponse(user interface{}) *UserResponse {
	userData := extractUserData(user)
	return &UserResponse{
		ID:        uint(userData["id"].(int)),
		UUID:      userData["uuid"].(string),
		Email:     userData["email"].(string),
		Phone:     userData["phone"].(string),
		Name:      userData["name"].(string),
		AvatarURL: "",
		Role:      userData["role"].(string),
	}
}

func calculateRefreshTokenExpiry() time.Time {
	// Simplified - real implementation would calculate based on config (e.g., 7 days from now)
	return time.Now().Add(7 * 24 * time.Hour)
}
