# Build stage
FROM golang:1.24-bullseye AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o mcp-netbird ./cmd/mcp-netbird

# Final stage
FROM debian:bullseye-slim

# Install ca-certificates for HTTPS requests
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Create a non-root user
RUN useradd -r -u 1000 -m mcp-netbird

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder --chown=1000:1000 /app/mcp-netbird /app/

# Use the non-root user
USER mcp-netbird

# Run the application in stdio mode
ENTRYPOINT ["/app/mcp-netbird"]
