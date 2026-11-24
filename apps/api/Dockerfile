# ── Dev stage ─────────────────────────────────────────────────────────────────
FROM golang:1.26-alpine AS dev

RUN apk add --no-cache git \
 && go install github.com/air-verse/air@latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Source is mounted at runtime — air will watch and rebuild
CMD ["air", "-c", ".air.toml"]

# ── Build stage ──────────────────────────────────────────────────────────────
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Install git (needed by some go modules)
RUN apk add --no-cache git

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source and build a statically-linked binary
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o /app/server ./cmd/api

# ── Run stage ─────────────────────────────────────────────────────────────────
FROM gcr.io/distroless/static-debian12 AS runner

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 3000

ENTRYPOINT ["/app/server"]
