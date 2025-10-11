#-----Build Stage-----
FROM golang:1.25.1 AS builder

WORKDIR /app

# Copy only go.mod and go.sum to cache module download
COPY go.mod go.sum ./
RUN go mod download

#Copy the source code
COPY . .

# Build a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o main main.go

# Make sure it’s executable
RUN chmod +x /app/main

# -----Final Stage-----
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/main .
ENTRYPOINT ["./main"]