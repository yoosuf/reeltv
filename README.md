# Reel TV - Vertical Short-Drama Streaming Platform

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org)
[![Next.js](https://img.shields.io/badge/Next.js-16.x-black?style=flat&logo=next.js)](https://nextjs.org/)
[![TypeScript](https://img.shields.io/badge/TypeScript-5.0+-blue?style=flat&logo=typescript)](https://www.typescriptlang.org/)
[![Docker](https://img.shields.io/badge/Docker-Ready-blue?style=flat&logo=docker)](https://www.docker.com/)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](http://makeapullrequest.com)
[![GitHub issues](https://img.shields.io/github/issues/yoosuf/reeltv)](https://github.com/yoosuf/reeltv/issues)
[![GitHub stars](https://img.shields.io/github/stars/yoosuf/reeltv?style=social)](https://github.com/yoosuf/reeltv/stargazers)

<div align="center">

**A curated vertical short-drama streaming platform built with Go (backend) and Next.js (frontend), following Domain-Driven Design (DDD) principles.**

[⭐ Star this project](https://github.com/yoosuf/reeltv/stargazers) · [🐛 Report a Bug](https://github.com/yoosuf/reeltv/issues/new?template=bug_report.md) · [💡 Feature Request](https://github.com/yoosuf/reeltv/issues/new?template=feature_request.md)

</div>

## License

Copyright © 2026 Yoosuf

## ✨ Features

- 🎬 **Short-Form Content**: Professionally produced 1-2 minute dramatic episodes
- 📱 **Mobile-First Design**: Optimized for mobile viewing experience
- 🔐 **Secure Authentication**: JWT-based auth with refresh tokens
- 📊 **Watch Progress Tracking**: Continue watching where you left off
- ❤️ **My List**: Save your favorite series for later
- 🎯 **Personalized Recommendations**: AI-powered content suggestions
- 💳 **Subscription Management**: Premium tier with subscription checks
- 📈 **Analytics**: Built-in event tracking and analytics
- 🛡️ **Admin CMS**: Content management for administrators
- 🚀 **High Performance**: Redis caching and optimized database queries
- 🐳 **Docker Ready**: Easy deployment with Docker and Docker Compose
- 📝 **Type-Safe**: Full TypeScript coverage on frontend

## 🏗️ Architecture

The project follows a monorepo structure with separate backend and frontend applications:

- **Backend**: Go-based REST API with DDD architecture
- **Frontend**: Next.js 16.x SPA with TypeScript and Tailwind CSS
- **Database**: PostgreSQL for data persistence
- **Cache**: Redis for caching and session management
- **Storage**: S3-compatible for media storage

## Documentation

- [Getting Started](#getting-started)
- [Deployment Guide](DEPLOYMENT.md)
- [Contributing Guide](CONTRIBUTING.md)
- [API Documentation](docs/04-API-Specification.md)
- [API Testing Guide](docs/API-Testing-Guide.md)
- [Security Policy](SECURITY.md)
- [Support](SUPPORT.md)
- [Changelog](CHANGELOG.md)

## 🚀 Quick Start

Get Reel TV running in under 5 minutes:

```bash
# Clone the repository
git clone https://github.com/yoosuf/reeltv.git
cd reeltv

# Copy environment variables
cp .env.example .env
cd frontend && cp .env.example .env.local && cd ..

# Start all services with Docker
make docker-up

# Run database migrations
make migrate-up

# Seed database with test data
make seed

# Start the applications
# Backend (in one terminal)
make run

# Frontend (in another terminal)
cd frontend
npm install
npm run dev
```

Visit http://localhost:3000 to see the frontend and http://localhost:8080 for the API.

## 🗺️ Roadmap

- [ ] Mobile apps (iOS/Android)
- [ ] Video streaming with adaptive bitrate
- [ ] Social features (comments, sharing)
- [ ] Advanced recommendation engine with ML
- [ ] Multi-language support
- [ ] Payment integration (Stripe, PayPal)
- [ ] Content delivery network (CDN) integration
- [ ] Real-time notifications
- [ ] Live streaming capabilities
- [ ] Content creator portal

## 🤝 Contributing

We welcome contributions from the community! Whether you're fixing a bug, adding a feature, or improving documentation, we appreciate your help.

**How to Contribute:**
1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

See [CONTRIBUTING.md](CONTRIBUTING.md) for detailed guidelines.

## 🌟 Star History

If you find this project useful, please consider giving it a ⭐ star on GitHub. It helps others discover the project and motivates us to continue development.

[![Star History Chart](https://api.star-history.com/svg?repos=yoosuf/reeltv&type=Date)](https://star-history.com/#yoosuf/reeltv&Date)

## 💬 Community

Join our community to connect with other developers and users:

- **GitHub Discussions**: [Join the conversation](https://github.com/yoosuf/reeltv/discussions)
- **GitHub Issues**: [Report bugs or request features](https://github.com/yoosuf/reeltv/issues)
- **Twitter**: Follow for updates [@yoosuf](https://twitter.com/yoosuf)

## 🙏 Acknowledgments

- Built with [Go](https://golang.org/) and [Next.js](https://nextjs.org/)
- Inspired by modern streaming platforms
- Thanks to all contributors who help improve this project

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 📞 Support

Need help? Check out our [SUPPORT.md](SUPPORT.md) or open an issue on GitHub.

---

<div align="center">

**Made with ❤️ by [Yoosuf](https://github.com/yoosuf)**

[⬆ Back to Top](#reel-tv---vertical-short-drama-streaming-platform)

</div>

## Tech Stack

### Backend
- **Language**: Go 1.21+
- **Architecture**: Modular DDD (Domain-Driven Design)
- **Database**: PostgreSQL 15
- **Cache**: Redis 7
- **Storage**: S3-compatible (MinIO for local dev)
- **ORM**: GORM
- **HTTP Framework**: Gin
- **Authentication**: JWT
- **Logging**: zerolog
- **Containerization**: Docker + Docker Compose

### Frontend
- **Framework**: Next.js 16.x
- **Language**: TypeScript
- **Styling**: Tailwind CSS
- **State Management**: React Hooks
- **HTTP Client**: Fetch API

## Architecture

The project follows a modular DDD approach with bounded contexts:

- **user**: User profile management
- **auth**: Authentication and authorization
- **catalog**: Series, seasons, episodes, genres, tags
- **playback**: Watch progress tracking
- **mylist**: User favorites
- **recommendation**: Content recommendations
- **subscription**: Subscription and entitlement checks
- **analytics**: Event ingestion
- **admin**: Content management APIs

Each bounded context has:
- `domain/`: Core business logic (entities, repository interfaces, domain services)
- `application/`: Use cases (application services, DTOs)
- `infrastructure/`: Technical implementations (persistence, caching, external services)
- `interface/http/`: HTTP handlers and DTOs

## Getting Started

### Prerequisites

- Docker and Docker Compose
- Go 1.21+ (for local development without Docker)
- Make (optional, for convenience commands)

### Quick Start with Docker

```bash
# Copy environment variables
cp .env.example .env
cd frontend && cp .env.example .env.local && cd ..

# Start all services
make docker-up

# Run database migrations
make migrate-up

# Seed database with test data
make seed

# The API will be available at http://localhost:8080
# The frontend will be available at http://localhost:3000
```

### Local Development

```bash
# Backend
# Install dependencies
make deps

# Start infrastructure services
make docker-up

# Run migrations
make migrate-up

# Seed database
make seed

# Run the application
make run

# Frontend
cd frontend

# Install dependencies
npm install

# Start development server
npm run dev
```

## Project Structure

```
reeltv/
├── backend/              # Backend application (Go)
│   ├── cmd/api/          # Application entry point
│   ├── internal/
│   │   ├── config/       # Configuration management
│   │   ├── shared/       # Shared domain and infrastructure
│   │   ├── user/         # User bounded context
│   │   ├── auth/         # Authentication bounded context
│   │   ├── catalog/      # Catalog bounded context
│   │   ├── watchprogress/# Watch progress bounded context
│   │   ├── mylist/       # My List bounded context
│   │   ├── recommendation/# Recommendation bounded context
│   │   ├── subscription/ # Subscription bounded context
│   │   ├── analytics/    # Analytics bounded context
│   │   ├── admin/        # Admin bounded context
│   │   └── interface/http/# Shared HTTP infrastructure
│   ├── pkg/              # Public utilities
│   ├── migrations/       # Database migrations
│   ├── seeds/            # Seed data
│   ├── tests/            # Integration tests
│   ├── go.mod            # Go module definition
│   └── go.sum            # Go module checksums
├── frontend/             # Frontend application (Next.js)
│   ├── app/              # Next.js App Router pages
│   ├── components/       # React components
│   ├── lib/              # Utility functions and API client
│   ├── public/           # Static assets
│   ├── package.json      # Node.js dependencies
│   └── tsconfig.json     # TypeScript configuration
├── deployments/          # Docker configurations
├── docs/                 # Project documentation
├── .github/              # GitHub configurations
└── Makefile              # Build commands
```

## API Documentation

See `docs/04-API-Specification.md` for complete API documentation.

### Base URL

- Development: `http://localhost:8080/api/v1`

### Health Check

```bash
curl http://localhost:8080/health
```

## Available Commands

### Backend (Go)
```bash
make build          # Build the backend application
make run            # Run the backend application locally
make test           # Run backend tests
make docker-up      # Start Docker services (DB, Redis, MinIO)
make docker-down    # Stop Docker services
make migrate-up     # Run database migrations
make migrate-down   # Rollback database migrations
make seed           # Seed database with test data
make deps           # Download Go dependencies
make fmt            # Format Go code
make clean          # Clean build artifacts
make dev            # Start development environment
```

### Frontend (Next.js)
```bash
cd frontend
npm install         # Install dependencies
npm run dev         # Start development server (http://localhost:3000)
npm run build       # Build for production
npm run start       # Start production server
npm run lint        # Run linter
```

### Docker
```bash
make docker-up      # Start all services
make docker-down    # Stop all services
make docker-logs    # View Docker logs
```

## Security Considerations

- JWT-based authentication with refresh tokens
- Password hashing using bcrypt
- Rate limiting via Redis
- CORS configuration
- Input validation at handler level
- SQL injection prevention via GORM parameterized queries
- Secrets managed via environment variables
- Non-root Docker container execution

## Performance Considerations

- Strategic database indexes
- Connection pooling
- Redis caching for catalog data
- Pagination for all list endpoints
- Response compression (gzip)
- N+1 query prevention via GORM Preload

## Migration Strategy

- Versioned SQL migrations in `/migrations`
- Up and down migrations for safe rollbacks
- Migration tracking table
- GORM AutoMigrate for development
- SQL migrations for production

## Seeding Strategy

- Reference data (genres, tags) always seeded
- Test data for development
- Demo data for staging
- Production starts with reference data only
- Idempotent seed operations

## Development Workflow

1. Start infrastructure: `make docker-up`
2. Run migrations: `make migrate-up`
3. Seed data: `make seed`
4. Run application: `make run`
5. Run tests: `make test`

## Documentation

- [MVP PRD](docs/01-MVP-PRD.md) - Product Requirements
- [System Architecture](docs/02-System-Architecture.md) - Architecture Overview
- [Domain Model & Schema](docs/03-Domain-Model-Database-Schema.md) - Database Design
- [API Specification](docs/04-API-Specification.md) - API Contract
- [Repository Structure](docs/05-Repository-Structure.md) - Code Organization

## License

Copyright © 2024 Crew Digital
