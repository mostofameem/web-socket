# Build stage
FROM golang:1.23.5-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN go build -o main main.go

# Run final stage
FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 80

# Command to run the executable
CMD ["/app/main", "serve-rest"]
