# First Stage: Build the Go application
FROM golang:1.24 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first to leverage caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project into the container
COPY . .

# Build the Go application from the cmd folder
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd

# Second Stage: Minimal Alpine image for runtime
FROM alpine:3.18

# Install CA certificates for HTTPS support
RUN apk --no-cache add ca-certificates

# Set up the working directory
WORKDIR /home/appuser/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Copy SSL certificate and key (Ensure these files exist in the project root)
COPY server.crt server.key .  

# Create a non-root user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
RUN chown -R appuser:appgroup /home/appuser && chmod +x /home/appuser/main

# Switch to the non-root user
USER appuser

# Expose the HTTPS port
EXPOSE 8443

# Run the application with SSL
CMD ["./main", "-cert", "server.crt", "-key", "server.key"]