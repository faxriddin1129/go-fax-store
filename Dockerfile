FROM golang:1.23.2-alpine AS builder

RUN apk add --no-cache git ca-certificates

WORKDIR /app

COPY go.mod go.sum ./

ENV GOPROXY=https://goproxy.cn,https://proxy.golang.org,direct
ENV GOSUMDB=off
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GO111MODULE=on
ENV GOCACHE=/go-cache
ENV GOMODCACHE=/go-mod-cache

RUN mkdir -p /go-cache /go-mod-cache

RUN --mount=type=cache,target=/go-mod-cache \
    --mount=type=cache,target=/go-cache \
    go mod download

COPY . .

RUN go build -ldflags="-w -s -X main.version=1.0.0" -a -installsuffix cgo -o main ./cmd/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata && \
    addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

WORKDIR /app

COPY --from=builder /app/main .

RUN chown appuser:appgroup main && chmod +x main

USER appuser

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD pgrep main || exit 1

EXPOSE ${PROJECT_PORT}

CMD ["./main"]