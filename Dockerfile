# ---- Build Stage ----
FROM golang:1.23-alpine AS builder

# Install git (required for go mod if using private repos)
RUN apk add --no-cache git

# Set the working directory inside the container
WORKDIR /app

# Copy go mod and sum files first, for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go binary
RUN go build -o url-shortener ./cmd/server

# ---- Run Stage ----
FROM alpine:latest

# Install ca-certificates for HTTPS
RUN apk add --no-cache ca-certificates

# Set working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/url-shortener .

# Copy any config files if needed (uncomment if you have config files)
# COPY config ./config

# Expose the port your app runs on (change if not 8080)
EXPOSE 8080

# Command to run the binary
ENTRYPOINT ["./url-shortener"]