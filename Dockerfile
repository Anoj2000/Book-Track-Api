# Build stage
FROM golang:1.21-alpine AS builder

# Set Go env
ENV GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux

WORKDIR /app

# First copy only module files for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy remaining files
COPY . .
RUN go build -o book-api .

# Runtime stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/book-api .
COPY --from=builder /app/books.db ./

EXPOSE 3000
CMD ["./book-api"]