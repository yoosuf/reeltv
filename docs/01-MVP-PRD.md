# Reel TV MVP Product Requirements Document

## Executive Summary

Reel TV is a curated vertical short-drama streaming platform. Unlike UGC platforms like TikTok, Reel TV features professionally produced, short-form dramatic content organized as series with seasons and episodes. Episodes are vertically formatted, typically 1-2 minutes each, optimized for mobile viewing.

## Product Vision

To become the premier destination for binge-worthy short-form drama content, offering users an addictive, mobile-first streaming experience with seamless episode progression and personalized recommendations.

## Target Audience

- Mobile-first users aged 18-45
- Users who enjoy dramatic storytelling but prefer shorter content formats
- Commuters, casual viewers seeking quick entertainment
- Users who value curated, high-quality content over user-generated content

## Core Value Propositions

1. **Curated Quality**: Professionally produced content, not UGC
2. **Binge-able Format**: 1-2 minute episodes perfect for quick viewing sessions
3. **Vertical Optimization**: Content designed for mobile screens
4. **Personalized Discovery**: Smart recommendations based on viewing behavior
5. **Flexible Access**: Freemium model with premium content tiers

## MVP Scope

### In Scope for MVP

#### User Features
- **Authentication**: Email/phone-based signup and login
- **Content Discovery**: Browse series by genre, trending, and new releases
- **Series Detail**: View series information, seasons, and episode lists
- **Episode Playback**: Watch episodes with vertical video player
- **Continue Watching**: Resume playback from where user left off
- **My List**: Save favorite series for later access
- **Search**: Search series by title, genre, or tags
- **Recommendations**: Personalized content suggestions
- **Premium Access**: Subscription-based premium content unlocking

#### Content Model
- Series with metadata (title, description, poster, genre, tags)
- Seasons organizing episodes
- Episodes with metadata (title, description, duration, thumbnail, video URL)
- Free vs Premium episode classification
- Genre and tag taxonomy
- Localized metadata support (future-ready)

#### Admin Features
- CMS for managing series, seasons, and episodes
- Content metadata management
- Premium/free tier configuration
- Basic analytics event ingestion

### Out of Scope for MVP (Phase 1)

- Social features (comments, shares, likes)
- User-generated content
- Live streaming
- Multi-language audio/subtitles (infrastructure ready, content not)
- Advanced analytics dashboards
- Payment processing (subscription check only, actual payment gateway integration deferred)
- Push notifications
- Offline viewing
- Multiple user profiles per account
- Content ratings and reviews
- Advanced recommendation algorithms (ML-based)

## User Stories

### Authentication
- As a new user, I can sign up with email or phone number
- As a returning user, I can login with my credentials
- As a user, I can logout and login again on different devices

### Content Discovery
- As a user, I can see a home feed with trending series
- As a user, I can browse series by genre
- As a user, I can see newly released episodes
- As a user, I can search for series by title

### Series & Episode Viewing
- As a user, I can view detailed information about a series
- As a user, I can see all seasons and episodes in a series
- As a user, I can watch an episode in vertical format
- As a user, I can autoplay the next episode in a series
- As a user, I can see which episodes are free vs premium

### Personalization
- As a user, I can see my continue watching list to resume playback
- As a user, I can add series to my list for later
- As a user, I can remove series from my list
- As a user, I can see personalized recommendations

### Premium Access
- As a free user, I can watch free episodes
- As a premium user, I can watch all episodes including premium content
- As a user, I can see clear indicators for premium content
- As a user, I can see unlock prompts for premium episodes

### Admin CMS
- As an admin, I can create and manage series
- As an admin, I can create and manage seasons
- As an admin, I can create and manage episodes
- As an admin, I can manage genres and tags
- As an admin, I can configure premium/free access for episodes

## Business Rules

### Content Access
- Episodes can be marked as free or premium
- Free episodes are accessible to all users
- Premium episodes require active subscription
- Users must be authenticated to access any content

### Watch Progress
- Watch progress is tracked per user per episode
- Progress is saved when user exits or periodically during playback
- Continue watching shows episodes with >0% and <90% progress
- Episodes with ≥90% progress are considered completed

### Recommendations (Heuristic-based for MVP)
- Prioritize recently watched series with new episodes
- Recommend series in same genres as user's completed content
- Boost popular/trending series
- Consider rewatch behavior for completed series
- Weight recency of user activity

### My List
- Users can add any series to their list regardless of premium status
- My list is personal to each user
- No limit on number of series in list for MVP

### Subscription
- Subscription status is checked via API
- Subscription can be active, inactive, or expired
- Premium access is granted only with active subscription
- Grace period handling for expired subscriptions (future)

## Non-Functional Requirements

### Performance
- API response time < 200ms for catalog endpoints
- API response time < 500ms for search
- Video playback start time < 2 seconds
- Support 1000 concurrent users for MVP

### Scalability
- Architecture designed for horizontal scaling
- Stateless API servers
- Database connection pooling
- Redis caching for frequently accessed data

### Security
- JWT-based authentication with reasonable expiration
- Password hashing (bcrypt)
- HTTPS for all API endpoints
- Input validation and sanitization
- Rate limiting on auth endpoints
- SQL injection prevention via parameterized queries (GORM)

### Reliability
- 99.5% uptime target for MVP
- Graceful degradation for non-critical features
- Database connection retry logic
- Health check endpoints for monitoring

### Observability
- Structured logging (JSON format)
- Request ID tracking
- Error tracking hooks
- Basic metrics collection (request counts, error rates)
- Analytics event ingestion for user behavior

## Technical Constraints

- Backend: Go 1.21+
- Database: PostgreSQL 15+
- Cache: Redis 7+
- Storage: S3-compatible object storage (MinIO for local dev)
- Video processing: FFmpeg
- Deployment: Docker + Docker Compose for local dev
- API: REST with JSON
- Authentication: JWT tokens

## Success Metrics (MVP)

- User registration rate
- Daily active users (DAU)
- Average session duration
- Episode completion rate
- Continue watching usage
- My list adoption rate
- Premium conversion rate (when payment gateway added)

## Assumptions

1. Video content will be pre-encoded and uploaded to S3-compatible storage
2. Initial content will be seed data for testing purposes
3. Payment gateway integration will be handled separately (subscription check only for MVP)
4. Email/phone verification may be deferred to post-MVP
5. Content moderation will be manual via admin CMS for MVP
6. Video CDN will be the S3-compatible storage directly for MVP (dedicated CDN post-MVP)
7. Recommendation algorithm will be heuristic-based, not ML-based for MVP
8. Single region deployment for MVP

## Risks and Mitigations

| Risk | Impact | Mitigation |
|------|--------|------------|
| Content acquisition challenges | High | Start with licensed or partner content, clear content strategy |
| Video streaming performance | High | Use reliable S3-compatible storage, optimize video encoding |
| User acquisition cost | Medium | Focus on organic growth, social sharing features post-MVP |
| Subscription conversion | Medium | Clear value proposition, free trial strategy post-MVP |
| Technical scalability | Medium | Architecture designed for scaling from day one |

## Phase 1 Focus

Phase 1 focuses exclusively on the backend foundation:
- Complete API implementation
- Database schema and migrations
- Authentication and authorization
- All core business logic
- Admin CMS APIs
- Docker-based local development environment
- Comprehensive API testing

No frontend development will occur until Phase 1 backend is complete and tested.
