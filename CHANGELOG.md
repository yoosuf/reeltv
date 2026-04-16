# Changelog

All notable changes to Reel TV will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Initial MVP backend implementation with DDD architecture
- Authentication module with JWT tokens
- User management module
- Catalog module (series, seasons, episodes)
- Watch progress tracking
- My List / Favorites functionality
- Recommendation engine (heuristic-based)
- Subscription and entitlement checks
- Admin CMS APIs
- Analytics event ingestion
- Docker and Docker Compose configurations
- Database migrations and seed data
- Integration tests with Docker
- Frontend application using Next.js 16.x with TypeScript
- Tailwind CSS for frontend styling
- React components for UI
- API client library for frontend-backend communication
- Authentication utilities for frontend

### Changed
- Restructured project to monorepo with backend and frontend folders
- Updated module path to `github.com/yoosuf/reeltv/backend`
- Updated .gitignore to include frontend-specific ignores
- Updated documentation to reflect current monorepo structure

### Security
- JWT-based authentication
- Password hashing
- Admin access control
- Subscription-based access control

## [0.1.0] - 2024-04-16

### Added
- Initial project setup
- MVP PRD documentation
- System architecture documentation
- Domain model and database schema design
- API specification
- Repository structure
- GitHub documentation files for publication
- LICENSE, CONTRIBUTING, SECURITY, CHANGELOG, CODE_OF_CONDUCT files
- Initial project setup
- Domain-driven design architecture
- Modular structure with bounded contexts
- PostgreSQL database schema
- Redis caching layer
- Gin HTTP framework
- GORM ORM
- Zero logger
- UUID entity support

---

## Version History

### v0.1.0 (MVP Release)
- Core streaming platform functionality
- User authentication and authorization
- Content management
- Basic recommendations
- Watch progress tracking
- Favorites/My List
- Subscription management
- Admin dashboard
- Analytics tracking

### Future Roadmap
- Advanced ML-based recommendations
- Payment gateway integration
- Video streaming with HLS/DASH
- Mobile app support
- Advanced analytics dashboard
- Content delivery network integration
- Multi-language support
- Social features (sharing, comments)
- Advanced admin CMS
