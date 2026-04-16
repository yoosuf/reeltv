package application

import (
	"context"
	"fmt"
	"time"

	sharedDomain "reeltv/backend/internal/shared/domain"
	"reeltv/backend/internal/mylist/domain"
	"reeltv/backend/pkg/utils"
)

var (
	ErrMyListItemNotFound = sharedDomain.ErrNotFound
	ErrAlreadyInMyList    = sharedDomain.ErrAlreadyExists
)

// Service implements my list use cases
type Service struct {
	myListRepo domain.MyListRepository
}

// NewService creates a new my list application service
func NewService(myListRepo domain.MyListRepository) *Service {
	return &Service{
		myListRepo: myListRepo,
	}
}

// AddToMyList adds a series to the user's my list
func (s *Service) AddToMyList(ctx context.Context, userID uint, req *AddToMyListRequest) (*MyListItemResponse, error) {
	// Check if already in my list
	existing, err := s.myListRepo.FindByUserAndSeries(ctx, userID, req.SeriesID)
	if err == nil && existing != nil {
		return nil, ErrAlreadyInMyList
	}

	// Create new my list item
	item := &domain.MyListItem{
		UUIDEntity: sharedDomain.UUIDEntity{UUID: utils.GenerateUUID()},
		UserID:     userID,
		SeriesID:   req.SeriesID,
		AddedAt:    utils.NowUTC(),
	}

	if err := s.myListRepo.Create(ctx, item); err != nil {
		return nil, fmt.Errorf("failed to add to my list: %w", err)
	}

	return s.toMyListItemResponse(item), nil
}

// GetMyList retrieves all items in the user's my list
func (s *Service) GetMyList(ctx context.Context, userID uint, offset, limit int) ([]*MyListItemResponse, error) {
	items, err := s.myListRepo.FindByUserID(ctx, userID, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get my list: %w", err)
	}

	responses := make([]*MyListItemResponse, len(items))
	for i, item := range items {
		responses[i] = s.toMyListItemResponse(item)
	}

	return responses, nil
}

// RemoveFromMyList removes a series from the user's my list
func (s *Service) RemoveFromMyList(ctx context.Context, userID, seriesID uint) error {
	return s.myListRepo.DeleteByUserAndSeries(ctx, userID, seriesID)
}

// Helper function to convert entity to response
func (s *Service) toMyListItemResponse(item *domain.MyListItem) *MyListItemResponse {
	return &MyListItemResponse{
		ID:       item.ID,
		UUID:     item.UUID,
		UserID:   item.UserID,
		SeriesID: item.SeriesID,
		AddedAt:  item.AddedAt.Format(time.RFC3339),
	}
}
