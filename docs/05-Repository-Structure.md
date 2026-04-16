# Reel TV Repository Structure (Modular DDD Approach)

```
reeltv/
├── cmd/
│   └── api/
│       └── main.go                 # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go              # Configuration loading and management
│   ├── shared/
│   │   ├── domain/
│   │   │   ├── base.go            # Base entity with ID, timestamps
│   │   │   └── errors.go          # Domain errors
│   │   ├── infrastructure/
│   │   │   ├── persistence/
│   │   │   │   ├── db.go         # Database connection setup
│   │   │   │   └── redis.go       # Redis client setup
│   │   │   └── messaging/         # Event bus (future)
│   │   └── application/
│   │       └── dto/               # Shared DTOs
│   ├── user/
│   │   ├── domain/
│   │   │   ├── entity.go          # User entity, value objects
│   │   │   ├── repository.go      # Repository interface
│   │   │   └── service.go         # Domain service (business rules)
│   │   ├── application/
│   │   │   ├── service.go         # Application service (use cases)
│   │   │   └── dto.go             # DTOs for this module
│   │   ├── infrastructure/
│   │   │   ├── persistence/
│   │   │   │   └── repository.go  # GORM repository implementation
│   │   │   └── messaging/         # Event publishers (future)
│   │   └── interface/
│   │       └── http/
│   │           ├── handler.go     # HTTP handlers
│   │           └── dto.go         # Request/Response DTOs
│   ├── auth/
│   │   ├── domain/
│   │   │   ├── entity.go          # RefreshToken entity
│   │   │   ├── repository.go      # Repository interface
│   │   │   └── service.go         # Auth domain service
│   │   ├── application/
│   │   │   ├── service.go         # Auth use cases (login, register, refresh)
│   │   │   └── dto.go
│   │   ├── infrastructure/
│   │   │   ├── persistence/
│   │   │   │   └── repository.go
│   │   │   └── jwt/
│   │   │       └── provider.go    # JWT token generation/validation
│   │   └── interface/
│   │       └── http/
│   │           └── handler.go
│   ├── catalog/
│   │   ├── domain/
│   │   │   ├── entity.go          # Series, Season, Episode, Genre, Tag entities
│   │   │   ├── repository.go      # Repository interfaces
│   │   │   └── service.go         # Catalog domain service
│   │   ├── application/
│   │   │   ├── service.go         # Catalog use cases
│   │   │   └── dto.go
│   │   ├── infrastructure/
│   │   │   ├── persistence/
│   │   │   │   └── repository.go # GORM implementations
│   │   │   └── cache/
│   │   │       └── provider.go    # Redis caching
│   │   └── interface/
│   │       └── http/
│   │           ├── handler.go
│   │           └── dto.go
│   ├── playback/
│   │   ├── domain/
│   │   │   ├── entity.go          # WatchProgress entity
│   │   │   ├── repository.go
│   │   │   └── service.go
│   │   ├── application/
│   │   │   ├── service.go         # Watch progress use cases
│   │   │   └── dto.go
│   │   ├── infrastructure/
│   │   │   └── persistence/
│   │   │       └── repository.go
│   │   └── interface/
│   │       └── http/
│   │           └── handler.go
│   ├── mylist/
│   │   ├── domain/
│   │   │   ├── entity.go          # MyList entity
│   │   │   ├── repository.go
│   │   │   └── service.go
│   │   ├── application/
│   │   │   ├── service.go
│   │   │   └── dto.go
│   │   ├── infrastructure/
│   │   │   └── persistence/
│   │   │       └── repository.go
│   │   └── interface/
│   │       └── http/
│   │           └── handler.go
│   ├── recommendation/
│   │   ├── domain/
│   │   │   ├── repository.go      # Repository interfaces for data access
│   │   │   └── service.go         # Recommendation algorithm
│   │   ├── application/
│   │   │   ├── service.go         # Recommendation use cases
│   │   │   └── dto.go
│   │   ├── infrastructure/
│   │   │   └── persistence/
│   │   │       └── repository.go
│   │   └── interface/
│   │       └── http/
│   │           └── handler.go
│   ├── subscription/
│   │   ├── domain/
│   │   │   ├── entity.go          # Subscription entity
│   │   │   ├── repository.go
│   │   │   └── service.go         # Entitlement logic
│   │   ├── application/
│   │   │   ├── service.go
│   │   │   └── dto.go
│   │   ├── infrastructure/
│   │   │   ├── persistence/
│   │   │   │   └── repository.go
│   │   │   └── payment/           # Payment gateway (future)
│   │   │       └── provider.go
│   │   └── interface/
│   │       └── http/
│   │           └── handler.go
│   ├── analytics/
│   │   ├── domain/
│   │   │   ├── entity.go          # AnalyticsEvent entity
│   │   │   ├── repository.go
│   │   │   └── service.go
│   │   ├── application/
│   │   │   ├── service.go
│   │   │   └── dto.go
│   │   ├── infrastructure/
│   │   │   ├── persistence/
│   │   │   │   └── repository.go
│   │   │   └── streaming/         # Event streaming (future)
│   │   │       └── producer.go
│   │   └── interface/
│   │       └── http/
│   │           └── handler.go
│   ├── admin/
│   │   ├── application/
│   │   │   ├── service.go         # Admin use cases (orchestrates other modules)
│   │   │   └── dto.go
│   │   └── interface/
│   │       └── http/
│   │           └── handler.go     # Admin HTTP handlers
│   └── interface/
│       └── http/
│           ├── router.go          # Route registration
│           ├── middleware/
│           │   ├── auth.go
│           │   ├── cors.go
│           │   ├── rate_limit.go
│           │   ├── logging.go
│           │   ├── request_id.go
│           │   └── error_handler.go
│           └── server.go          # HTTP server setup
├── pkg/
│   ├── logger/
│   │   └── logger.go              # Structured logging
│   ├── password/
│   │   └── password.go            # Password hashing (bcrypt)
│   ├── validator/
│   │   └── validator.go           # Request validation
│   ├── utils/
│   │   ├── uuid.go                # UUID generation
│   │   ├── slug.go                # Slug generation
│   │   └── time.go                # Time utilities
│   └── storage/
│       └── storage.go            # S3-compatible storage interface
├── migrations/
│   ├── 000001_init_schema.up.sql
│   └── 000001_init_schema.down.sql
├── configs/
│   ├── config.yaml                 # Configuration template
│   └── config.example.yaml         # Example configuration
├── scripts/
│   ├── migrate.sh                  # Database migration script
│   ├── seed.sh                     # Seed data script
│   └── test.sh                     # API test script
├── test/
│   ├── integration/
│   │   ├── auth_test.go
│   │   ├── catalog_test.go
│   │   └── playback_test.go
│   └── fixtures/
│       └── seed_data.go           # Test fixtures
├── deployments/
│   ├── Dockerfile                  # Docker image for API
│   ├── docker-compose.yml          # Docker Compose for local dev
│   └── docker-compose.prod.yml    # Docker Compose for production
├── docs/
│   ├── 01-MVP-PRD.md
│   ├── 02-System-Architecture.md
│   ├── 03-Domain-Model-Database-Schema.md
│   ├── 04-API-Specification.md
│   ├── 05-Repository-Structure.md
│   ├── 06-Implementation-Plan.md
│   └── API-Examples.md
├── .gitignore
├── .dockerignore
├── .env.example                    # Environment variables template
├── go.mod                          # Go module definition
├── go.sum                          # Go dependency checksums
├── Makefile                        # Build and development commands
└── README.md                       # Project documentation
```

## Directory Descriptions

### `/cmd/api`
Application entry point. Contains the `main.go` file that initializes the application and starts the HTTP server.

### `/internal`
Private application code organized by bounded contexts (DDD). Each module represents a domain context with its own layers.

#### `/internal/shared`
Cross-cutting concerns shared across all bounded contexts.

- **domain**: Base entities, common domain errors
- **infrastructure**: Shared infrastructure (database connection, Redis client, event bus)
- **application**: Shared DTOs and application-level utilities

#### Bounded Context Modules
Each bounded context (user, auth, catalog, playback, mylist, recommendation, subscription, analytics, admin) follows the same DDD layered structure:

**domain/**: Core business logic
- `entity.go`: Domain entities with business rules (rich domain model)
- `repository.go`: Repository interfaces (contracts for data access)
- `service.go`: Domain services (complex business logic that doesn't fit in entities)

**application/**: Application use cases
- `service.go`: Application services (orchestrate domain operations, transaction boundaries)
- `dto.go`: Data Transfer Objects for this module's use cases

**infrastructure/**: Technical implementations
- `persistence/repository.go`: GORM implementations of repository interfaces
- `cache/provider.go`: Redis caching implementations
- `jwt/provider.go`: JWT token generation/validation (auth module)
- `payment/provider.go`: Payment gateway integration (subscription module)
- `messaging/**`: Event publishers/subscribers (future)

**interface/http/**: External interfaces
- `handler.go`: HTTP handlers/controllers (thin, delegate to application services)
- `dto.go`: HTTP request/response DTOs

## DDD Principles Applied

- **Bounded Contexts**: Each module (user, auth, catalog, etc.) is a bounded context with clear boundaries
- **Ubiquitous Language**: Each module uses language specific to its domain
- **Domain Entities**: Rich domain models with business logic, not just data holders
- **Repository Pattern**: Interfaces in domain layer, implementations in infrastructure
- **Dependency Inversion**: Domain layer doesn't depend on infrastructure
- **Application Services**: Orchestrate use cases, manage transactions
- **Thin Controllers**: HTTP handlers are thin, delegate to application services
- **Separation of Concerns**: Clear separation between domain, application, infrastructure, and interface layers

## Security Considerations

### Authentication & Authorization
- JWT tokens stored only in infrastructure layer (auth module)
- Role-based access control enforced at handler middleware level
- Refresh tokens stored in database with expiration tracking
- Password hashing using bcrypt with cost factor 10+ in pkg/password

### Data Protection
- Sensitive data (passwords, tokens) never logged
- PII fields encrypted at rest in database (future enhancement)
- Environment variables for secrets, never hardcoded
- SQL injection prevention via GORM parameterized queries

### API Security
- Rate limiting middleware (Redis-backed) per endpoint
- CORS configured for specific origins only
- Request validation at handler level before business logic
- Input sanitization for all user inputs
- HTTPS enforcement in production (via middleware)

### Infrastructure Security
- Database credentials via environment variables
- Redis password protection in production
- S3 access keys via environment variables
- Docker containers run as non-root user
- Secrets managed via Docker secrets or Kubernetes secrets (production)

## Performance Considerations

### Database Optimization
- Strategic indexes on foreign keys and frequently queried fields
- Connection pooling via GORM configuration
- N+1 query prevention via GORM Preload
- Read replicas for catalog queries (future)
- Query result caching in Redis

### Caching Strategy
- Cache-aside pattern for catalog data
- Redis for session storage and rate limiting
- TTL-based cache invalidation
- Cache warming for frequently accessed content
- Distributed cache for horizontal scaling (future)

### API Performance
- Pagination for all list endpoints
- Cursor-based pagination for large datasets
- Response compression (gzip middleware)
- Lazy loading of nested relationships
- Asynchronous event publishing (analytics)

### Resource Management
- Database connection limits configured
- Redis connection pooling
- HTTP client timeouts for external services
- Graceful shutdown handling
- Memory profiling hooks for production

## Migration Strategy

### Migration Files
Located in `/migrations` directory with naming convention:
- `NNNNNN_description.up.sql` - Apply migration
- `NNNNNN_description.down.sql` - Rollback migration

### Migration Tooling
- **Development**: GORM AutoMigrate for schema changes
- **Production**: Versioned SQL migrations with tracking table
- **Tool**: golang-migrate or custom migration runner
- **Rollback**: Down migrations for safe rollbacks

### Migration Workflow
1. Create migration file with descriptive name
2. Write up migration (DDL changes)
3. Write down migration (rollback DDL)
4. Test migration in development environment
5. Apply to staging before production
6. Monitor production migration execution

### Migration Tracking
```sql
CREATE TABLE schema_migrations (
    version BIGINT PRIMARY KEY,
    applied_at TIMESTAMP NOT NULL DEFAULT NOW()
);
```

### Best Practices
- Never modify existing migrations (create new ones)
- Keep migrations backward compatible when possible
- Use transactions in migration files
- Test migrations with realistic data volumes
- Document breaking changes in migration comments

### Security in Migrations
- No sensitive data in migration files
- Use parameterized queries if custom SQL
- Validate migration files in version control
- Review migrations in code review process

## Seeding Strategy

### Seed Data Location
- `/scripts/seed.sh` - Seed execution script
- `/test/fixtures/seed_data.go` - Test fixtures
- `/migrations/seeds/` - Production seed data (optional)

### Seed Data Categories
1. **Reference Data**: Genres, tags (always seeded)
2. **Test Data**: Sample series, episodes for development
3. **Demo Data**: Curated content for staging/demo
4. **Empty**: Production starts with reference data only

### Seed Execution
```bash
# Development - full seed
make seed

# Production - reference data only
make seed-reference

# Staging - demo data
make seed-demo
```

### Seed Data Security
- No real user credentials in seeds
- Placeholder content only
- S3 URLs pointing to sample media
- Test passwords clearly marked
- Seed data excluded from production builds

### Seed Data Performance
- Batch inserts for large datasets
- Disable indexes during bulk inserts
- Re-enable indexes after seeding
- Use COPY commands for PostgreSQL bulk load
- Parallel seed execution where safe

### Seed Data Versioning
- Seed files versioned alongside migrations
- Idempotent seed operations (check before insert)
- Seed data can be re-run without duplication
- Clear separation between reference and test data

#### `/internal/interface/http`
Shared HTTP infrastructure across all modules.
- `router.go`: Central route registration
- `middleware/`: Reusable HTTP middleware (auth, CORS, rate limiting, logging)
- `server.go`: HTTP server setup and configuration

### `/pkg`
Public library code that can be reused across projects or imported externally.

- **logger**: Structured JSON logging
- **password**: Password hashing (bcrypt)
- **validator**: Request validation helpers
- **utils**: Common utilities (UUID, slug generation, time)
- **storage**: S3-compatible storage abstraction

### `/migrations`
Database migration files (up and down SQL scripts).

### `/configs`
Configuration files and templates.

### `/scripts`
Utility scripts for development and deployment.

### `/test`
Integration tests and test fixtures organized by bounded context.

### `/deployments`
Docker configurations for local development and production.

### `/docs`
Project documentation (PRD, architecture, API spec, etc.).

## File Naming Conventions

- Go files: `snake_case.go`
- Test files: `*_test.go`
- Config files: `config.yaml`, `config.example.yaml`
- Migration files: `NNNNNN_description.up.sql`, `NNNNNN_description.down.sql`

## Import Path

The Go module will be defined as:
```
module github.com/yoosuf/reeltv
```

Internal imports will use:
```
import "github.com/yoosuf/reeltv/internal/model"
import "github.com/yoosuf/reeltv/internal/repository"
```

Package imports will use:
```
import "github.com/yoosuf/reeltv/pkg/logger"
import "github.com/yoosuf/reeltv/pkg/jwt"
```

## Environment Variables

Key environment variables (defined in `.env.example`):
- `APP_ENV`: Application environment (development, staging, production)
- `APP_PORT`: HTTP server port (default: 8080)
- `DB_HOST`: PostgreSQL host
- `DB_PORT`: PostgreSQL port (default: 5432)
- `DB_USER`: PostgreSQL user
- `DB_PASSWORD`: PostgreSQL password
- `DB_NAME`: PostgreSQL database name
- `REDIS_HOST`: Redis host
- `REDIS_PORT`: Redis port (default: 6379)
- `REDIS_PASSWORD`: Redis password
- `JWT_SECRET`: JWT signing secret
- `JWT_ACCESS_EXPIRATION`: Access token expiration (default: 15m)
- `JWT_REFRESH_EXPIRATION`: Refresh token expiration (default: 168h)
- `S3_ENDPOINT`: S3-compatible storage endpoint
- `S3_ACCESS_KEY`: S3 access key
- `S3_SECRET_KEY`: S3 secret key
- `S3_BUCKET`: S3 bucket name
- `S3_REGION`: S3 region

## Build Commands

Makefile will provide:
- `make build`: Build the application
- `make run`: Run the application locally
- `make test`: Run tests
- `make docker-up`: Start Docker Compose services
- `make docker-down`: Stop Docker Compose services
- `make migrate-up`: Run database migrations
- `make migrate-down`: Rollback database migrations
- `make seed`: Seed database with test data

## Development Workflow

1. Start services: `make docker-up`
2. Run migrations: `make migrate-up`
3. Seed data: `make seed`
4. Run application: `make run`
5. Run tests: `make test`

## Testing Strategy

- Unit tests for service layer logic
- Integration tests for API endpoints
- Test fixtures in `/test/fixtures`
- Tests runnable in Docker via `make test`
- API contract testing using the specification
