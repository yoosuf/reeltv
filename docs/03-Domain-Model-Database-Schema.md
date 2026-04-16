# Reel TV Domain Model and Database Schema

## Domain Model Overview

The Reel TV domain is centered around content catalog management, user interactions, and playback tracking. The core entities are Series, Seasons, Episodes, and Users, with supporting entities for genres, tags, watch progress, favorites, and subscriptions.

## Entity Relationship Diagram

```
┌─────────────┐       ┌──────────────┐       ┌─────────────┐
│    User     │──1:N──│WatchProgress │──N:1──│  Episode    │
└─────────────┘       └──────────────┘       └──────┬──────┘
       │                                            │
       │1:N                                        │N:1
       │                                            │
┌──────▼──────┐                             ┌──────▼──────┐
│   MyList    │                             │   Season    │
└─────────────┘                             └──────┬──────┘
                                                 │
                                                 │N:1
                                                 │
                                          ┌──────▼──────┐
                                          │   Series    │──N:M──┌─────────┐
                                          └─────────────┘        │ Genre   │
                                                 │              └─────────┘
                                                 │N:M
                                                 │
                                          ┌──────▼──────┐
                                          │    Tag      │
                                          └─────────────┘

┌─────────────┐       ┌──────────────┐
│    User     │──1:N──│ Subscription  │
└─────────────┘       └──────────────┘

┌─────────────┐       ┌──────────────┐
│    User     │──1:N──│RefreshToken  │
└─────────────┘       └──────────────┘

┌─────────────┐       ┌──────────────┐
│    User     │──1:N──│AnalyticsEvent│
└─────────────┘       └──────────────┘
```

## GORM Model Definitions

### 1. User Model

```go
package model

import (
    "time"
    "gorm.io/gorm"
)

type Role string

const (
    RoleUser  Role = "user"
    RoleAdmin Role = "admin"
)

type User struct {
    ID        uint           `gorm:"primaryKey" json:"id"`
    UUID      string         `gorm:"type:varchar(36);uniqueIndex;not null" json:"uuid"`
    Email     string         `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
    Phone     string         `gorm:"type:varchar(20);uniqueIndex" json:"phone"`
    Password  string         `gorm:"type:varchar(255);not null" json:"-"`
    Name      string         `gorm:"type:varchar(255)" json:"name"`
    AvatarURL string         `gorm:"type:varchar(500)" json:"avatar_url"`
    Role      Role           `gorm:"type:varchar(20);default:user;not null" json:"role"`
    IsActive  bool           `gorm:"default:true;not null" json:"is_active"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

    // Relationships
    WatchProgresses   []WatchProgress   `gorm:"foreignKey:UserID" json:"-"`
    MyList            []MyList          `gorm:"foreignKey:UserID" json:"-"`
    Subscriptions     []Subscription    `gorm:"foreignKey:UserID" json:"-"`
    RefreshTokens     []RefreshToken    `gorm:"foreignKey:UserID" json:"-"`
    AnalyticsEvents   []AnalyticsEvent  `gorm:"foreignKey:UserID" json:"-"`
}
```

### 2. RefreshToken Model (JWT Refresh Tokens)

```go
type RefreshToken struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    Token     string    `gorm:"type:varchar(500);uniqueIndex;not null" json:"-"`
    UserID    uint      `gorm:"not null;index" json:"user_id"`
    ExpiresAt time.Time `gorm:"not null;index" json:"expires_at"`
    CreatedAt time.Time `json:"created_at"`
    
    // Relationship
    User User `gorm:"foreignKey:UserID" json:"-"`
}
```

### 3. Genre Model

```go
type Genre struct {
    ID          uint           `gorm:"primaryKey" json:"id"`
    Name        string         `gorm:"type:varchar(100);uniqueIndex;not null" json:"name"`
    Slug        string         `gorm:"type:varchar(100);uniqueIndex;not null" json:"slug"`
    Description string         `gorm:"type:text" json:"description"`
    IsActive    bool           `gorm:"default:true;not null" json:"is_active"`
    DisplayOrder int           `gorm:"default:0" json:"display_order"`
    CreatedAt   time.Time      `json:"created_at"`
    UpdatedAt   time.Time      `json:"updated_at"`
    DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

    // Relationships
    Series []Series `gorm:"many2many:series_genres;" json:"-"`
}
```

### 4. Tag Model

```go
type Tag struct {
    ID        uint           `gorm:"primaryKey" json:"id"`
    Name      string         `gorm:"type:varchar(100);uniqueIndex;not null" json:"name"`
    Slug      string         `gorm:"type:varchar(100);uniqueIndex;not null" json:"slug"`
    IsActive  bool           `gorm:"default:true;not null" json:"is_active"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

    // Relationships
    Series []Series `gorm:"many2many:series_tags;" json:"-"`
}
```

### 5. Series Model

```go
type SeriesStatus string

const (
    SeriesStatusDraft     SeriesStatus = "draft"
    SeriesStatusPublished SeriesStatus = "published"
    SeriesStatusArchived  SeriesStatus = "archived"
)

type Series struct {
    ID               uint           `gorm:"primaryKey" json:"id"`
    UUID             string         `gorm:"type:varchar(36);uniqueIndex;not null" json:"uuid"`
    Title            string         `gorm:"type:varchar(255);not null;index" json:"title"`
    Slug             string         `gorm:"type:varchar(255);uniqueIndex;not null" json:"slug"`
    Description      string         `gorm:"type:text" json:"description"`
    PosterURL        string         `gorm:"type:varchar(500)" json:"poster_url"`
    BackdropURL      string         `gorm:"type:varchar(500)" json:"backdrop_url"`
    Year             int            `gorm:"index" json:"year"`
    Status           SeriesStatus   `gorm:"type:varchar(20);default:draft;not null;index" json:"status"`
    IsPremium        bool           `gorm:"default:false;not null;index" json:"is_premium"`
    TotalEpisodes    int            `gorm:"default:0" json:"total_episodes"`
    TotalSeasons     int            `gorm:"default:0" json:"total_seasons"`
    Rating           float32        `gorm:"default:0" json:"rating"`
    ViewCount        int64          `gorm:"default:0" json:"view_count"`
    ReleasedAt       *time.Time     `gorm:"index" json:"released_at"`
    CreatedAt        time.Time      `json:"created_at"`
    UpdatedAt        time.Time      `json:"updated_at"`
    DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`

    // Relationships
    Seasons    []Season       `gorm:"foreignKey:SeriesID" json:"seasons,omitempty"`
    Genres     []Genre        `gorm:"many2many:series_genres;" json:"genres,omitempty"`
    Tags       []Tag          `gorm:"many2many:series_tags;" json:"tags,omitempty"`
    MyList     []MyList       `gorm:"foreignKey:SeriesID" json:"-"`
}
```

### 6. Season Model

```go
type Season struct {
    ID          uint           `gorm:"primaryKey" json:"id"`
    UUID        string         `gorm:"type:varchar(36);uniqueIndex;not null" json:"uuid"`
    SeriesID    uint           `gorm:"not null;index" json:"series_id"`
    SeasonNumber int           `gorm:"not null" json:"season_number"`
    Title       string         `gorm:"type:varchar(255)" json:"title"`
    Description string         `gorm:"type:text" json:"description"`
    PosterURL   string         `gorm:"type:varchar(500)" json:"poster_url"`
    EpisodeCount int           `gorm:"default:0" json:"episode_count"`
    ReleasedAt  *time.Time     `gorm:"index" json:"released_at"`
    CreatedAt   time.Time      `json:"created_at"`
    UpdatedAt   time.Time      `json:"updated_at"`
    DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

    // Relationships
    Series  Series    `gorm:"foreignKey:SeriesID" json:"series,omitempty"`
    Episodes []Episode `gorm:"foreignKey:SeasonID" json:"episodes,omitempty"`
}
```

### 7. Episode Model

```go
type EpisodeStatus string

const (
    EpisodeStatusDraft     EpisodeStatus = "draft"
    EpisodeStatusPublished EpisodeStatus = "published"
    EpisodeStatusArchived  EpisodeStatus = "archived"
)

type Episode struct {
    ID              uint           `gorm:"primaryKey" json:"id"`
    UUID            string         `gorm:"type:varchar(36);uniqueIndex;not null" json:"uuid"`
    SeriesID        uint           `gorm:"not null;index" json:"series_id"`
    SeasonID        uint           `gorm:"not null;index" json:"season_id"`
    EpisodeNumber   int            `gorm:"not null" json:"episode_number"`
    Title           string         `gorm:"type:varchar(255);not null" json:"title"`
    Description     string         `gorm:"type:text" json:"description"`
    ThumbnailURL    string         `gorm:"type:varchar(500)" json:"thumbnail_url"`
    VideoURL        string         `gorm:"type:varchar(500);not null" json:"video_url"`
    Duration        int            `gorm:"not null" json:"duration"` // in seconds
    Status          EpisodeStatus  `gorm:"type:varchar(20);default:draft;not null;index" json:"status"`
    IsPremium       bool           `gorm:"default:false;not null;index" json:"is_premium"`
    ViewCount       int64          `gorm:"default:0" json:"view_count"`
    ReleasedAt      *time.Time     `gorm:"index" json:"released_at"`
    CreatedAt       time.Time      `json:"created_at"`
    UpdatedAt       time.Time      `json:"updated_at"`
    DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`

    // Relationships
    Series         Series         `gorm:"foreignKey:SeriesID" json:"series,omitempty"`
    Season         Season         `gorm:"foreignKey:SeasonID" json:"season,omitempty"`
    WatchProgresses []WatchProgress `gorm:"foreignKey:EpisodeID" json:"-"`
}
```

### 8. WatchProgress Model

```go
type WatchProgress struct {
    ID          uint      `gorm:"primaryKey" json:"id"`
    UserID      uint      `gorm:"not null;index:idx_user_episode" json:"user_id"`
    EpisodeID   uint      `gorm:"not null;index:idx_user_episode;index" json:"episode_id"`
    Progress    int       `gorm:"not null" json:"progress"` // in seconds
    Completed   bool      `gorm:"default:false;not null;index" json:"completed"`
    LastWatched time.Time `gorm:"not null;index" json:"last_watched"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`

    // Relationships
    User    User    `gorm:"foreignKey:UserID" json:"-"`
    Episode Episode `gorm:"foreignKey:EpisodeID" json:"episode,omitempty"`
}
```

### 9. MyList Model (Favorites)

```go
type MyList struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    UserID    uint      `gorm:"not null;index:idx_user_series" json:"user_id"`
    SeriesID  uint      `gorm:"not null;index:idx_user_series" json:"series_id"`
    AddedAt   time.Time `gorm:"not null;index" json:"added_at"`
    CreatedAt time.Time `json:"created_at"`

    // Relationships
    User   User   `gorm:"foreignKey:UserID" json:"-"`
    Series Series `gorm:"foreignKey:SeriesID" json:"series,omitempty"`
}
```

### 10. Subscription Model

```go
type SubscriptionStatus string

const (
    SubscriptionStatusActive   SubscriptionStatus = "active"
    SubscriptionStatusExpired  SubscriptionStatus = "expired"
    SubscriptionStatusCanceled SubscriptionStatus = "canceled"
    SubscriptionStatusPending  SubscriptionStatus = "pending"
)

type SubscriptionPlan string

const (
    SubscriptionPlanFree  SubscriptionPlan = "free"
    SubscriptionPlanBasic SubscriptionPlan = "basic"
    SubscriptionPlanPro   SubscriptionPlan = "pro"
)

type Subscription struct {
    ID             uint              `gorm:"primaryKey" json:"id"`
    UUID           string            `gorm:"type:varchar(36);uniqueIndex;not null" json:"uuid"`
    UserID         uint              `gorm:"not null;index" json:"user_id"`
    Plan           SubscriptionPlan  `gorm:"type:varchar(20);not null" json:"plan"`
    Status         SubscriptionStatus `gorm:"type:varchar(20);not null;index" json:"status"`
    StartDate      time.Time         `gorm:"not null" json:"start_date"`
    EndDate        *time.Time        `gorm:"index" json:"end_date"`
    AutoRenew      bool              `gorm:"default:true" json:"auto_renew"`
    ExternalID     string            `gorm:"type:varchar(255);index" json:"external_id"` // Payment gateway ID
    CreatedAt      time.Time         `json:"created_at"`
    UpdatedAt      time.Time         `json:"updated_at"`

    // Relationships
    User User `gorm:"foreignKey:UserID" json:"-"`
}
```

### 11. AnalyticsEvent Model

```go
type EventType string

const (
    EventTypePageView       EventType = "page_view"
    EventTypeVideoStart     EventType = "video_start"
    EventTypeVideoComplete EventType = "video_complete"
    EventTypeVideoPause     EventType = "video_pause"
    EventTypeSearch         EventType = "search"
    EventTypeAddToList      EventType = "add_to_list"
    EventTypeRemoveFromList EventType = "remove_from_list"
    EventTypeSubscription  EventType = "subscription"
)

type AnalyticsEvent struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    UserID    uint      `gorm:"index" json:"user_id"`
    EventType EventType `gorm:"type:varchar(50);not null;index" json:"event_type"`
    EventData string    `gorm:"type:jsonb" json:"event_data"` // JSON payload
    Timestamp time.Time `gorm:"not null;index" json:"timestamp"`
    UserAgent string    `gorm:"type:varchar(500)" json:"user_agent"`
    IPAddress string    `gorm:"type:varchar(45)" json:"ip_address"`
    CreatedAt time.Time `json:"created_at"`

    // Relationships
    User User `gorm:"foreignKey:UserID" json:"-"`
}
```

## Database Schema (DDL)

### Indexes

**Users Table**
- PRIMARY KEY: `id`
- UNIQUE: `uuid`, `email`, `phone`
- INDEX: `deleted_at`

**RefreshTokens Table**
- PRIMARY KEY: `id`
- UNIQUE: `token`
- INDEX: `user_id`, `expires_at`

**Genres Table**
- PRIMARY KEY: `id`
- UNIQUE: `name`, `slug`
- INDEX: `deleted_at`

**Tags Table**
- PRIMARY KEY: `id`
- UNIQUE: `name`, `slug`
- INDEX: `deleted_at`

**Series Table**
- PRIMARY KEY: `id`
- UNIQUE: `uuid`, `slug`
- INDEX: `title`, `status`, `is_premium`, `year`, `released_at`, `deleted_at`

**Seasons Table**
- PRIMARY KEY: `id`
- UNIQUE: `uuid`
- INDEX: `series_id`, `released_at`, `deleted_at`

**Episodes Table**
- PRIMARY KEY: `id`
- UNIQUE: `uuid`
- INDEX: `series_id`, `season_id`, `status`, `is_premium`, `released_at`, `deleted_at`

**WatchProgress Table**
- PRIMARY KEY: `id`
- UNIQUE INDEX: `idx_user_episode` (`user_id`, `episode_id`)
- INDEX: `episode_id`, `last_watched`

**MyList Table**
- PRIMARY KEY: `id`
- UNIQUE INDEX: `idx_user_series` (`user_id`, `series_id`)
- INDEX: `added_at`

**Subscriptions Table**
- PRIMARY KEY: `id`
- UNIQUE: `uuid`
- INDEX: `user_id`, `status`, `end_date`, `external_id`

**AnalyticsEvents Table**
- PRIMARY KEY: `id`
- INDEX: `user_id`, `event_type`, `timestamp`

**SeriesGenres (Join Table)**
- INDEX: `series_id`, `genre_id`

**SeriesTags (Join Table)**
- INDEX: `series_id`, `tag_id`

## Foreign Key Relationships

```sql
-- WatchProgress
ALTER TABLE watch_progresses 
ADD CONSTRAINT fk_watch_progress_user 
FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;

ALTER TABLE watch_progresses 
ADD CONSTRAINT fk_watch_progress_episode 
FOREIGN KEY (episode_id) REFERENCES episodes(id) ON DELETE CASCADE;

-- MyList
ALTER TABLE my_lists 
ADD CONSTRAINT fk_my_list_user 
FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;

ALTER TABLE my_lists 
ADD CONSTRAINT fk_my_list_series 
FOREIGN KEY (series_id) REFERENCES series(id) ON DELETE CASCADE;

-- Season
ALTER TABLE seasons 
ADD CONSTRAINT fk_season_series 
FOREIGN KEY (series_id) REFERENCES series(id) ON DELETE CASCADE;

-- Episode
ALTER TABLE episodes 
ADD CONSTRAINT fk_episode_series 
FOREIGN KEY (series_id) REFERENCES series(id) ON DELETE CASCADE;

ALTER TABLE episodes 
ADD CONSTRAINT fk_episode_season 
FOREIGN KEY (season_id) REFERENCES seasons(id) ON DELETE CASCADE;

-- RefreshToken
ALTER TABLE refresh_tokens 
ADD CONSTRAINT fk_refresh_token_user 
FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;

-- Subscription
ALTER TABLE subscriptions 
ADD CONSTRAINT fk_subscription_user 
FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;

-- AnalyticsEvent
ALTER TABLE analytics_events 
ADD CONSTRAINT fk_analytics_event_user 
FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL;
```

## Data Types and Constraints

### String Lengths
- UUID: 36 characters (standard UUID format)
- Email: 255 characters
- Phone: 20 characters (international format)
- Password: 255 characters (bcrypt hash)
- Title: 255 characters
- Slug: 255 characters
- URL fields: 500 characters
- Token: 500 characters

### Numeric Types
- IDs: unsigned integer (auto-increment)
- Duration: integer (seconds)
- Episode/Season numbers: integer
- View counts: bigint (int64)
- Rating: float32 (0.0 - 10.0)
- Progress: integer (seconds)

### Boolean Flags
- `is_active`: User account status
- `is_premium`: Content access level
- `completed`: Episode completion status
- `auto_renew`: Subscription auto-renewal

### Timestamps
- `created_at`: Record creation time
- `updated_at`: Last update time
- `deleted_at`: Soft delete timestamp (GORM)
- `released_at`: Content release date
- `last_watched`: Last watch timestamp
- `expires_at`: Token expiration
- `start_date`/`end_date`: Subscription period

### Enums
- `Role`: user, admin
- `SeriesStatus`: draft, published, archived
- `EpisodeStatus`: draft, published, archived
- `SubscriptionStatus`: active, expired, canceled, pending
- `SubscriptionPlan`: free, basic, pro
- `EventType`: Various analytics event types

## JSONB Usage

### AnalyticsEvent.EventData
Stores flexible JSON payload for different event types:
```json
{
  "episode_id": 123,
  "series_id": 456,
  "progress": 45,
  "duration": 120,
  "device_type": "mobile",
  "connection_type": "wifi"
}
```

## Soft Deletes

Tables with soft delete capability (using `gorm.DeletedAt`):
- Users
- Genres
- Tags
- Series
- Seasons
- Episodes

These tables retain deleted records for audit purposes and can be restored if needed.

## Cascade Delete Rules

- **CASCADE**: WatchProgress, MyList, RefreshToken (when user is deleted)
- **CASCADE**: Episodes (when season is deleted)
- **CASCADE**: Seasons (when series is deleted)
- **SET NULL**: AnalyticsEvent (when user is deleted - keep analytics data)

## Database Optimization Notes

### Indexing Strategy
1. Foreign keys are indexed for JOIN performance
2. Frequently queried fields (status, is_premium) have indexes
3. Composite indexes for common query patterns (user_id + episode_id)
4. Timestamp indexes for time-based queries (continue watching, new releases)

### Query Patterns
- **Catalog browsing**: Filter by status, genre, is_premium
- **Continue watching**: Filter by user_id, last_watched DESC, completed=false
- **My list**: Filter by user_id, order by added_at DESC
- **Search**: Full-text search on title (PostgreSQL tsvector, future enhancement)
- **Recommendations**: Complex joins across watch history, genres, tags

### Partitioning (Future)
Consider partitioning large tables by:
- AnalyticsEvents by month (time-based partitioning)
- WatchProgress by user_id hash (if table grows large)

## Migration Strategy

GORM AutoMigrate will be used for initial schema creation. For production:
1. Versioned migration files in `/migrations` directory
2. Up/down migration scripts
3. Migration tracking table
4. Rollback capability

## Seed Data Requirements

Initial seed data will include:
- 1 admin user
- 5-10 test users
- 10-15 genres (Drama, Romance, Thriller, Comedy, etc.)
- 20-30 tags
- 20-30 series (mix of free and premium)
- 50-100 episodes across series
- Sample watch progress data
- Sample my list entries
- Sample subscriptions

This schema supports all MVP features and provides a solid foundation for future enhancements.
