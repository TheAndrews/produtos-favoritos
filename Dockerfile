# Build stage
FROM golang:1.24-bookworm AS builder

WORKDIR /app

# Copy go.mod and go.sum first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source
COPY . .

# Build the binary
RUN go build -o main ./src/main.go

# Final stage
FROM debian:bookworm-slim

# Install CA certificates so HTTPS works
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates && \
	rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]