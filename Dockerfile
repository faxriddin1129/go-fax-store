# Multi-stage build: Builder stage
FROM golang:1.23.2-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git ca-certificates

# Set working directory
WORKDIR /app

# Copy dependency files first (for better Docker layer caching)
COPY go.mod go.sum ./

# Set Go build environment variables
ENV GOPROXY=https://goproxy.cn,https://proxy.golang.org,direct
ENV GOSUMDB=off
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GO111MODULE=on
ENV GOCACHE=/go-cache
ENV GOMODCACHE=/go-mod-cache

# Create cache directories
RUN mkdir -p /go-cache /go-mod-cache

# Download and cache dependencies
RUN --mount=type=cache,target=/go-mod-cache \
    --mount=type=cache,target=/go-cache \
    go mod download

# Copy source code
COPY . .

# Build the application with optimizations
RUN --mount=type=cache,target=/go-mod-cache \
    --mount=type=cache,target=/go-cache \
    go build -ldflags="-w -s -X main.version=1.0.0" \
    -a -installsuffix cgo \
    -o main ./main/main.go

# Production stage - minimal Alpine image
FROM alpine:latest

# Install runtime dependencies and create non-root user
RUN apk --no-cache add ca-certificates tzdata wget && \
    addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

# Set working directory
WORKDIR /app

# Copy compiled binary from builder stage
COPY --from=builder /app/main .

# Set proper ownership and permissions
RUN chown -R appuser:appgroup /app && \
    chmod +x main

# Switch to non-root user for security
USER appuser

# Health check using HTTP endpoint
HEALTHCHECK --interval=30s --timeout=10s --start-period=40s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8081/health || exit 1

# Expose application port
EXPOSE 8081

# Run the application
CMD ["./main"]