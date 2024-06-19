FROM golang:1.22-bullseye AS builder
LABEL authors="developer"
ENV CGO_ENABLED=1

WORKDIR /app

# Download necessary Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o user-svc .

# Stage 2: Run the Go application
FROM debian:bullseye-slim

# Install SQLite3
RUN apt-get update && \
    apt-get install -y sqlite3 && \
    rm -rf /var/lib/apt/lists/*

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/user-svc .
COPY --from=builder /app/*.sqlite .

# Expose port 3000 to the outside world
EXPOSE 3000

# Command to run the executable
CMD ["./user-svc"]
