# Multi-stage build for Reel TV API

# Stage 1: Build
FROM golang:1.26-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git ca-certificates

# Set working directory
WORKDIR /build

# Copy go mod files
COPY backend/go.mod backend/go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY backend/ .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o reeltv-api ./cmd/api

# Stage 2: Runtime
FROM alpine:latest

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates curl

# Create non-root user
RUN addgroup -g 1000 reeltv && \
    adduser -D -u 1000 -G reeltv reeltv

# Set working directory
WORKDIR /app

# Copy binary from builder
COPY --from=builder /build/reeltv-api .

# Copy migrations
COPY --from=builder /build/migrations ./migrations

# Change ownership to non-root user
RUN chown -R reeltv:reeltv /app

# Switch to non-root user
USER reeltv

# Expose port
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
    CMD curl -f http://localhost:8080/health || exit 1

# Run the application
CMD ["./reeltv-api"]
