# Stage 1: Build stage
FROM golang:1.23.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

# Build the Go binary with optimization flags
RUN go build -trimpath -ldflags="-s -w" -o cutlink-api ./cmd

# Stage 2: Final image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/cutlink-api .

EXPOSE 8080

CMD ["./cutlink-api"]