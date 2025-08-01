# Build stage
FROM golang:1.22-alpine AS builder

# Install git for fetching dependencies
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o wipeOs main.go

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Create non-root user
RUN addgroup -g 1001 -S wipeOs && \
    adduser -u 1001 -S wipeOs -G wipeOs

WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/wipeOs .

# Copy assets
COPY --from=builder /app/assets ./assets

# Change ownership
RUN chown -R wipeOs:wipeOs /root

# Switch to non-root user
USER wipeOs

# Expose port (if needed for web interface in future)
EXPOSE 8080

# Command to run
ENTRYPOINT ["./wipeOs"] 