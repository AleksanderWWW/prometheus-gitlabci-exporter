# Stage 1: build
FROM golang:1.25.1 AS builder

WORKDIR /app

# Cache dependency installation
COPY go.mod go.sum ./
RUN go mod download

COPY exporter/ ./exporter/
COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o exporter .

# Stage 2: runtime (as non-root user)
FROM gcr.io/distroless/base-debian12:nonroot

WORKDIR /app

COPY --from=builder /app/exporter .

ENTRYPOINT ["/app/exporter"]
