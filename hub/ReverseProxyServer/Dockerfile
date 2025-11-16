# Stage 1 — build the Go binary
FROM golang:1.23.3 AS builder

WORKDIR /app

# 1️⃣ Copy only go.mod and go.sum first — this allows Docker to cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# 2️⃣ Now copy the rest of the source code
COPY . .

# 3️⃣ Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o main

# Stage 2 — minimal runtime image
FROM scratch
COPY --from=builder /app/main /main
EXPOSE 4445
ENTRYPOINT ["/main"]
