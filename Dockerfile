# =========================
# Stage 1: Build
# =========================
FROM golang:1.24-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first (to leverage Docker cache)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the Go binary
RUN go build -o user-service main.go

# =========================
# Stage 2: Run
# =========================
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/user-service .

# Copy any other needed files (optional)
# COPY config/ ./config/
# COPY models/ ./models/
# COPY controllers/ ./controllers/
# COPY utils/ ./utils/
# COPY routes/ ./routes/

# Expose port
EXPOSE 8081

# Command to run the binary
CMD ["./user-service"]

