# Reel TV API Specification

## Base URL

- **Development**: `http://localhost:8080/api/v1`
- **Production**: `https://api.reeltv.com/api/v1`

## Authentication

Most endpoints require JWT authentication. Include the access token in the `Authorization` header:

```
Authorization: Bearer <access_token>
```

## Standard Response Format

### Success Response
```json
{
  "success": true,
  "data": { /* response data */ },
  "message": "Operation successful",
  "request_id": "uuid"
}
```

### Error Response
```json
{
  "success": false,
  "error": {
    "code": "ERROR_CODE",
    "message": "Human-readable error message",
    "details": { /* additional error details */ }
  },
  "request_id": "uuid"
}
```

### Error Codes
- `INVALID_REQUEST`: Invalid request parameters
- `UNAUTHORIZED`: Authentication required or invalid token
- `FORBIDDEN`: Insufficient permissions
- `NOT_FOUND`: Resource not found
- `CONFLICT`: Resource conflict (e.g., duplicate email)
- `INTERNAL_ERROR`: Internal server error
- `SERVICE_UNAVAILABLE`: Service temporarily unavailable
- `RATE_LIMIT_EXCEEDED`: Too many requests

## Pagination

List endpoints support pagination via query parameters:

- `page`: Page number (default: 1)
- `limit`: Items per page (default: 20, max: 100)
- `cursor`: Cursor for cursor-based pagination (where applicable)

**Paginated Response Format:**
```json
{
  "success": true,
  "data": {
    "items": [ /* array of items */ ],
    "pagination": {
      "page": 1,
      "limit": 20,
      "total": 100,
      "total_pages": 5,
      "has_next": true,
      "has_prev": false
    }
  }
}
```

---

## API Endpoints

### 1. Authentication

#### 1.1 Register
**POST** `/auth/register`

Register a new user account.

**Request Body:**
```json
{
  "email": "user@example.com",
  "phone": "+1234567890",
  "password": "SecurePassword123",
  "name": "John Doe"
}
```

**Response (201):**
```json
{
  "success": true,
  "data": {
    "user": {
      "id": 1,
      "uuid": "uuid",
      "email": "user@example.com",
      "phone": "+1234567890",
      "name": "John Doe",
      "avatar_url": null,
      "role": "user",
      "created_at": "2024-01-01T00:00:00Z"
    },
    "access_token": "jwt_access_token",
    "refresh_token": "jwt_refresh_token",
    "token_type": "Bearer",
    "expires_in": 900
  }
}
```

**Validation:**
- `email` or `phone` required (at least one)
- `password`: min 8 characters
- `name`: required if provided

#### 1.2 Login
**POST** `/auth/login`

Authenticate a user.

**Request Body:**
```json
{
  "email": "user@example.com",
  "phone": "+1234567890",
  "password": "SecurePassword123"
}
```

**Response (200):**
```json
{
  "success": true,
  "data": {
    "user": {
      "id": 1,
      "uuid": "uuid",
      "email": "user@example.com",
      "phone": "+1234567890",
      "name": "John Doe",
      "avatar_url": null,
      "role": "user"
    },
    "access_token": "jwt_access_token",
    "refresh_token": "jwt_refresh_token",
    "token_type": "Bearer",
    "expires_in": 900
  }
}
```

#### 1.3 Refresh Token
**POST** `/auth/refresh`

Refresh access token using refresh token.

**Request Body:**
```json
{
  "refresh_token": "jwt_refresh_token"
}
```

**Response (200):**
```json
{
  "success": true,
  "data": {
    "access_token": "new_jwt_access_token",
    "refresh_token": "new_jwt_refresh_token",
    "token_type": "Bearer",
    "expires_in": 900
  }
}
```

#### 1.4 Logout
**POST** `/auth/logout`
**Authentication Required**

Invalidate the current session.

**Request Headers:**
```
Authorization: Bearer <access_token>
```

**Response (200):**
```json
{
  "success": true,
  "message": "Logged out successfully"
}
```

---

### 2. User Profile

#### 2.1 Get Profile
**GET** `/users/me`
**Authentication Required**

Get current user profile.

**Response (200):**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "uuid": "uuid",
    "email": "user@example.com",
    "phone": "+1234567890",
    "name": "John Doe",
    "avatar_url": null,
    "role": "user",
    "is_active": true,
    "created_at": "2024-01-01T00:00:00Z",
    "subscription": {
      "plan": "basic",
      "status": "active",
      "start_date": "2024-01-01T00:00:00Z",
      "end_date": "2024-02-01T00:00:00Z",
      "auto_renew": true
    }
  }
}
```

#### 2.2 Update Profile
**PUT** `/users/me`
**Authentication Required**

Update user profile.

**Request Body:**
```json
{
  "name": "John Smith",
  "avatar_url": "https://example.com/avatar.jpg"
}
```

**Response (200):**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "uuid": "uuid",
    "email": "user@example.com",
    "phone": "+1234567890",
    "name": "John Smith",
    "avatar_url": "https://example.com/avatar.jpg",
    "role": "user",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

#### 2.3 Change Password
**POST** `/users/me/password`
**Authentication Required**

Change user password.

**Request Body:**
```json
{
  "current_password": "OldPassword123",
  "new_password": "NewPassword123"
}
```

**Response (200):**
```json
{
  "success": true,
  "message": "Password changed successfully"
}
```

---

### 3. Catalog - Series

#### 3.1 List Series
**GET** `/series`

List all published series with filtering and pagination.

**Query Parameters:**
- `page`: Page number (default: 1)
- `limit`: Items per page (default: 20, max: 100)
- `genre_id`: Filter by genre ID
- `tag_id`: Filter by tag ID
- `is_premium`: Filter by premium status (true/false)
- `status`: Filter by status (default: published)
- `sort`: Sort order (trending, newest, rating, title)
- `search`: Search in title

**Response (200):**
```json
{
  "success": true,
  "data": {
    "items": [
      {
        "id": 1,
        "uuid": "uuid",
        "title": "Drama Series Title",
        "slug": "drama-series-title",
        "description": "Series description",
        "poster_url": "https://example.com/poster.jpg",
        "backdrop_url": "https://example.com/backdrop.jpg",
        "year": 2024,
        "is_premium": false,
        "total_episodes": 24,
        "total_seasons": 2,
        "rating": 8.5,
        "view_count": 10000,
        "released_at": "2024-01-01T00:00:00Z",
        "genres": [
          {
            "id": 1,
            "name": "Drama",
            "slug": "drama"
          }
        ],
        "tags": [
          {
            "id": 1,
            "name": "romance",
            "slug": "romance"
          }
        ]
      }
    ],
    "pagination": {
      "page": 1,
      "limit": 20,
      "total": 100,
      "total_pages": 5,
      "has_next": true,
      "has_prev": false
    }
  }
}
```

#### 3.2 Get Series Detail
**GET** `/series/{id}` or `/series/{slug}`

Get detailed information about a series including seasons.

**Path Parameters:**
- `id`: Series ID or slug

**Response (200):**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "uuid": "uuid",
    "title": "Drama Series Title",
    "slug": "drama-series-title",
    "description": "Full series description",
    "poster_url": "https://example.com/poster.jpg",
    "backdrop_url": "https://example.com/backdrop.jpg",
    "year": 2024,
    "is_premium": false,
    "total_episodes": 24,
    "total_seasons": 2,
    "rating": 8.5,
    "view_count": 10000,
    "released_at": "2024-01-01T00:00:00Z",
    "genres": [
      {
        "id": 1,
        "name": "Drama",
        "slug": "drama"
      }
    ],
    "tags": [
      {
        "id": 1,
        "name": "romance",
        "slug": "romance"
      }
    ],
    "seasons": [
      {
        "id": 1,
        "uuid": "uuid",
        "season_number": 1,
        "title": "Season 1",
        "description": "Season description",
        "poster_url": "https://example.com/season1.jpg",
        "episode_count": 12,
        "released_at": "2024-01-01T00:00:00Z"
      }
    ],
    "in_my_list": false
  }
}
```

#### 3.3 Get Season Detail
**GET** `/series/{series_id}/seasons/{season_number}`

Get episodes in a specific season.

**Path Parameters:**
- `series_id`: Series ID or slug
- `season_number`: Season number

**Response (200):**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "uuid": "uuid",
    "series_id": 1,
    "season_number": 1,
    "title": "Season 1",
    "description": "Season description",
    "poster_url": "https://example.com/season1.jpg",
    "episode_count": 12,
    "released_at": "2024-01-01T00:00:00Z",
    "episodes": [
      {
        "id": 1,
        "uuid": "uuid",
        "series_id": 1,
        "season_id": 1,
        "episode_number": 1,
        "title": "Episode 1",
        "description": "Episode description",
        "thumbnail_url": "https://example.com/ep1.jpg",
        "duration": 120,
        "is_premium": false,
        "released_at": "2024-01-01T00:00:00Z",
        "watch_progress": {
          "progress": 45,
          "completed": false
        }
      }
    ]
  }
}
```

#### 3.4 Get Episode Detail
**GET** `/episodes/{id}` or `/episodes/{uuid}`

Get detailed information about an episode.

**Path Parameters:**
- `id`: Episode ID or UUID

**Query Parameters:**
- `include_series`: Include series information (default: false)

**Response (200):**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "uuid": "uuid",
    "series_id": 1,
    "season_id": 1,
    "episode_number": 1,
    "title": "Episode 1",
    "description": "Episode description",
    "thumbnail_url": "https://example.com/ep1.jpg",
    "video_url": "https://s3.example.com/videos/ep1.mp4",
    "duration": 120,
    "is_premium": false,
    "released_at": "2024-01-01T00:00:00Z",
    "series": {
      "id": 1,
      "title": "Drama Series Title",
      "slug": "drama-series-title",
      "poster_url": "https://example.com/poster.jpg"
    },
    "season": {
      "id": 1,
      "season_number": 1,
      "title": "Season 1"
    },
    "watch_progress": {
      "progress": 45,
      "completed": false
    },
    "can_play": true,
    "lock_reason": null
  }
}
```

**Lock Reason Values:**
- `null`: Episode can be played
- `premium`: Requires subscription
- `login`: Requires authentication

---

### 4. Home Feed

#### 4.1 Get Home Feed
**GET** `/home`

Get personalized home feed with multiple sections.

**Authentication Required** (Optional - returns generic feed if not authenticated)

**Response (200):**
```json
{
  "success": true,
  "data": {
    "continue_watching": [
      {
        "episode": {
          "id": 1,
          "uuid": "uuid",
          "title": "Episode 1",
          "thumbnail_url": "https://example.com/ep1.jpg",
          "duration": 120,
          "is_premium": false
        },
        "series": {
          "id": 1,
          "title": "Drama Series Title",
          "poster_url": "https://example.com/poster.jpg"
        },
        "progress": 45,
        "last_watched": "2024-01-01T00:00:00Z"
      }
    ],
    "trending": [
      {
        "id": 1,
        "title": "Drama Series Title",
        "poster_url": "https://example.com/poster.jpg",
        "rating": 8.5,
        "view_count": 10000
      }
    ],
    "new_releases": [
      {
        "id": 2,
        "title": "New Series",
        "poster_url": "https://example.com/poster2.jpg",
        "released_at": "2024-01-01T00:00:00Z"
      }
    ],
    "recommended_for_you": [
      {
        "id": 3,
        "title": "Recommended Series",
        "poster_url": "https://example.com/poster3.jpg",
        "reason": "Because you watched Drama Series Title"
      }
    ],
    "genres": [
      {
        "id": 1,
        "name": "Drama",
        "slug": "drama",
        "poster_url": "https://example.com/drama.jpg",
        "series_count": 50
      }
    ]
  }
}
```

---

### 5. Watch Progress

#### 5.1 Update Watch Progress
**POST** `/watch-progress`
**Authentication Required**

Update or create watch progress for an episode.

**Request Body:**
```json
{
  "episode_id": 1,
  "progress": 45,
  "completed": false
}
```

**Response (200):**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "user_id": 1,
    "episode_id": 1,
    "progress": 45,
    "completed": false,
    "last_watched": "2024-01-01T00:00:00Z"
  }
}
```

#### 5.2 Get Continue Watching
**GET** `/watch-progress/continue-watching`
**Authentication Required**

Get list of episodes in progress.

**Query Parameters:**
- `limit`: Max items (default: 20)

**Response (200):**
```json
{
  "success": true,
  "data": [
    {
      "episode": {
        "id": 1,
        "uuid": "uuid",
        "title": "Episode 1",
        "thumbnail_url": "https://example.com/ep1.jpg",
        "duration": 120,
        "is_premium": false
      },
      "series": {
        "id": 1,
        "title": "Drama Series Title",
        "poster_url": "https://example.com/poster.jpg"
      },
      "progress": 45,
      "completed": false,
      "last_watched": "2024-01-01T00:00:00Z"
    }
  ]
}
```

#### 5.3 Get Episode Progress
**GET** `/watch-progress/episode/{episode_id}`
**Authentication Required**

Get watch progress for a specific episode.

**Response (200):**
```json
{
  "success": true,
  "data": {
    "progress": 45,
    "completed": false,
    "last_watched": "2024-01-01T00:00:00Z"
  }
}
```

---

### 6. My List (Favorites)

#### 6.1 Add to My List
**POST** `/my-list`
**Authentication Required**

Add a series to user's my list.

**Request Body:**
```json
{
  "series_id": 1
}
```

**Response (201):**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "user_id": 1,
    "series_id": 1,
    "added_at": "2024-01-01T00:00:00Z"
  }
}
```

#### 6.2 Remove from My List
**DELETE** `/my-list/{series_id}`
**Authentication Required**

Remove a series from user's my list.

**Response (200):**
```json
{
  "success": true,
  "message": "Removed from my list"
}
```

#### 6.3 Get My List
**GET** `/my-list`
**Authentication Required**

Get user's my list.

**Query Parameters:**
- `page`: Page number (default: 1)
- `limit`: Items per page (default: 20, max: 100)

**Response (200):**
```json
{
  "success": true,
  "data": {
    "items": [
      {
        "id": 1,
        "series_id": 1,
        "added_at": "2024-01-01T00:00:00Z",
        "series": {
          "id": 1,
          "title": "Drama Series Title",
          "poster_url": "https://example.com/poster.jpg",
          "is_premium": false,
          "total_episodes": 24,
          "rating": 8.5
        }
      }
    ],
    "pagination": {
      "page": 1,
      "limit": 20,
      "total": 10,
      "total_pages": 1
    }
  }
}
```

#### 6.4 Check if in My List
**GET** `/my-list/check/{series_id}`
**Authentication Required**

Check if a series is in user's my list.

**Response (200):**
```json
{
  "success": true,
  "data": {
    "in_list": true
  }
}
```

---

### 7. Search

#### 7.1 Search Series
**GET** `/search`

Search for series by title, genre, or tags.

**Query Parameters:**
- `q`: Search query
- `page`: Page number (default: 1)
- `limit`: Items per page (default: 20, max: 100)
- `genre_id`: Filter by genre
- `is_premium`: Filter by premium status

**Response (200):**
```json
{
  "success": true,
  "data": {
    "items": [
      {
        "id": 1,
        "title": "Drama Series Title",
        "slug": "drama-series-title",
        "poster_url": "https://example.com/poster.jpg",
        "year": 2024,
        "is_premium": false,
        "rating": 8.5,
        "genres": [
          {
            "id": 1,
            "name": "Drama",
            "slug": "drama"
          }
        ]
      }
    ],
    "pagination": {
      "page": 1,
      "limit": 20,
      "total": 10,
      "total_pages": 1
    }
  }
}
```

---

### 8. Recommendations

#### 8.1 Get Recommendations
**GET** `/recommendations`
**Authentication Required**

Get personalized recommendations.

**Query Parameters:**
- `limit`: Max items (default: 20)

**Response (200):**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "title": "Drama Series Title",
      "poster_url": "https://example.com/poster.jpg",
      "is_premium": false,
      "rating": 8.5,
      "reason": "Because you watched Similar Series"
    }
  ]
}
```

---

### 9. Subscription / Entitlement

#### 9.1 Get Subscription Status
**GET** `/subscription`
**Authentication Required**

Get current subscription status.

**Response (200):**
```json
{
  "success": true,
  "data": {
    "uuid": "uuid",
    "plan": "basic",
    "status": "active",
    "start_date": "2024-01-01T00:00:00Z",
    "end_date": "2024-02-01T00:00:00Z",
    "auto_renew": true,
    "can_access_premium": true
  }
}
```

#### 9.2 Check Episode Access
**POST** `/subscription/check-access`
**Authentication Required**

Check if user can access a specific episode.

**Request Body:**
```json
{
  "episode_id": 1
}
```

**Response (200):**
```json
{
  "success": true,
  "data": {
    "can_access": true,
    "reason": null
  }
}
```

**Reason Values:**
- `null`: Can access
- `premium_required`: Requires subscription
- `not_authenticated`: Requires login

---

### 10. Genres

#### 10.1 List Genres
**GET** `/genres`

List all active genres.

**Response (200):**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "name": "Drama",
      "slug": "drama",
      "description": "Drama series",
      "display_order": 1
    }
  ]
}
```

#### 10.2 Get Genre Detail
**GET** `/genres/{id}` or `/genres/{slug}`

Get genre details with series.

**Path Parameters:**
- `id`: Genre ID or slug

**Query Parameters:**
- `page`: Page number (default: 1)
- `limit`: Items per page (default: 20)

**Response (200):**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "Drama",
    "slug": "drama",
    "description": "Drama series",
    "series": {
      "items": [
        {
          "id": 1,
          "title": "Drama Series Title",
          "poster_url": "https://example.com/poster.jpg",
          "year": 2024,
          "rating": 8.5
        }
      ],
      "pagination": {
        "page": 1,
        "limit": 20,
        "total": 50
      }
    }
  }
}
```

---

### 11. Admin CMS APIs

#### 11.1 Create Series
**POST** `/admin/series`
**Authentication Required** (Admin Role)

Create a new series.

**Request Body:**
```json
{
  "title": "New Series",
  "slug": "new-series",
  "description": "Series description",
  "poster_url": "https://example.com/poster.jpg",
  "backdrop_url": "https://example.com/backdrop.jpg",
  "year": 2024,
  "is_premium": false,
  "genre_ids": [1, 2],
  "tag_ids": [1, 2],
  "released_at": "2024-01-01T00:00:00Z"
}
```

**Response (201):**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "uuid": "uuid",
    "title": "New Series",
    "slug": "new-series",
    "status": "draft",
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

#### 11.2 Update Series
**PUT** `/admin/series/{id}`
**Authentication Required** (Admin Role)

Update a series.

**Request Body:**
```json
{
  "title": "Updated Title",
  "description": "Updated description",
  "status": "published"
}
```

**Response (200):**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "title": "Updated Title",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

#### 11.3 Delete Series
**DELETE** `/admin/series/{id}`
**Authentication Required** (Admin Role)

Soft delete a series.

**Response (200):**
```json
{
  "success": true,
  "message": "Series deleted successfully"
}
```

#### 11.4 Create Season
**POST** `/admin/series/{series_id}/seasons`
**Authentication Required** (Admin Role)

Create a new season for a series.

**Request Body:**
```json
{
  "season_number": 1,
  "title": "Season 1",
  "description": "Season description",
  "poster_url": "https://example.com/season1.jpg",
  "released_at": "2024-01-01T00:00:00Z"
}
```

**Response (201):**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "uuid": "uuid",
    "series_id": 1,
    "season_number": 1,
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

#### 11.5 Create Episode
**POST** `/admin/seasons/{season_id}/episodes`
**Authentication Required** (Admin Role)

Create a new episode for a season.

**Request Body:**
```json
{
  "episode_number": 1,
  "title": "Episode 1",
  "description": "Episode description",
  "thumbnail_url": "https://example.com/ep1.jpg",
  "video_url": "https://s3.example.com/videos/ep1.mp4",
  "duration": 120,
  "is_premium": false,
  "released_at": "2024-01-01T00:00:00Z"
}
```

**Response (201):**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "uuid": "uuid",
    "series_id": 1,
    "season_id": 1,
    "episode_number": 1,
    "status": "draft",
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

#### 11.6 Update Episode
**PUT** `/admin/episodes/{id}`
**Authentication Required** (Admin Role)

Update an episode.

**Request Body:**
```json
{
  "title": "Updated Episode Title",
  "status": "published",
  "is_premium": true
}
```

**Response (200):**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "title": "Updated Episode Title",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

#### 11.7 Delete Episode
**DELETE** `/admin/episodes/{id}`
**Authentication Required** (Admin Role)

Soft delete an episode.

**Response (200):**
```json
{
  "success": true,
  "message": "Episode deleted successfully"
}
```

#### 11.8 Create Genre
**POST** `/admin/genres`
**Authentication Required** (Admin Role)

Create a new genre.

**Request Body:**
```json
{
  "name": "Action",
  "slug": "action",
  "description": "Action series",
  "display_order": 1
}
```

**Response (201):**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "Action",
    "slug": "action",
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

#### 11.9 Create Tag
**POST** `/admin/tags`
**Authentication Required** (Admin Role)

Create a new tag.

**Request Body:**
```json
{
  "name": "thriller",
  "slug": "thriller"
}
```

**Response (201):**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "thriller",
    "slug": "thriller",
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

---

### 12. Analytics Events

#### 12.1 Ingest Analytics Event
**POST** `/analytics/events`
**Authentication Required**

Submit an analytics event.

**Request Body:**
```json
{
  "event_type": "video_start",
  "event_data": {
    "episode_id": 1,
    "series_id": 1,
    "progress": 0,
    "duration": 120
  }
}
```

**Response (202):**
```json
{
  "success": true,
  "message": "Event accepted"
}
```

**Event Types:**
- `page_view`: User viewed a page
- `video_start`: User started watching a video
- `video_complete`: User completed watching a video
- `video_pause`: User paused a video
- `search`: User performed a search
- `add_to_list`: User added series to my list
- `remove_from_list`: User removed series from my list
- `subscription`: User subscribed

---

### 13. Health Check

#### 13.1 Health Check
**GET** `/health`

Basic health check.

**Response (200):**
```json
{
  "success": true,
  "status": "healthy",
  "timestamp": "2024-01-01T00:00:00Z"
}
```

#### 13.2 Readiness Check
**GET** `/health/ready`

Readiness check with dependency status.

**Response (200):**
```json
{
  "success": true,
  "status": "ready",
  "dependencies": {
    "database": "healthy",
    "redis": "healthy",
    "storage": "healthy"
  },
  "timestamp": "2024-01-01T00:00:00Z"
}
```

---

## Rate Limiting

- **Authentication endpoints**: 5 requests per minute per IP
- **General API**: 100 requests per minute per user
- **Admin endpoints**: 50 requests per minute per admin

Rate limit headers are included in responses:
```
X-RateLimit-Limit: 100
X-RateLimit-Remaining: 95
X-RateLimit-Reset: 1609459200
```

---

## CORS

Allowed origins for development:
- `http://localhost:3000`
- `http://localhost:8080`

Production origins will be configured via environment variables.

---

## Versioning

API version is specified in the URL path: `/api/v1/`

Future versions will follow semantic versioning:
- `/api/v2/` for breaking changes
- `/api/v1/` will remain supported for a deprecation period

---

## Webhooks (Future)

Webhook support for:
- Subscription events (created, renewed, expired)
- Content publishing events
- User registration events

This will be implemented in a future phase.
