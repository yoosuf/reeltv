# API Testing Guide

This guide provides curl examples for testing the Reel TV MVP API endpoints.

## Base URL

- Development: `http://localhost:8080/api/v1`

## Health Check

```bash
curl http://localhost:8080/health
```

Expected response:
```json
{
  "status": "healthy"
}
```

## Authentication

### Register

```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "phone": "+1234567890",
    "password": "Test123!",
    "name": "Test User"
  }'
```

### Login

```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "Test123!"
  }'
```

Expected response:
```json
{
  "success": true,
  "data": {
    "access_token": "jwt_access_token_here",
    "refresh_token": "jwt_refresh_token_here",
    "user": {
      "id": 1,
      "uuid": "user-uuid-here",
      "email": "test@example.com",
      "name": "Test User"
    }
  }
}
```

### Refresh Token

```bash
curl -X POST http://localhost:8080/api/v1/auth/refresh \
  -H "Content-Type: application/json" \
  -d '{
    "refresh_token": "your_refresh_token_here"
  }'
```

### Logout

```bash
curl -X POST http://localhost:8080/api/v1/auth/logout \
  -H "Content-Type: application/json" \
  -d '{
    "refresh_token": "your_refresh_token_here"
  }'
```

## User Profile (requires authentication)

Set your access token as an environment variable:
```bash
export TOKEN="your_access_token_here"
```

### Get Profile

```bash
curl http://localhost:8080/api/v1/users/me \
  -H "Authorization: Bearer $TOKEN"
```

### Update Profile

```bash
curl -X PUT http://localhost:8080/api/v1/users/me \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Updated Name",
    "avatar_url": "https://example.com/avatar.jpg"
  }'
```

### Change Password

```bash
curl -X POST http://localhost:8080/api/v1/users/me/change-password \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "current_password": "Test123!",
    "new_password": "NewTest123!"
  }'
```

## Catalog

### List Series

```bash
curl "http://localhost:8080/api/v1/catalog/series?offset=0&limit=20"
```

### Get Series by ID

```bash
curl http://localhost:8080/api/v1/catalog/series/1
```

### Get Series by Slug

```bash
curl http://localhost:8080/api/v1/catalog/series/slug/breaking-bad
```

### Get Seasons for a Series

```bash
curl http://localhost:8080/api/v1/catalog/series/1/seasons
```

### Get Episodes for a Season

```bash
curl "http://localhost:8080/api/v1/catalog/seasons/1/episodes?offset=0&limit=20"
```

### Create Series (admin only - requires authentication)

```bash
curl -X POST http://localhost:8080/api/v1/catalog/series \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "New Series",
    "slug": "new-series",
    "description": "Series description",
    "year": 2026,
    "genre": "Drama",
    "language": "en"
  }'
```

### Update Series (admin only - requires authentication)

```bash
curl -X PUT http://localhost:8080/api/v1/catalog/series/1 \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Updated Series Title",
    "status": "active"
  }'
```

### Delete Series (admin only - requires authentication)

```bash
curl -X DELETE http://localhost:8080/api/v1/catalog/series/1 \
  -H "Authorization: Bearer $TOKEN"
```

## Watch Progress (requires authentication)

### Update Watch Progress

```bash
curl -X POST http://localhost:8080/api/v1/watch-progress \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "episode_id": 1,
    "duration": 120,
    "percentage": 50.0,
    "completed": false
  }'
```

### Get Watch Progress

```bash
curl "http://localhost:8080/api/v1/watch-progress?offset=0&limit=20" \
  -H "Authorization: Bearer $TOKEN"
```

## My List (requires authentication)

### Add to My List

```bash
curl -X POST http://localhost:8080/api/v1/my-list \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "series_id": 1
  }'
```

### Get My List

```bash
curl "http://localhost:8080/api/v1/my-list?offset=0&limit=20" \
  -H "Authorization: Bearer $TOKEN"
```

### Remove from My List

```bash
curl -X DELETE http://localhost:8080/api/v1/my-list/1 \
  -H "Authorization: Bearer $TOKEN"
```

## Error Responses

All endpoints return consistent error responses:

```json
{
  "success": false,
  "error": {
    "code": "ERROR_CODE",
    "message": "Human readable error message"
  }
}
```

Common error codes:
- `UNAUTHORIZED` - Authentication required or invalid
- `INVALID_REQUEST` - Request validation failed
- `NOT_FOUND` - Resource not found
- `ALREADY_EXISTS` - Resource already exists
- `INTERNAL_ERROR` - Server error

## Testing with Docker

1. Start services:
```bash
make docker-up
```

2. Run migrations:
```bash
make migrate-up
```

3. Seed database:
```bash
make seed
```

4. Run the application:
```bash
make run
```

5. Test endpoints using the curl examples above.
