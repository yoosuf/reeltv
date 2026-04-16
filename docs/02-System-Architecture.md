# Reel TV System Architecture

## Architecture Overview

Reel TV follows a clean modular monolith architecture designed for scalability and future microservices extraction. The system is API-first with clear separation of concerns across layers.

## High-Level Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                         Client Layer                        │
│  (Mobile App - Flutter | Admin Web UI - Future Phase 2)    │
└───────────────────────────┬─────────────────────────────────┘
                            │ HTTPS/REST
┌───────────────────────────▼─────────────────────────────────┐
│                      API Gateway Layer                      │
│  (Gin Router | Middleware | Rate Limiting | CORS)          │
└───────────────────────────┬─────────────────────────────────┘
                            │
┌───────────────────────────▼─────────────────────────────────┐
│                    Handler/Controller Layer                 │
│  (Request Validation | Response Formatting | Auth Check)     │
└───────────────────────────┬─────────────────────────────────┘
                            │
┌───────────────────────────▼─────────────────────────────────┐
│                       Service Layer                         │
│  (Business Logic | Domain Rules | Orchestration)            │
│  ┌─────────┐ ┌──────────┐ ┌────────────┐ ┌──────────────┐  │
│  │  Auth   │ │ Catalog  │ │  Playback  │ │ Recommendation│  │
│  │ Service │ │ Service  │ │  Service   │ │   Service    │  │
│  └─────────┘ └──────────┘ └────────────┘ └──────────────┘  │
└───────────────────────────┬─────────────────────────────────┘
                            │
┌───────────────────────────▼─────────────────────────────────┐
│                    Repository Layer                         │
│  (Data Access | GORM Operations | Caching Strategy)          │
└───────────────────────────┬─────────────────────────────────┘
                            │
        ┌───────────────────┼───────────────────┐
        │                   │                   │
┌───────▼────────┐  ┌──────▼─────────┐  ┌─────▼──────────┐
│  PostgreSQL    │  │     Redis      │  │  S3 Storage    │
│  (Primary DB)  │  │   (Cache)      │  │  (Video/Media) │
└────────────────┘  └────────────────┘  └────────────────┘
```

## Component Layers

### 1. API Gateway Layer
- **Framework**: Gin (HTTP router)
- **Responsibilities**:
  - Request routing
  - Middleware chain execution
  - CORS handling
  - Request ID generation
  - Panic recovery
  - Rate limiting (Redis-based)
  - Request logging

### 2. Handler/Controller Layer
- **Responsibilities**:
  - HTTP request/response handling
  - Request validation (struct tags)
  - JWT token extraction and validation
  - Response formatting (consistent JSON structure)
  - Error translation to HTTP status codes
  - Input sanitization

### 3. Service Layer
- **Responsibilities**:
  - Business logic implementation
  - Domain rule enforcement
  - Transaction coordination
  - External service integration (future)
  - Cache orchestration
  - Event publishing (analytics)

**Service Modules**:
- **AuthService**: User registration, login, token generation, password reset
- **CatalogService**: Series, seasons, episodes, genres, tags management
- **PlaybackService**: Watch progress tracking, continue watching
- **UserService**: Profile management, my list operations
- **RecommendationService**: Heuristic-based content recommendations
- **SubscriptionService**: Entitlement checks, subscription validation
- **AdminService**: CMS operations, content management
- **AnalyticsService**: Event ingestion, basic metrics

### 4. Repository Layer
- **Responsibilities**:
  - Database operations via GORM
  - Cache read/write operations
  - Query optimization
  - Data mapping between domain models and DB entities
  - Transaction management

### 5. Model Layer
- **Responsibilities**:
  - Domain entity definitions
  - GORM model definitions
  - Validation rules
  - Database relationships

## Data Storage Strategy

### PostgreSQL (Primary Database)
- **Purpose**: Persistent data storage
- **Data**:
  - Users and authentication
  - Catalog metadata (series, seasons, episodes)
  - Watch progress
  - My list
  - Subscriptions
  - Genres and tags
- **Connection Pooling**: Configured via GORM
- **Migrations**: Auto-migration with GORM + versioned migration files

### Redis (Cache and Session Store)
- **Purpose**:
  - Session storage (if using session-based auth)
  - Rate limiting counters
  - Catalog caching (frequently accessed series/episodes)
  - JWT token blacklist (logout)
  - Temporary data (verification codes, future)
- **TTL Strategy**:
  - Catalog: 5-15 minutes
  - Rate limits: 1 minute to 1 hour
  - Sessions: 24 hours

### S3-Compatible Storage (Media)
- **Purpose**: Video files, thumbnails, posters
- **Implementation**:
  - Local dev: MinIO (Docker)
  - Production: AWS S3 or compatible
- **Organization**:
  - `/videos/series/{series_id}/seasons/{season_id}/episodes/{episode_id}/`
  - `/images/posters/{series_id}/`
  - `/images/thumbnails/{episode_id}/`

## Authentication & Authorization

### Authentication Flow
```
1. User Signup/Login Request
   ↓
2. Handler validates credentials
   ↓
3. Service hashes password (bcrypt)
   ↓
4. Repository stores/retrieves user
   ↓
5. Service generates JWT (access + refresh)
   ↓
6. Handler returns tokens
```

### JWT Token Structure
- **Access Token**: 15 minutes expiration, contains user_id, role
- **Refresh Token**: 7 days expiration, stored in DB or Redis
- **Claims**:
  - `sub`: user_id
  - `role`: user | admin
  - `iat`: issued at
  - `exp`: expiration

### Authorization
- **Role-based**: User vs Admin
- **Resource-based**: Users can only access their own data (my list, watch progress)
- **Middleware**: JWT validation on protected routes
- **Admin Routes**: Require admin role

## Caching Strategy

### Cache-Aside Pattern
```
1. Service checks Redis cache
2. If cache hit → return data
3. If cache miss → query PostgreSQL
4. Store result in Redis
5. Return data
```

### Cache Keys
- `catalog:series:{id}` → Series details
- `catalog:episode:{id}` → Episode details
- `catalog:trending` → Trending series list
- `catalog:genre:{id}` → Series by genre
- `user:watch_progress:{user_id}` → User's watch progress
- `rate_limit:{user_id}:{endpoint}` → Rate limit counter

### Cache Invalidation
- **Write-through**: Invalidate cache on content updates (admin operations)
- **TTL-based**: Automatic expiration
- **Manual**: Cache flush endpoint for admin

## Video Pipeline (Architecture Ready)

### Video Upload Flow
```
1. Admin uploads video file via CMS
   ↓
2. Handler receives file
   ↓
3. Service uploads to temporary S3 location
   ↓
4. FFmpeg worker processes video:
   - Transcode to vertical format (9:16)
   - Generate multiple qualities (480p, 720p, 1080p)
   - Generate thumbnail
   ↓
5. Processed files stored in final S3 location
   ↓
6. Metadata updated in database
```

### Video Playback Flow
```
1. Client requests episode
   ↓
2. Service returns video URL (presigned S3 URL or CDN URL)
   ↓
3. Client streams video directly from S3/CDN
   ↓
4. Client sends watch progress updates
   ↓
5. Service updates progress in database
```

## API Design Principles

### RESTful Conventions
- **GET**: Retrieve data (idempotent)
- **POST**: Create resources
- **PUT/PATCH**: Update resources
- **DELETE**: Remove resources (idempotent)

### Response Format
```json
{
  "success": true,
  "data": { ... },
  "message": "Operation successful",
  "request_id": "uuid"
}
```

### Error Response Format
```json
{
  "success": false,
  "error": {
    "code": "ERROR_CODE",
    "message": "Human-readable error message",
    "details": { ... }
  },
  "request_id": "uuid"
}
```

### Pagination
- Cursor-based for large datasets (watch history)
- Offset-based for catalog browsing
- Default page size: 20
- Max page size: 100

## Scalability Considerations

### Horizontal Scaling
- **Stateless API servers**: No session state in memory
- **Load balancer ready**: Multiple API instances behind LB
- **Database**: Read replicas for catalog queries (future)
- **Cache**: Redis Cluster for distributed caching (future)

### Performance Optimization
- **Database indexing**: Strategic indexes on foreign keys, timestamps
- **Query optimization**: N+1 prevention via GORM Preload
- **Connection pooling**: GORM connection pool configuration
- **Compression**: Gzip middleware for API responses

### Future Microservices Extraction
The modular monolith design allows extraction of:
- **Catalog Service**: Content metadata and search
- **Playback Service**: Watch progress and recommendations
- **User Service**: Authentication and profiles
- **Analytics Service**: Event processing and metrics

## Security Measures

### Input Validation
- Struct tag validation on all inputs
- SQL injection prevention via GORM parameterized queries
- XSS prevention via input sanitization

### Rate Limiting
- Auth endpoints: 5 requests per minute per IP
- General API: 100 requests per minute per user
- Admin endpoints: Stricter limits

### Data Protection
- Passwords: bcrypt with cost factor 10
- PII: Encrypted at rest (future enhancement)
- API keys: Environment variables, never in code

### Network Security
- HTTPS only in production
- CORS configured for specific origins
- Security headers (CSP, X-Frame-Options)

## Observability

### Logging
- **Format**: Structured JSON logs
- **Levels**: DEBUG, INFO, WARN, ERROR
- **Context**: Request ID, user ID, timestamp
- **Destinations**: STDOUT (Docker logs), file (optional)

### Metrics (Basic for MVP)
- Request counts by endpoint
- Error rates
- Response times (p50, p95, p99)
- Database query times
- Cache hit/miss ratios

### Health Checks
- `/health`: Basic liveness
- `/health/ready`: Readiness (DB, Redis, S3 connectivity)
- `/metrics`: Prometheus metrics endpoint (future)

## Deployment Architecture (Local Dev)

### Docker Compose Services
```
reeltv-api          (Go application)
reeltv-db           (PostgreSQL 15)
reeltv-redis        (Redis 7)
reeltv-minio        (MinIO S3-compatible)
reeltv-ffmpeg-worker (FFmpeg worker - optional for MVP)
```

### Service Communication
- All services on shared Docker network
- API connects to DB via internal DNS
- Environment variables for configuration
- Volume mounts for persistence (DB, Redis, MinIO)

## Technology Stack Rationale

### Go
- **Performance**: Excellent for high-throughput APIs
- **Concurrency**: Goroutines for parallel processing
- **Type Safety**: Compile-time error detection
- **Deployment**: Single binary, easy containerization
- **Ecosystem**: Rich libraries for web, database, caching

### GORM
- **Productivity**: ORM reduces boilerplate
- **Migrations**: Built-in migration support
- **Relationships**: Easy handling of complex relationships
- **Community**: Well-maintained, widely used

### PostgreSQL
- **Reliability**: ACID compliance, mature
- **Features**: JSONB for flexible metadata, full-text search
- **Performance**: Excellent for read-heavy workloads
- **Scalability**: Replication, partitioning support

### Redis
- **Speed**: In-memory, sub-millisecond operations
- **Features**: Rich data structures, pub/sub
- **Use Cases**: Caching, rate limiting, sessions
- **Durability**: Optional persistence (AOF/RDB)

### Gin
- **Performance**: Fast HTTP router
- **Middleware**: Extensive middleware ecosystem
- **JSON**: Built-in JSON binding/validation
- **Developer Experience**: Clean API, good documentation

## Migration to Microservices (Future)

### Extraction Triggers
- Team size > 10 developers
- Single service deployment time > 10 minutes
- Resource contention between modules
- Different scaling requirements per module

### Extraction Strategy
1. Identify bounded contexts
2. Extract service with API gateway
3. Implement service-to-service communication (gRPC/HTTP)
4. Data migration and synchronization
5. Gradual traffic shifting
6. Decommission old code

## Phase 1 Scope

Phase 1 implements the complete backend foundation:
- All layers (Handler → Service → Repository → Model)
- All MVP modules (Auth, Catalog, Playback, User, Recommendation, Subscription, Admin, Analytics)
- Docker Compose setup
- Database migrations
- Seed data
- API tests

No frontend development until Phase 1 backend is complete and tested.
