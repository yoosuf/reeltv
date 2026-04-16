package persistence

import (
	"context"

	"gorm.io/gorm"

	"reeltv/backend/internal/mylist/domain"
)

// myListRepository implements MyListRepository interface
type myListRepository struct {
	db *gorm.DB
}

// NewMyListRepository creates a new my list repository
func NewMyListRepository(db *gorm.DB) domain.MyListRepository {
	return &myListRepository{db: db}
}

func (r *myListRepository) Create(ctx context.Context, item *domain.MyListItem) error {
	return r.db.WithContext(ctx).Create(item).Error
}

func (r *myListRepository) FindByID(ctx context.Context, id uint) (*domain.MyListItem, error) {
	var item domain.MyListItem
	err := r.db.WithContext(ctx).Preload("Series").First(&item, id).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *myListRepository) FindByUserAndSeries(ctx context.Context, userID, seriesID uint) (*domain.MyListItem, error) {
	var item domain.MyListItem
	err := r.db.WithContext(ctx).Where("user_id = ? AND series_id = ?", userID, seriesID).First(&item).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *myListRepository) FindByUserID(ctx context.Context, userID uint, offset, limit int) ([]*domain.MyListItem, error) {
	var items []*domain.MyListItem
	err := r.db.WithContext(ctx).Preload("Series").Where("user_id = ?", userID).Order("added_at DESC").Offset(offset).Limit(limit).Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (r *myListRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.MyListItem{}, id).Error
}

func (r *myListRepository) DeleteByUserAndSeries(ctx context.Context, userID, seriesID uint) error {
	return r.db.WithContext(ctx).Where("user_id = ? AND series_id = ?", userID, seriesID).Delete(&domain.MyListItem{}).Error
}
