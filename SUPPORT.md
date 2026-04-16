# Support

## Getting Help

If you need help with Reel TV, here are the best ways to get support:

## Documentation

First, check our comprehensive documentation:

- [README](README.md) - Project overview and getting started
- [Deployment Guide](DEPLOYMENT.md) - Backend and frontend deployment
- [API Testing Guide](docs/API-Testing-Guide.md) - API usage examples
- [System Architecture](docs/02-System-Architecture.md) - Architecture details
- [API Specification](docs/04-API-Specification.md) - Complete API reference

## Community Support

### GitHub Issues
For bug reports and feature requests, please open an issue on GitHub:
- [Report a Bug](https://github.com/yoosuf/reeltv/issues/new?template=bug_report.md)
- [Request a Feature](https://github.com/yoosuf/reeltv/issues/new?template=feature_request.md)

### Discussions
For general questions, ideas, or discussions:
- [GitHub Discussions](https://github.com/yoosuf/reeltv/discussions)

## Professional Support

For enterprise support, custom development, or consulting:
- Email: support@yoosuf.com
- Website: https://yoosuf.com

## Common Issues

### Database Connection Issues
If you're having trouble connecting to the database:
1. Ensure Docker is running: `docker ps`
2. Check environment variables in `.env`
3. Verify PostgreSQL is accessible: `make docker-logs`

### Redis Connection Issues
If Redis is not responding:
1. Check Redis container status: `docker ps`
2. View Redis logs: `docker logs reeltv-redis-1`
3. Verify Redis configuration in `.env`

### Build Errors
If you encounter build errors:
1. Ensure you have Go 1.21+ installed: `go version`
2. Update dependencies: `make deps`
3. Clean build artifacts: `make clean`
4. Rebuild: `make build`

### Test Failures
If tests are failing:
1. Ensure test infrastructure is running: `make docker-test-up`
2. Run tests with verbose output: `go test -v ./...`
3. Check test database connection

## Reporting Security Issues

**Do not** report security vulnerabilities through public issues.

For security vulnerabilities, please email: security@yoosuf.com

See our [Security Policy](SECURITY.md) for more details.

## Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

## Status Page

Check our service status at: https://status.yoosuf.com

## Response Times

- GitHub Issues: Within 2-3 business days
- Security Issues: Within 24 hours
- Enterprise Support: As per SLA

## Additional Resources

- [Changelog](CHANGELOG.md) - Version history and release notes
- [Code of Conduct](CODE_OF_CONDUCT.md) - Community guidelines
- [License](LICENSE) - MIT License information
