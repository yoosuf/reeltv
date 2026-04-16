# Contributing to Reel TV

Thank you for your interest in contributing to Reel TV! This document provides guidelines and instructions for contributing to the project.

## Project Structure

Reel TV is a monorepo with two main applications:

- **Backend** (`backend/`): Go-based REST API with DDD architecture
- **Frontend** (`frontend/`): Next.js 16.x SPA with TypeScript and Tailwind CSS

## Getting Started

### Prerequisites
- Go 1.21 or higher (backend)
- Node.js 18+ and npm (frontend)
- Docker and Docker Compose
- PostgreSQL 15+
- Redis 7+

### Setup Development Environment

1. Clone the repository:
```bash
git clone https://github.com/yoosuf/reeltv.git
cd reeltv
```

2. Copy environment variables:
```bash
cp .env.example .env
cd frontend && cp .env.example .env.local && cd ..
```

3. Start Docker services:
```bash
make docker-up
```

4. Run database migrations:
```bash
make migrate-up
```

5. Seed database:
```bash
make seed
```

6. Install dependencies:
```bash
# Backend
make deps

# Frontend
cd frontend
npm install
cd ..
```

7. Run the applications:
```bash
# Backend (in separate terminal)
make run

# Frontend (in separate terminal)
cd frontend
npm run dev
```

The API will be available at `http://localhost:8080`
The frontend will be available at `http://localhost:3000`

## Development Workflow

### Backend Development
- Follow Go conventions and best practices
- Use `gofmt` to format code
- Run linter before committing: `make lint`
- Write tests for new features
- Run tests: `make test`

### Frontend Development
- Follow TypeScript and React best practices
- Use ESLint and Prettier for code formatting
- Run linter: `cd frontend && npm run lint`
- Write tests for new components
- Run tests: `cd frontend && npm test`
- Add tests for new features
- Ensure all tests pass: `make test`

### Branching Strategy
- `main` - production branch
- `develop` - development branch
- `feature/*` - feature branches
- `bugfix/*` - bug fix branches
- `hotfix/*` - urgent production fixes

### Commit Message Format
Use conventional commit format:
```
type(scope): description

[optional body]

[optional footer]
```

Types:
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes (formatting, etc.)
- `refactor`: Code refactoring
- `test`: Adding or updating tests
- `chore`: Maintenance tasks

Example:
```
feat(auth): add refresh token support

Implement JWT refresh token rotation to improve security.
Users can now refresh their access tokens without re-authenticating.
```

### Pull Request Process

1. Create a feature branch from `develop`
2. Make your changes
3. Add tests and ensure they pass
4. Update documentation if needed
5. Submit a pull request to `develop`
6. Address review feedback
7. Once approved, merge to `develop`

### Testing

Run tests locally:
```bash
make test
```

Run tests with coverage:
```bash
make test-coverage
```

Run tests with Docker:
```bash
make test-docker
```

## Project Structure

```
reeltv/
├── cmd/                    # Application entry points
│   └── api/               # API server
├── internal/              # Private application code
│   ├── auth/             # Authentication module
│   ├── user/             # User management
│   ├── catalog/          # Content catalog
│   ├── watchprogress/    # Watch progress tracking
│   ├── mylist/           # Favorites/My List
│   ├── recommendation/    # Recommendations
│   ├── subscription/     # Subscription management
│   ├── admin/            # Admin CMS
│   ├── analytics/        # Analytics tracking
│   ├── shared/           # Shared utilities
│   ├── interface/        # HTTP layer
│   └── app/              # Application bootstrap
├── pkg/                  # Public libraries
├── migrations/           # Database migrations
├── seeds/               # Seed data
├── docs/                # Documentation
├── deployments/         # Docker configurations
└── tests/               # Integration tests
```

## Architecture

This project follows Domain-Driven Design (DDD) principles with clean architecture:

- **Domain Layer**: Core business logic and entities
- **Application Layer**: Use cases and application services
- **Infrastructure Layer**: External dependencies (database, Redis, etc.)
- **Interface Layer**: HTTP handlers and routing

## Documentation

- [MVP PRD](docs/01-MVP-PRD.md)
- [System Architecture](docs/02-System-Architecture.md)
- [Domain Model & Database Schema](docs/03-Domain-Model-Database-Schema.md)
- [API Specification](docs/04-API-Specification.md)
- [Repository Structure](docs/05-Repository-Structure.md)
- [API Testing Guide](docs/API-Testing-Guide.md)

## Reporting Issues

When reporting bugs or suggesting features:
1. Check existing issues first
2. Use the issue template if available
3. Provide clear description and steps to reproduce
4. Include environment details (OS, Go version, etc.)

## Questions

For questions or discussions:
- Open a GitHub issue with the "question" label
- Join our community discussions

## License

By contributing, you agree that your contributions will be licensed under the MIT License.
