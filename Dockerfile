# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Allow Go to automatically download newer toolchains (like 1.25.0) if required by go.mod
ENV GOTOOLCHAIN=auto

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
# CGO_ENABLED=0 is used to ensure a statically linked binary (github.com/glebarez/sqlite is a pure Go driver)
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

# Run stage
FROM alpine:latest

WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/server .

# Copy the web directory (static files and templates)
COPY --from=builder /app/web ./web

# Copy the database directory so the existing database and path structure is maintained
COPY --from=builder /app/internal/database ./internal/database

# Copy the .env file to ensure default variables (like PORT and DATABASE_FILE_PATH) are set.
# Note: In production, consider passing environment variables securely via docker run --env-file or secrets.
COPY --from=builder /app/.env.example .env

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./server"]
