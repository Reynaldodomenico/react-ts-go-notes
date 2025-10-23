# ---- Stage 1: Build ----
FROM golang:1.23-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server .

# ---- Stage 2: Run ----
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .

COPY .env .env

EXPOSE 8080

CMD ["./server"]
