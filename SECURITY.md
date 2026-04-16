# Security Policy

## Supported Versions

Currently, only the latest version of Reel TV is supported with security updates.

## Reporting a Vulnerability

If you discover a security vulnerability, please report it responsibly.

**Do not** create a public issue for security vulnerabilities.

### How to Report

Send an email to: security@yoosuf.com

Please include:
- Description of the vulnerability
- Steps to reproduce
- Potential impact
- Suggested fix (if known)

### What to Expect

- We will acknowledge receipt within 48 hours
- We will provide a detailed response within 7 days
- We will work with you to understand and resolve the issue
- Once resolved, we will release a security update and credit you in the release notes

## Security Best Practices

### For Developers
- Keep dependencies up to date
- Use environment variables for sensitive data
- Never commit secrets or API keys
- Review code changes carefully
- Use security scanning tools

### For Operators
- Use HTTPS in production
- Keep the application updated
- Use strong passwords and JWT secrets
- Implement rate limiting
- Monitor logs for suspicious activity
- Regularly backup data

### Known Security Considerations

1. **JWT Tokens**: Ensure JWT secrets are kept secure and rotated regularly
2. **Database**: Use strong passwords and enable SSL in production
3. **Redis**: Enable authentication and restrict network access
4. **API Keys**: Never commit API keys to the repository
5. **Dependencies**: Regularly audit and update dependencies

## Dependency Scanning

We regularly scan our dependencies for known vulnerabilities using:
- `go mod` for Go dependencies
- Docker security scanning
- GitHub Dependabot

## Security Updates

Security updates will be:
- Released as soon as feasible
- Documented in the changelog
- Backported to supported versions if necessary
- Announced via security advisories
