package application

import (
	"context"
	"errors"
	"fmt"
	"time"

	sharedDomain "reeltv/backend/internal/shared/domain"
	"reeltv/backend/internal/user/domain"
	"reeltv/backend/pkg/password"
	"reeltv/backend/pkg/utils"
)

var (
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

// Service implements user use cases
type Service struct {
	userRepo domain.UserRepository
}

// NewService creates a new user application service
func NewService(userRepo domain.UserRepository) *Service {
	return &Service{
		userRepo: userRepo,
	}
}

// CreateUser creates a new user
func (s *Service) CreateUser(ctx context.Context, req *CreateUserRequest) (*UserResponse, error) {
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

	// Create user entity
	user := &domain.User{
		UUIDEntity: sharedDomain.UUIDEntity{UUID: utils.GenerateUUID()},
		Email:      req.Email,
		Phone:      req.Phone,
		Password:   hashedPassword,
		Name:       req.Name,
		Role:       "user",
		Status:     "active",
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return s.toUserResponse(user), nil
}

// GetUserByID retrieves a user by ID
func (s *Service) GetUserByID(ctx context.Context, id uint) (*UserResponse, error) {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, ErrUserNotFound
	}

	return s.toUserResponse(user), nil
}

// GetUserByUUID retrieves a user by UUID
func (s *Service) GetUserByUUID(ctx context.Context, uuid string) (*UserResponse, error) {
	user, err := s.userRepo.FindByUUID(ctx, uuid)
	if err != nil {
		return nil, ErrUserNotFound
	}

	return s.toUserResponse(user), nil
}

// UpdateUser updates a user
func (s *Service) UpdateUser(ctx context.Context, id uint, req *UpdateUserRequest) (*UserResponse, error) {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, ErrUserNotFound
	}

	// Update fields
	if req.Name != "" {
		user.Name = req.Name
	}
	if req.AvatarURL != "" {
		user.AvatarURL = req.AvatarURL
	}

	if err := s.userRepo.Update(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return s.toUserResponse(user), nil
}

// ChangePassword changes a user's password
func (s *Service) ChangePassword(ctx context.Context, id uint, req *ChangePasswordRequest) error {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return ErrUserNotFound
	}

	// Verify current password
	if !password.Verify(req.CurrentPassword, user.Password) {
		return ErrInvalidCredentials
	}

	// Hash new password
	hashedPassword, err := password.Hash(req.NewPassword)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	// Update password
	if err := s.userRepo.UpdatePassword(ctx, id, hashedPassword); err != nil {
		return fmt.Errorf("failed to update password: %w", err)
	}

	return nil
}

// UpdateLastLogin updates the last login timestamp
func (s *Service) UpdateLastLogin(ctx context.Context, id uint) error {
	now := utils.NowUTC()
	return s.userRepo.UpdateLastLogin(ctx, id, now)
}

// ListUsers lists users with pagination
func (s *Service) ListUsers(ctx context.Context, offset, limit int) ([]*UserResponse, error) {
	users, err := s.userRepo.List(ctx, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to list users: %w", err)
	}

	responses := make([]*UserResponse, len(users))
	for i, user := range users {
		responses[i] = s.toUserResponse(user)
	}

	return responses, nil
}

// CountUsers returns the total count of users
func (s *Service) CountUsers(ctx context.Context) (int64, error) {
	return s.userRepo.Count(ctx)
}

// Helper function to convert entity to response
func (s *Service) toUserResponse(user *domain.User) *UserResponse {
	return &UserResponse{
		ID:        user.ID,
		UUID:      user.UUID,
		Email:     user.Email,
		Phone:     user.Phone,
		Name:      user.Name,
		AvatarURL: user.AvatarURL,
		Role:      user.Role,
		Status:    user.Status,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}
}
