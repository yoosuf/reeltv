package domain

import (
	"context"
)

// MyListRepository defines the interface for my list data access
type MyListRepository interface {
	Create(ctx context.Context, item *MyListItem) error
	FindByID(ctx context.Context, id uint) (*MyListItem, error)
	FindByUserAndSeries(ctx context.Context, userID, seriesID uint) (*MyListItem, error)
	FindByUserID(ctx context.Context, userID uint, offset, limit int) ([]*MyListItem, error)
	Delete(ctx context.Context, id uint) error
	DeleteByUserAndSeries(ctx context.Context, userID, seriesID uint) error
}
