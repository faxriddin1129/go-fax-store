# Stage 1: build
FROM golang:1.24 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# build main.go with static linking for Alpine
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/main.go

# Stage 2: runtime
FROM alpine:latest

# Install ca-certificates for HTTPS requests, wget for healthcheck, and timezone data
RUN apk add --no-cache ca-certificates wget tzdata

# Create a non-root user
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

WORKDIR /app

# Copy the binary from builder stage
COPY --from=builder /app/app .
COPY --from=builder /app/.env .

# Create logs directory and set permissions
RUN mkdir -p /app/storage/logs && \
    chown -R appuser:appgroup /app

    
USER appuser

EXPOSE 8081
CMD ["./app"]