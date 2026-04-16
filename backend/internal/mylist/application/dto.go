package application

// AddToMyListRequest represents a request to add a series to my list
type AddToMyListRequest struct {
	SeriesID uint `json:"series_id" binding:"required"`
}

// MyListItemResponse represents my list item data in responses
type MyListItemResponse struct {
	ID       uint   `json:"id"`
	UUID     string `json:"uuid"`
	UserID   uint   `json:"user_id"`
	SeriesID uint   `json:"series_id"`
	AddedAt  string `json:"added_at"`
}
