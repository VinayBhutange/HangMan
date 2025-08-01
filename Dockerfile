# Build stage
FROM golang:1.21-alpine AS builder

# Install git and ca-certificates (needed for fetching dependencies)
RUN apk add --no-cache git ca-certificates

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download && go mod verify

# Copy source code
COPY . .

# Build the application
ARG VERSION=dev
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-s -w -X main.version=${VERSION}" \
    -o hangman \
    main.go

# Runtime stage
FROM alpine:latest

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Create non-root user
RUN addgroup -g 1001 hangman && \
    adduser -D -u 1001 -G hangman hangman

# Set working directory
WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/hangman .
COPY --from=builder /app/data ./data

# Change ownership
RUN chown -R hangman:hangman /app

# Switch to non-root user
USER hangman

# Create volume for statistics persistence
VOLUME ["/app/.hangman"]

# Set entrypoint
ENTRYPOINT ["./hangman"]

# Metadata
LABEL org.opencontainers.image.title="Hangman Game"
LABEL org.opencontainers.image.description="A terminal-based Hangman game written in Go"
LABEL org.opencontainers.image.vendor="Hangman Game Project"
LABEL org.opencontainers.image.licenses="MIT"
