# Stage 1: build
FROM golang:1.21 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# build main.go
RUN go build -o app ./cmd/main.go

# Stage 2: runtime
FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/app .
COPY .env .  # .env ni ham image ichiga qo‘yamiz

EXPOSE 8080
CMD ["./app"]
