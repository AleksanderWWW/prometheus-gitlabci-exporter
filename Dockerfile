# Stage 1: build
FROM golang:1.25.1@sha256:d7098379b7da665ab25b99795465ec320b1ca9d4addb9f77409c4827dc904211 AS builder

WORKDIR /app

# Cache dependency installation
COPY go.mod go.sum ./
RUN go mod download

COPY exporter/ ./exporter/
COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o prom_exporter .

# Stage 2: runtime (as non-root user)
FROM gcr.io/distroless/base-debian12:nonroot@sha256:10136f394cbc891efa9f20974a48843f21a6b3cbde55b1778582195d6726fa85

WORKDIR /app

COPY --from=builder /app/prom_exporter .

ENTRYPOINT ["/app/prom_exporter"]
