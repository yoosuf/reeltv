package application

// UpdateWatchProgressRequest represents a request to update watch progress
type UpdateWatchProgressRequest struct {
	EpisodeID  uint    `json:"episode_id" binding:"required"`
	Duration   int     `json:"duration" binding:"required"`   // seconds watched
	Percentage float32 `json:"percentage" binding:"required"` // 0-100
	Completed  bool    `json:"completed"`
}

// WatchProgressResponse represents watch progress data in responses
type WatchProgressResponse struct {
	ID         uint    `json:"id"`
	UUID       string  `json:"uuid"`
	UserID     uint    `json:"user_id"`
	EpisodeID  uint    `json:"episode_id"`
	WatchedAt  string  `json:"watched_at"`
	Duration   int     `json:"duration"`
	Completed  bool    `json:"completed"`
	Percentage float32 `json:"percentage"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}
