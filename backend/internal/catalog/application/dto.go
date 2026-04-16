package application

// CreateSeriesRequest represents a request to create a series
type CreateSeriesRequest struct {
	Title       string  `json:"title" binding:"required"`
	Slug        string  `json:"slug" binding:"required"`
	Description string  `json:"description"`
	PosterURL   string  `json:"poster_url"`
	BackdropURL string  `json:"backdrop_url"`
	Year        int     `json:"year" binding:"required"`
	Rating      float32 `json:"rating"`
	Genre       string  `json:"genre"`
	Language    string  `json:"language"`
	ReleaseDate string  `json:"release_date"`
}

// UpdateSeriesRequest represents a request to update a series
type UpdateSeriesRequest struct {
	Title       string  `json:"title" binding:"omitempty"`
	Slug        string  `json:"slug" binding:"omitempty"`
	Description string  `json:"description"`
	PosterURL   string  `json:"poster_url"`
	BackdropURL string  `json:"backdrop_url"`
	Year        int     `json:"year"`
	Rating      float32 `json:"rating"`
	Genre       string  `json:"genre"`
	Language    string  `json:"language"`
	Status      string  `json:"status"`
}

// SeriesResponse represents series data in responses
type SeriesResponse struct {
	ID          uint   `json:"id"`
	UUID        string `json:"uuid"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	PosterURL   string `json:"poster_url"`
	BackdropURL string `json:"backdrop_url"`
	Year        int    `json:"year"`
	Rating      float32 `json:"rating"`
	Genre       string `json:"genre"`
	Language    string `json:"language"`
	Status      string `json:"status"`
	ReleaseDate string `json:"release_date"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// CreateSeasonRequest represents a request to create a season
type CreateSeasonRequest struct {
	SeriesID    uint   `json:"series_id" binding:"required"`
	SeasonNumber int    `json:"season_number" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	PosterURL   string `json:"poster_url"`
	ReleaseDate string `json:"release_date"`
}

// SeasonResponse represents season data in responses
type SeasonResponse struct {
	ID           uint   `json:"id"`
	UUID         string `json:"uuid"`
	SeriesID     uint   `json:"series_id"`
	SeasonNumber int    `json:"season_number"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	PosterURL    string `json:"poster_url"`
	ReleaseDate  string `json:"release_date"`
	EpisodeCount int    `json:"episode_count"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

// CreateEpisodeRequest represents a request to create an episode
type CreateEpisodeRequest struct {
	SeasonID      uint   `json:"season_id" binding:"required"`
	EpisodeNumber int    `json:"episode_number" binding:"required"`
	Title         string `json:"title" binding:"required"`
	Description   string `json:"description"`
	ThumbnailURL  string `json:"thumbnail_url"`
	Duration      int    `json:"duration" binding:"required"`
	VideoURL      string `json:"video_url"`
	ReleaseDate   string `json:"release_date"`
	IsPremium     bool   `json:"is_premium"`
}

// EpisodeResponse represents episode data in responses
type EpisodeResponse struct {
	ID           uint   `json:"id"`
	UUID         string `json:"uuid"`
	SeasonID     uint   `json:"season_id"`
	EpisodeNumber int    `json:"episode_number"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	ThumbnailURL string `json:"thumbnail_url"`
	Duration     int    `json:"duration"`
	VideoURL     string `json:"video_url"`
	ReleaseDate  string `json:"release_date"`
	IsPremium    bool   `json:"is_premium"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
