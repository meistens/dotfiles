FROM golang:1.24-alpine3.21 AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application, need to specify directory if you get a build error
# with this cmd line
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd

# Final stage
FROM alpine:latest

# RUN apk --no-cache add ca-certificates
RUN apk add ca-certificates
WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/main .

# Create logs directory
RUN mkdir -p /app/logs

# Expose port
EXPOSE 8080

# Run the application
CMD ["./main"]
