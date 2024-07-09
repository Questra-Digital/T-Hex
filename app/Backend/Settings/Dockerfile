# Start from an official Golang image
FROM golang:1.19 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy only the go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

# Start from a minimal Alpine base image for the final image
FROM alpine:latest

# Install ca-certificates for SSL
RUN apk add --no-cache ca-certificates

# Set the working directory inside the container
WORKDIR /app

# Copy only the built binary from the builder stage
COPY --from=builder /app/server .

# Expose the port your application runs on (adjust as needed)
EXPOSE 9090

# Command to run the application
CMD ["./server"]
