# BUILDER
FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod .
RUN go mod download

COPY . .

RUN go build -o go-app

# RUNTIME
FROM debian:stable-slim

RUN useradd -m appuser

WORKDIR /app

COPY --from=builder /app/go-app /app/go-app
RUN chown -R appuser:appuser /app \
    && chmod 755 /app/go-app

USER appuser
    
EXPOSE 8080

CMD ["./go-app"]