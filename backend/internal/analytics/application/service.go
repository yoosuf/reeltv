package application

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"reeltv/backend/pkg/logger"
	"github.com/rs/zerolog"
)

// Event represents an analytics event
type Event struct {
	EventName  string                 `json:"event_name"`
	UserID     uint                   `json:"user_id,omitempty"`
	Timestamp  time.Time              `json:"timestamp"`
	Properties map[string]interface{} `json:"properties"`
}

// Service implements analytics event tracking
type Service struct {
	logger zerolog.Logger
}

// NewService creates a new analytics service
func NewService() *Service {
	return &Service{
		logger: logger.GetLogger(),
	}
}

// TrackEvent tracks an analytics event
func (s *Service) TrackEvent(ctx context.Context, event *Event) error {
	event.Timestamp = time.Now()

	// Log the event (in production, this would send to an analytics service)
	eventData, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}

	s.logger.Info().
		Str("event", event.EventName).
		Str("data", string(eventData)).
		Msg("Analytics event tracked")

	return nil
}

// TrackVideoView tracks a video view event
func (s *Service) TrackVideoView(ctx context.Context, userID, episodeID uint, duration int) error {
	event := &Event{
		EventName: "video_view",
		UserID:    userID,
		Properties: map[string]interface{}{
			"episode_id": episodeID,
			"duration":   duration,
		},
	}
	return s.TrackEvent(ctx, event)
}

// TrackUserAction tracks a user action event
func (s *Service) TrackUserAction(ctx context.Context, userID uint, action string, properties map[string]interface{}) error {
	event := &Event{
		EventName:  action,
		UserID:     userID,
		Properties: properties,
	}
	return s.TrackEvent(ctx, event)
}
