FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/server/main.go

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# Copy the binary
COPY --from=builder /app/main .

# Copy templates and static files
COPY --from=builder /app/internal/templates ./internal/templates
COPY --from=builder /app/internal/static ./internal/static

# Expose port
EXPOSE 8080

# Run the application
CMD ["./main"]
