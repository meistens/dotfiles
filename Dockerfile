# Select container to use for build
FROM golang:1.24-alpine3.21 AS builder
WORKDIR /app
# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download
# Copy source code
COPY . .
# Build the application, or build locally and copy, whichever works
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd

# Image/Container
FROM alpine:latest
# Install only required packages for the application
RUN apk add --no-cache ca-certificates

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/main .

# Create logs directory
RUN mkdir -p /app/logs

# Expose port
EXPOSE 8080

# Run the application directly
CMD ["./main"]
